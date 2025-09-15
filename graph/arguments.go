package graph

import (
	"github.com/graphql-go/graphql"
)

func getIntOrDefault(raw any, defaultValue int) int {
	val, ok := raw.(int)
	if !ok {
		return defaultValue
	}

	return val
}

func getIntPointerOrDefault(raw any, defaultValue *int) *int {
	intVal, ok := raw.(int)
	if !ok {
		ptrInt, ok := raw.(*int)
		if !ok {
			return defaultValue
		}
		return ptrInt
	}

	return &intVal
}

func getStringOrDefault(raw any, defaultValue string) string {
	val, ok := raw.(string)
	if !ok {
		return defaultValue
	}

	return val
}

func getBoolOrDefault(raw any, defaultValue bool) bool {
	val, ok := raw.(bool)
	if !ok {
		return defaultValue
	}

	return val
}

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
		"fieldsets": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.NewInputObject(graphql.InputObjectConfig{
				Name:        "FieldSetInput",
				Description: "Field set to naturally group a number of input fields such as physical mailing address",
				Fields: graphql.InputObjectConfigFieldMap{
					"legend": &graphql.InputObjectFieldConfig{
						Type:         graphql.String,
						DefaultValue: "Provides a description for the form input grouping",
					},
					"inputs": &graphql.InputObjectFieldConfig{
						Type:         graphql.NewList(FieldSetInputInputConfig),
						Description:  "Input Field Configurations",
						DefaultValue: nil,
					},
				},
			})),
		},
	}
)
