package database

import (
	"context"

	"github.com/artHouse-Docs/backend/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	host     string
	port     string
	username string
	password string
}

func (database *Database) NewClient(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+database.username+":"+database.password+"@"+database.host+":"+database.port))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database("main"), err
}

func NewCollection(ctx context.Context, name string) (coll *mongo.Collection, err error) {
	cfg := config.Configure().Database

	database := Database{cfg.Host, cfg.Port, cfg.Username, cfg.Password}
	client, err := database.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return client.Collection(name), nil
}
