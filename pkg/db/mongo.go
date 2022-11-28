package db

import (
	"context"
	"fmt"
	"os"

	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MustGetClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	godotenv.Load()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("db connection established")
	return client
}

// for creating migrations later on
func CreateUniqueIndex(collection *mongo.Collection, name string) (string, error) {
	IndexName, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: name, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return "", err
	}
	return IndexName, nil
}

func DropUniqueIndex(collection *mongo.Collection, name string) (bson.Raw, error) {
	IndexName, err := collection.Indexes().DropOne(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return IndexName, nil
}
