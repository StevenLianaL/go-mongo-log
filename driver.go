package mongo_log

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(dbName string, collectionName string, dbUser string, dbPswd string,
	dbHost string, dbPort int,
	ctx *context.Context) *mongo.Collection {

	dbUrl := "mongodb://" + dbUser + ":" + dbPswd + "@" + dbHost + fmt.Sprintf(":%d/", dbPort) + dbName + "?authSource=admin"
	client, _ := mongo.Connect(*ctx, options.Client().ApplyURI(dbUrl))
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
