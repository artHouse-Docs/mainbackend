package user

type (
	UserId   = string
	UserName = string
)

type User struct {
	ID           UserId   `json:"id" bson:"_id"`
	Name         UserName `json:"name" bson:"name"`
	Surname      string   `json:"surname" bson:"surname"`
	PasswordHash string   `json:"password" bson:"password"`
	Email        string   `json:"email" bson:"email"`
	AccessToken  string   `json:"-" bson:"-"`
	RefreshToken string   `json:"-" bson:"-"`
}
