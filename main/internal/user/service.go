package user

import (
	"context"
	"errors"
	"github.com/artHouse-Docs/backend/pkg/dto"
	"github.com/artHouse-Docs/backend/pkg/grpcBridge"
	"google.golang.org/grpc"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/artHouse-Docs/backend/pkg/hashing"
)

func (u *User) Register(ctx context.Context) (err error) {
	coll, err := NewUserCollection(ctx)
	if err != nil {
		return errors.New("database unavailable")
	}

	result, err := coll.InsertOne(ctx, bson.D{
		{"name", u.Name},
		{"surname", u.Surname},
		{"password", hashing.MakeHash(u.PasswordHash)},
		{"email", u.Email},
	})

	if err != nil {
		return errors.New("user email already exists")
	}

	u.ID = result.InsertedID.(primitive.ObjectID).Hex()

	cc := grpcBridge.CreateConnection()
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal("Unexpected connection error\t", err.Error())
		}
	}(cc)

	client := dto.NewAuthServiceClient(cc)

	tokens, err := client.Login(context.Background(), &dto.Payload{Id: u.ID})
	if err != nil {
		log.Fatal("Unexpected token error\t", err.Error())
	}

	u.AccessToken = tokens.AccessToken
	u.RefreshToken = tokens.RefreshToken

	return nil
}

func (u *User) Login(ctx context.Context) (result bool, err error) {
	coll, err := NewUserCollection(ctx)
	if err != nil {
		return false, errors.New("database unavailable")
	}

	var userByEmail User

	coll.FindOne(ctx, bson.D{
		{"email", u.Email},
	}).Decode(&userByEmail)

	if userByEmail.ID == "" {
		return false, errors.New("user not found")
	} else if !hashing.CompareHash(u.PasswordHash, userByEmail.PasswordHash) {
		return false, errors.New("wrong password")
	} else {
		u.ID = userByEmail.ID
		u.Name = userByEmail.Name
		u.Surname = userByEmail.Surname
	}
	return true, nil
}
