package graph

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
)

func getFormsRepository(params *graphql.ResolveParams) (db.FormsRepository, error) {
	formsRepository, ok := params.Context.Value(db.FormsRepositoryContextKey).(db.FormsRepository)
	if !ok {
		return nil, errors.New("Could not fetch repository from user context")
	}
	return formsRepository, nil
}
