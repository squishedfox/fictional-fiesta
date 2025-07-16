package graph

import "github.com/graphql-go/graphql"

var (
	FormQueries = graphql.NewObject(graphql.ObjectConfig{
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
			},
		},
	})
)
