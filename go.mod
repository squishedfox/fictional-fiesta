module github.com/squishedfox/fictional-fiesta

go 1.24.3

replace github.com/squishedfox/fictional-fiesta/graph => ./graph

require (
	github.com/graphql-go/graphql v0.8.1
	github.com/graphql-go/handler v0.2.4
	github.com/squishedfox/fictional-fiesta/graph v0.0.0-20250716153411-243c3f1b2077
)
