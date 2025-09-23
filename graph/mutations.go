package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	CreateFormMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CreateFormMutation",
		Description: "Create a brand new form",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Args:        CreateNewFormArgument,
				Description: "Create a new form",
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "CreateNewFormResponse",
					Fields: graphql.Fields{
						"id": &graphql.Field{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
				}),
				Resolve: createFormResolver,
			},
		},
	})
)
