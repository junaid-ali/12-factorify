package dblayer

import (
	"fmt"
	"eventservice/lib/persistence"
	"eventservice/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB		DBTYPE = "mongodb"
	DYNAMODB	DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		fmt.Println("mongo connection")
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
