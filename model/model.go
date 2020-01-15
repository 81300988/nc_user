package mdel

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"FirstName"`
	LastName  string `json:"last_name" bson:"LastName"`
	Age       int    `json:"age" bson:"Age"`
	Email     string `json:"email" bson:"Email"`
	Phone     string `json:"phone" bson:"Phone"`
	PassWord  string `json:"password" bson:"PassWord"`
}
type UserLoggin struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"FirstName"`
	LastName  string `json:"last_name" bson:"LastName"`
	Age       int    `json:"age" bson:"Age"`
	Email     string `json:"email" bson:"Email"`
	Phone     string `json:"phone" bson:"Phone"`
	PassWord  string `json:"password" bson:"PassWord"`
	Token     string `json:"token" bson:"Token"`
}

type UserUpdateReq struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"FirstName"`
	LastName  string `json:"last_name" bson:"LastName"`
	PassWord  string `json:"password" bson:"PassWord"`
}

type UserClaims struct {
	UserID int    `json: "uid"`
	Phone  string `json: "p"`
	Email  string `json: "e"`
	jwt.StandardClaims
}
type Error struct {
	Code int
	Msg  string
}

type LoginReq struct {
	Email    string `json:"email" bson:"email"`
	PassWord string `json:"password" bson:"password"`
}
type LoginRes struct {
	*User
	Token string `json:"token"`
}
