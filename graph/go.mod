module github.com/squishedfox/fictional-fiesta/graph

go 1.24.3

require (
	github.com/graphql-go/graphql v0.8.1
	github.com/squishedfox/fictional-fiesta/db v0.0.0-20250717024829-65b24ed3a148
	gotest.tools/v3 v3.5.2
)

require github.com/google/go-cmp v0.6.0 // indirect

replace github.com/squishedfox/fictional-fiesta/db => ../db
