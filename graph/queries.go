package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	FormQueries = graphql.NewObject(graphql.ObjectConfig{
		Name: "FormsQueries",
		Fields: graphql.Fields{
			"list": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "FormList",
					Fields: graphql.Fields{
						"results": &graphql.Field{
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
				Resolve: formListResolver,
			},
		},
	})
)
