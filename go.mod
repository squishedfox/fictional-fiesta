module github.com/squishedfox/fictional-fiesta

go 1.24.3

require (
	github.com/graphql-go/graphql v0.8.1
	github.com/graphql-go/handler v0.2.4
	github.com/jackc/pgx/v5 v5.7.5
	github.com/squishedfox/fictional-fiesta/graph v0.0.0-20250716153411-243c3f1b2077
)

require (
	github.com/golang/snappy v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.mongodb.org/mongo-driver/v2 v2.2.2 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)

replace github.com/squishedfox/fictional-fiesta/graph => ./graph

replace github.com/squishedfox/fictional-fiesta/db => ./db

replace github.com/squishedfox/fictional-fiesta/db/mongodb => ./db/mongodb
