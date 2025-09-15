package mongodb

import (
	"context"
	"log"
	"slices"

	"github.com/squishedfox/fictional-fiesta/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	validFields = []string{"id", "_id", "name"}
)

func getFilterFormRequest(req *db.GetFormsModel) bson.D {
	filter := bson.D{}

	if req == nil {
		return filter
	}
	if len(req.Filters) == 0 {
		return filter
	}

	// supported operations https://www.mongodb.com/docs/manual/reference/mql/query-predicates/#std-label-query-predicates-ref
	for _, operation := range req.Filters {
		if !slices.Contains(validFields, operation.Field) {
			continue
		}
		var operationKey string
		switch operation.Operation {
		case db.EqualsOperation:
			operationKey = "$eq"
			break
		case db.GreaterThanOperation:
			operationKey = "$gt"
			break
		case db.LessThanOperation:
			operationKey = "$lt"
			break
		case db.NotEqualsOperation:
			operationKey = "$ne"
			break
		default:
			break
		}

		if len(operationKey) != 0 {
			filter = append(filter, bson.E{
				Key: operation.Field,
				Value: bson.E{
					Key:   operationKey,
					Value: operation.Value,
				},
			})
		}
	}

	return filter
}

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
	filter := getFilterFormRequest(model)

	// count filter should not incldue things like limit or skip
	countFilter := bson.D{}
	copy(countFilter, filter)

	opts := options.Find().SetLimit(model.Limit).SetSkip(model.Skip)
	forms := make([]*db.FormModel, 0)
	cursor, err := collection.Find(r.context, filter, opts)
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
	count, err := collection.CountDocuments(r.context, countFilter)
	if err != nil {
		return nil, err
	}
	return &db.FormsModel{
		Forms: forms,
		Count: count,
	}, nil
}
