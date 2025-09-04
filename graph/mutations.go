package graph

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
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
				Resolve: func(p graphql.ResolveParams) (any, error) {
					name := p.Args["name"].(string)
					repository := p.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
					if repository == nil {
						return nil, errors.New("Could not fetch repository from user context")
					}

					id, err := repository.CreateForm(&db.CreateFormModel{
						Name: name,
					})
					if err != nil {
						return nil, err
					}

					return &struct {
						ID any `json:"id"`
					}{id}, err
				},
			},
		},
	})
)
