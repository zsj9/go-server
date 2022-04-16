package models

import "github.com/globalsign/mgo/bson"

type User struct {
	Id bson.ObjectId `bson:"_id"`
	UserName string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Token string	`bson:"token"`
	Created_at string `bson:"created_at" json:"created_at"`
	Updated_at string `bson:"updated_at" json:"updated_at"`
	Deleted_at string `bson:"deleted_at" json:"deleted_at"`
}

type JwtToken struct {
	Token string `json:"token"`
}
