package user

import (
	"context"
	"errors"

	db "github.com/artHouse-Docs/backend/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserCollection(ctx context.Context) (coll *mongo.Collection, err error) {
	coll, err = db.NewCollection(ctx, "users")
	if err != nil {
		return nil, errors.New("database unavaliable")
	}

	return coll, nil
}
