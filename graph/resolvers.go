package graph

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

func mapFieldSetInput(raw map[string]any) (db.FormInputModel, error) {
	inputModel := db.FormInputModel{}

	for key, value := range raw {
		switch key {
		case "label":
			inputModel.Label = ""
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

func convertArgsToFieldsets(arg any) ([]db.FieldSetModel, error) {

	fieldsets := []db.FieldSetModel{}
	if arg == nil {
		return fieldsets, nil
	}

	fieldsetList, ok := arg.([]any)
	if !ok {
		return fieldsets, fmt.Errorf("arg is of wrong type %s", reflect.TypeOf(arg))
	}

	for _, rawFieldset := range fieldsetList {
		m, ok := rawFieldset.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("Error trying to convert fieldset %v to be a map[string]interface{}", rawFieldset)
		}

		rawInputs, ok := m["inputs"].([]any)
		if !ok {
			return nil, fmt.Errorf("Error trying to convert input %v to be an array type of any", m["inputs"])
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

func getFormsModel(args map[string]any) db.GetFormsModel {
	limit := getIntOrDefault(args["limit"], DEFAULT_LIMIT_VALUE)
	page := getIntOrDefault(args["page"], DEFAULT_PAGE_VALUE)
	skip := limit * page

	filters := []*db.SearchFieldFilter{}
	if name, ok := args["name"]; ok {
		filters = append(filters, &db.SearchFieldFilter{
			Operation: db.EqualsOperation,
			Value:     name,
			Field:     "name",
		})
	}
	return db.GetFormsModel{
		Limit:   int64(limit),
		Skip:    int64(skip),
		Filters: filters,
	}
}

func formListResolver(p graphql.ResolveParams) (any, error) {
	repository, err := getFormsRepository(&p)
	if err != nil {
		return nil, err
	}
	model := getFormsModel(p.Args)
	result, err := repository.GetForms(&model)
	if err != nil {
		return nil, err
	}
	return &struct {
		Results []*db.FormModel `json:"results"`
		Count   int64           `json:"count"`
	}{
		result.Forms,
		result.Count,
	}, nil
}

func createFormResolver(p graphql.ResolveParams) (any, error) {
	name, ok := p.Args["name"].(string)
	if !ok {
		return nil, errors.New("name is required")
	}

	repository, err := getFormsRepository(&p)
	if err != nil {
		return nil, err
	}

	fieldsets, err := convertArgsToFieldsets(p.Args["fieldsets"])
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
}
