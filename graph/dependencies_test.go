package graph

import (
	"context"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/fictional-fiesta/db"
	"gotest.tools/v3/assert"
)

func TestGetformsRepository(t *testing.T) {

	getResolverParams := func(repo db.FormsRepository) graphql.ResolveParams {
		emptyContext := context.TODO()
		if repo != nil {
			emptyContext = context.WithValue(emptyContext, db.FormsRepositoryContextKey, repo)
		}
		return graphql.ResolveParams{
			Context: emptyContext,
			Source:  &struct{}{},
			Args:    make(map[string]any),
			Info:    graphql.ResolveInfo{},
		}
	}

	t.Run("getFormsRepository should return error when not present in context", func(tt *testing.T) {
		// arrange
		params := getResolverParams(nil)

		// act
		repo, err := getFormsRepository(&params)

		// assert
		assert.Assert(tt, repo == nil, "Expect Repo to be nil but got %v", repo)
		assert.Assert(tt, err != nil, "Expected Error to have value but got nil instead")
	})
}
