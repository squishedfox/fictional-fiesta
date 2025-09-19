package graph

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

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
	repository := p.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
	if repository == nil {
		return nil, errors.New("Could not fetch repository from user context")
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
}
