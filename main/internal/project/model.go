package project

import "github.com/artHouse-Docs/backend/internal/user"

type (
	ProjectId   = string
	ProjectName = string
)

type CommandMember struct {
	Member user.User `json:"-" bson:"member"`
	Role   string    `json:"role" bson:"role"`
}

type Project struct {
	ID        ProjectId       `json:"id" bson:"_id"`
	Name      ProjectName     `json:"name" bson:"name"`
	Customers any             `json:"-" bson:"customers" `
	Commands  []CommandMember `json:"-" bson:"commands"`
}
