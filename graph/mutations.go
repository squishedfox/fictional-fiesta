package graph

import (
	"github.com/graphql-go/graphql"
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
				Type:        FormObject,
			},
		},
	})
)
