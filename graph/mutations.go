package graph

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

var (
	CreateNewFormArguments = graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A unique human readable name for the form",
		},
	}
	CreateFormMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CreateFormMutation",
		Description: "Create a brand new form",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Args:        CreateNewFormArguments,
				Description: "Create a new form",
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "CreateNewformResponseType",
					Fields: graphql.Fields{
						"id": &graphql.Field{
							Name: "IDType",
							Type: graphql.String,
						},
					},
				}),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					repository := p.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
					if repository == nil {
						return nil, errors.New("Could not fetch repository from user context")
					}
					model := &db.CreateFormModel{
						Name: p.Args["name"].(string),
					}
					id, err := repository.CreateForm(model)
					return struct {
						ID any `json:"id"`
					}{
						id,
					}, err
				},
			},
		},
	})
)
