package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	LabelInputConfig = graphql.InputObjectFieldConfig{
		Type:         graphql.String,
		DefaultValue: "",
		Description:  "The display text that will be tagged with a label",
	}
	FieldSetInputInputConfig = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "FieldSetInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"label": &LabelInputConfig,
			"type": &graphql.InputObjectFieldConfig{
				Type:         InputTypeEnum,
				DefaultValue: "text",
				Description:  "The type of input the user will see such as phone, email, calendar, select, text, etc.",
			},
			"required": &graphql.InputObjectFieldConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
				Description:  "True if the field is required to be filled out, else false",
			},
		},
	})
	CreateNewFormArgument = graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A unique human readable name for the form",
		},
		"active": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "Whether the form is active or usable",
		},
		"fieldset": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name: "FieldSetArguments",
				Fields: graphql.InputObjectConfigFieldMap{
					"label": &LabelInputConfig,
					"inputs": &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(FieldSetInputInputConfig),
					},
				},
			}),
			Description: "Field set to naturally group a number of input fields such as physical mailing address",
		},
	}
)
