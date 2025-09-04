package mongodb

import (
	"context"
	"log"

	"github.com/squishedfox/fictional-fiesta/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type (
	formRepository struct {
		context        context.Context
		session        *mongo.Session
		collectionName string
	}
)

// CreateForm implements db.FormsRepository.
func (r *formRepository) CreateForm(model *db.CreateFormModel) (any, error) {
	if err := r.session.StartTransaction(); err != nil {
		return nil, err
	}
	database := r.session.Client().Database("fictional-fiesta", options.Database())
	collection := database.Collection("forms", options.Collection())

	document := mapFormModelToDocument(model)
	result, err := collection.InsertOne(r.context, document)
	if err != nil {
		return nil, err
	}

	if err := r.session.CommitTransaction(r.context); err != nil {
		return nil, err
	}
	return result.InsertedID.(bson.ObjectID).Hex(), nil
}

func NewFormRepository(ctx context.Context, session *mongo.Session) db.FormsRepository {
	return &formRepository{
		context:        ctx,
		session:        session,
		collectionName: "forms",
	}
}

func (r *formRepository) GetForms(model *db.GetFormsModel) (*db.FormsModel, error) {
	// don't need a transaction for read only session
	database := r.session.Client().Database("fictional-fiesta", options.Database())
	collection := database.Collection("forms", options.Collection())
	filter := bson.D{}

	forms := make([]*db.FormModel, 0)
	cursor, err := collection.Find(r.context, filter)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cursor.Close(r.context); err != nil {
			log.Println(err)
		}
	}()
	for cursor.Next(r.context) {
		bsonForm := bson.M{}
		if err := cursor.Decode(&bsonForm); err != nil {
			return nil, err
		}
		form := db.FormModel{
			ID:   bsonForm["_id"].(bson.ObjectID).Hex(),
			Name: bsonForm["name"].(string),
		}
		forms = append(forms, &form)
	}
	count, err := collection.CountDocuments(r.context, filter)
	if err != nil {
		return nil, err
	}
	return &db.FormsModel{
		Forms: forms,
		Count: count,
	}, nil
}
