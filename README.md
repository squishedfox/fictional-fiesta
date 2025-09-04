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
2. run `echo "MONGO_DB_URL=mongodb://root:example@0.0.0.0:27017/\n
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

## Examples

All examples use the localhost environment for testing. They demonstrate variables, queries, and mutations.

### Basic Fetching

#### Query

```graphql
query filterForms($id: String!, $page: Int, $limit: Int) {
  list(id: $id, page: $page, limit: $limit) {
    count,
    results {
      id,
      fieldsets {
        fields {
          ordinal,
          name,
          type,
        }
        legend,
      }
    }
  }
}
```

#### Variables

```json
{
  "id": "68b99132f30dfb2ed5aed633"
}
```

### Creating a New Form

#### mutation

```graphql
mutation createForm($name: String!, $active: Boolean!, $fieldsets: FieldSetInput) {
  create(name: $name, active: $active, fieldset: $fieldsets) {
      id
  }
}
```

#### Variables

```json
{
  "active": true,
  "fieldsets": {
    "inputs": [
      {
        "label": "First Name",
        "multiple": false,
        "required": true,
        "type": "text"
      },
      {
        "label": "Last Name",
        "multiple": false,
        "required": true,
        "type": "text"
      }
    ],
    "label": "Basic Contact"
  },
  "name": "A name that you want this form to be"
}
```
