package mongodb

import "go.mongodb.org/mongo-driver/v2/mongo"

type (
	Repository struct {
		conn *mongo.Session
	}
)
