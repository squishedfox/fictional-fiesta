package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/squishedfox/fictional-fiesta/graph"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graph.FormQueries,
		Mutation: graph.CreateFormMutation,
	})
	if err != nil {
		log.Fatal(err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	s := http.NewServeMux()
	s.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"status\": \"OK\"}"))
	}))
	s.Handle("/graphql", h)

	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatal(err)
	}
}
