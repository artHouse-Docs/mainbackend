package hashing

import (
	"github.com/artHouse-Docs/backend/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var salt = config.Configure().Server.Salt

func MakeHash(password string) (hash string) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashBytes)
}

func CompareHash(password string, hash string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}
