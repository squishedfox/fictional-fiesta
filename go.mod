module github.com/squishedfox/fictional-fiesta

go 1.24.3

require (
	github.com/graphql-go/graphql v0.8.1
	github.com/graphql-go/handler v0.2.4
	github.com/joho/godotenv v1.5.1
	github.com/squishedfox/fictional-fiesta/db v0.0.0-20250717183151-a8de32585fe5
	github.com/squishedfox/fictional-fiesta/db/mongodb v0.0.0-20250717183151-a8de32585fe5
	github.com/squishedfox/fictional-fiesta/graph v0.0.0-20250717024829-65b24ed3a148
	go.mongodb.org/mongo-driver/v2 v2.2.2
)

require (
	github.com/golang/snappy v1.0.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)

replace github.com/squishedfox/fictional-fiesta/graph => ./graph

replace github.com/squishedfox/fictional-fiesta/db => ./db

replace github.com/squishedfox/fictional-fiesta/db/mongodb => ./db/mongodb
