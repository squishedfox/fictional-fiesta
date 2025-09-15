package graph

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

func mapFieldSetInput(raw map[string]any) (db.FormInputModel, error) {
	inputModel := db.FormInputModel{}

	for key, value := range raw {
		switch key {
		case "label":
			break
		case "type":
			inputModel.Type = getStringOrDefault(value, "")
			break
		case "minLength":
			inputModel.MinLength = getIntPointerOrDefault(value, nil)
			break
		case "maxLength":
			inputModel.MaxLength = getIntPointerOrDefault(value, nil)
			break
		case "min":
			inputModel.Min = getStringOrDefault(value, "")
			break
		case "max":
			inputModel.Max = getStringOrDefault(value, "")
			break
		case "required":
			inputModel.Required = getBoolOrDefault(value, false)
			break
		case "list":
			fmt.Println("list not implemented")
			break
		case "multiple":
			inputModel.Multiple = getBoolOrDefault(value, false)
			break
		default:
			break
		}
	}

	return inputModel, nil
}

func convertArgsToFieldsets(arg []any) ([]db.FieldSetModel, error) {

	fieldsets := []db.FieldSetModel{}
	for _, rawFieldset := range arg {
		m, ok := rawFieldset.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("Error trying to convert fieldset %v to be a map[string]interface{}", rawFieldset)
		}

		rawInputs, ok := m["inputs"].([]any)
		if !ok {
			return nil, fmt.Errorf("Error trying to convert input %v to be a map[string]interface{}", m["inputs"])
		}
		inputList := []db.FormInputModel{}
		for _, rawInput := range rawInputs {
			inputMap, ok := rawInput.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("Error trying to convert input %v to be a map[string]interface{}", m["inputs"])
			}
			input, err := mapFieldSetInput(inputMap)
			if err != nil {
				return nil, err
			}
			inputList = append(inputList, input)
		}

		legend, ok := m["legend"].(string)
		if !ok {
			return nil, fmt.Errorf("Error trying to convert legend %v to be string", m["legened"])

		}
		fieldsets = append(fieldsets, db.FieldSetModel{
			Legend: legend,
			Inputs: inputList,
		})
	}
	return fieldsets, nil
}

var (
	CreateFormMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CreateFormMutation",
		Description: "Create a brand new form",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Args:        CreateNewFormArgument,
				Description: "Create a new form",
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "CreateNewFormResponse",
					Fields: graphql.Fields{
						"id": &graphql.Field{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
				}),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					name := p.Args["name"].(string)
					repository := p.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
					if repository == nil {
						return nil, errors.New("Could not fetch repository from user context")
					}

					fieldsetMap, ok := p.Args["fieldsets"].([]any)
					if !ok {
						return nil, errors.New("Invalid type for fieldsets")
					}
					fieldsets, err := convertArgsToFieldsets(fieldsetMap)
					if err != nil {
						return nil, err
					}

					id, err := repository.CreateForm(&db.CreateFormModel{
						Name:      name,
						Fieldsets: fieldsets,
					})
					if err != nil {
						return nil, err
					}

					return &struct {
						ID any `json:"id"`
					}{id}, err
				},
			},
		},
	})
)
