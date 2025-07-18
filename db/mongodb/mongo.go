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
		form := &db.FormsModel{}
		if err := cursor.Decode(&form); err != nil {
			return nil, err
		}
	}
	return &db.FormsModel{
		Forms: forms,
	}, nil
}
