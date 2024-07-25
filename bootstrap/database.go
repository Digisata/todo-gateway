package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/digisata/auth-service/pkg/mongo"
)

func NewMongoDatabase(cfg *Config) (mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := cfg.Mongo.DBHost
	dbPort := cfg.Mongo.DBPort
	dbUser := cfg.Mongo.DBUser
	dbPass := cfg.Mongo.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%d", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
