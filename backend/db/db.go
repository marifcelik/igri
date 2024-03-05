package db

import (
	"context"
	"time"

	"go-chat/config"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.C[config.MongoURI]))
	if err != nil {
		log.Fatal(err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	if err = client.Ping(pingCtx, nil); err != nil {
		log.Fatal(err)
	}

	DB = client.Database("go-chat")
}
