package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	FormObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Form",
		Description: "Dynamic Form with groups or fields",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
)
