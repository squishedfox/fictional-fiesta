package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	CreateFormMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CreateFormMutation",
		Description: "Create a brand new form",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
)
