package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/squishedfox/fictional-fiesta/graph"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	connUrl := os.Getenv("MONGO_DB_URL")
	if len(connUrl) == 0 {
		log.Fatal("Could not find environment variable MONGO_DB_URL")
	}

	opts := options.Client().ApplyURI(connUrl)
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	//
	// setup schema
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

	//
	// setup database
	s := http.NewServeMux()
	s.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"status\": \"OK\"}"))
	}))
	s.Handle("/graphql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contextWithMongoConnection := context.WithValue(r.Context(), graph.ContextKey, client)
		h.ContextHandler(contextWithMongoConnection, w, r)
	}))

	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatal(err)
	}
}
