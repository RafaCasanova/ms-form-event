package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv("MONGODB_URL")
	mongodb_port := os.Getenv("MONGODB_PORT")
	mongodb_user := os.Getenv("MONGODB_USER_DB")
	mongodb_pass := os.Getenv("MONGODB_PASS_DB")
	mongodb_database := os.Getenv("MONGODB_DATA_BASE")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongodb_user, mongodb_pass, mongodb_uri, mongodb_port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
