# fictional-fiesta

## Summary

A proof of concept HTML Form Builder created using Golang, GraphQL, and MongoDB

## Build With

- [graphql-go/graphql](https://github.com/graphql-go/graphql)
- [graphql-go/handlers](https://github.com/graphql-go/handlers)
- [Golang (Go)](https://go.dev)
- [MongoDB](https://github.com/graphql-go/graphql)
- [GraphQL](https://graphql.com/)

## Getting Started

### Setup MongoDB Database

1. make sure you have docker installed
2. run `echo "MONGO_DB_URL=mongodb://root:example@0.0.0.0:27017/\
    DATABASE_NAME=forms" > .env`
3. run `docker compose up -d` to start mongodb server locally

### Running the Project

1. run `go get` to install modules
2. run `go run main.go` to start the project
3. run `curl http://localhost:8080/`
4. You should received

```json
{ "status": "ok" }
```

