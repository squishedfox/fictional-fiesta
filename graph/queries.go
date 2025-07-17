package graph

import (
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/squishedfox/fictional-fiesta/db"
)

type (
	DbContextKey string
)

var (
	ContextKey  DbContextKey = "DbContext"
	FormQueries              = graphql.NewObject(graphql.ObjectConfig{
		Name: "FormsQueries",
		Fields: graphql.Fields{
			"forms": &graphql.Field{
				Type: graphql.NewList(FormObject),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						DefaultValue: "",
						Type:         graphql.ID,
					},
					"name": &graphql.ArgumentConfig{
						DefaultValue: "",
						Type:         graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
						DefaultValue: 10,
						Description:  "Total number records to retreive and skip per page",
						Type:         graphql.Int,
					},
					"page": &graphql.ArgumentConfig{
						DefaultValue: 0,
						Description:  "Number of pages to skip before retreiving records",
						Type:         graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					// don't need a transaction for read only session
					session := p.Context.Value(ContextKey).(*mongo.Session)
					if session == nil {
						log.Fatal("could not find database driver")
					}

					database := session.Client().Database("forms", options.Database())
					collection := database.Collection("forms", options.Collection())
					filter := bson.D{}

					forms := make([]db.FormModel, 0)
					results, err := collection.Find(p.Context, filter)
					if err != nil {
						return nil, err
					}

					if err := results.All(p.Context, &forms); err != nil {
						return nil, err
					}
					log.Printf("Forms = %s", forms)
					return forms, nil
				},
			},
		},
	})
)
