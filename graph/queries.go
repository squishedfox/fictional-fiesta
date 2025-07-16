package graph

import "github.com/graphql-go/graphql"

var (
	FormQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Form Query",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name:        "ID",
				Type:        graphql.ID,
				Description: "The unique Identifier of the Form itself",
			},
			"name": &graphql.Field{
				Name:        "Name",
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The unique name given to the form that is readable to the user",
			},
		},
	})
)
