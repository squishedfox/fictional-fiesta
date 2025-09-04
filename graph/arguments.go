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
		Name: "FieldSetInputInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"label": &LabelInputConfig,
			"type": &graphql.InputObjectFieldConfig{
				Type:         InputTypeEnum,
				DefaultValue: "text",
				Description:  "The type of input the user will see such as phone, email, calendar, select, text, etc.",
			},
			"minLength": &graphql.InputObjectFieldConfig{
				Type:         graphql.Int,
				DefaultValue: nil,
				Description:  "The minimum number of characters required for the field. See https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#maxlength for valid input types",
			},
			"maxLength": &graphql.InputObjectFieldConfig{
				Type:         graphql.Int,
				DefaultValue: nil,
				Description:  "The maximum number of characters required for the field. See https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#maxlength for valid input types",
			},
			"min": &graphql.InputObjectFieldConfig{
				Type:         graphql.String,
				DefaultValue: nil,
				Description:  "Minimum value that can be applied to the input. See https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/min for valid input types",
			},
			"max": &graphql.InputObjectFieldConfig{
				Type:         graphql.String,
				DefaultValue: nil,
				Description:  "Maximum value that can be applied to the input. See https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/max for valid input types",
			},
			"required": &graphql.InputObjectFieldConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
				Description:  "True if the field is required to be filled out, else false",
			},
			"list": &graphql.InputObjectFieldConfig{
				Type:         graphql.NewList(graphql.String),
				Description:  "List of selectable options. See https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#list for valid input types",
				DefaultValue: []string{},
			},
			"multiple": &graphql.InputObjectFieldConfig{
				Type:         graphql.Boolean,
				Description:  "Flag indicating that a user can select multiple values for a input field. See https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/multiple for valid input types",
				DefaultValue: false,
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
				Name: "FieldSetInput",
				Fields: graphql.InputObjectConfigFieldMap{
					"label": &LabelInputConfig,
					"inputs": &graphql.InputObjectFieldConfig{
						Type:         graphql.NewList(FieldSetInputInputConfig),
						Description:  "Input Field Configurations",
						DefaultValue: nil,
					},
				},
			}),
			Description: "Field set to naturally group a number of input fields such as physical mailing address",
		},
	}
)
