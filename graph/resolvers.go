package graph

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

func formListResolver(p graphql.ResolveParams) (any, error) {
	repository := p.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
	if repository == nil {
		return nil, errors.New("Could not fetch repository from user context")
	}

	limit := getIntOrDefault(p.Args["limit"], DEFAULT_LIMIT_VALUE)
	page := getIntOrDefault(p.Args["page"], DEFAULT_PAGE_VALUE)
	skip := limit * page

	result, err := repository.GetForms(&db.GetFormsModel{
		Limit: int64(limit),
		Skip:  int64(skip),
	})
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
