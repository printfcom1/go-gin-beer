package database

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDataBaseMongoDB() *mongo.Database {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		panic(err.Error())
	}
	url := envMap["MONGODB_URL"]
	dbName := envMap["MONGODB_DB"]
	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	mongoDB := client.Database(dbName)
	fmt.Println("Connected to MongoDB!")
	return mongoDB
}
