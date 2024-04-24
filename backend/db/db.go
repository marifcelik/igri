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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.C.MongoURI))
	if err != nil {
		log.Fatal("db connection error", "err", err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	if err = client.Ping(pingCtx, nil); err != nil {
		log.Fatal("db ping error", "err", err)
	}

	DB = client.Database("go-chat")
}
