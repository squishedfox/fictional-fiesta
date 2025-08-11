package graph

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

var (
	FormQueries = graphql.NewObject(graphql.ObjectConfig{
		Name: "FormsQueries",
		Fields: graphql.Fields{
			"forms": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "FormList",
					Fields: graphql.Fields{
						"list": &graphql.Field{
							Type: graphql.NewList(FormObject),
						},
						"count": &graphql.Field{
							Type: graphql.Int,
						},
					},
				}),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						DefaultValue: "",
						Type:         graphql.String,
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
					repository := p.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
					if repository == nil {
						return nil, errors.New("Could not fetch repository from user context")
					}
					result, err := repository.GetForms(&db.GetFormsModel{})
					if err != nil {
						return nil, err
					}
					return &struct {
						List  []*db.FormModel `json:"list"`
						Count int64           `json:"count"`
					}{result.Forms, result.Count}, nil
				},
			},
		},
	})
)
