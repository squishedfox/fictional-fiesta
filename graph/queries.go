package graph

import (
	"log"
	"reflect"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					client := p.Context.Value(ContextKey).(*mongo.Client)
					if client == nil {
						log.Fatal("could not find database driver")
					}

					opts := options.Find().SetLimit()
					limit := p.Args["limit"]
					if reflect.TypeOf(limit).Kind() == reflect.Int {
						opts = opts.SetLimit(int64(limit.(int)))
					}
					page := p.Args["page"]
					database := client.Database("database", options.Database())
					collection := database.Collection("forms", options.Collection())
					filter := bson.D{}
					results, err := collection.Find(p.Context, filter)
					return nil, nil
				},
			},
		},
	})
)
