package models

import "github.com/globalsign/mgo/bson"

type Userinfo struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	User_id string `bson:"user_id" json:"user_id"`
	Name string `bson:"name" json:"name"`
	Avatar string	`bson:"avatar" json:"avatar"`
	Birth_at string `bson:"birth_at" json:"birth_at"`
	Sex string `bson:"sex" json:"sex"`
	Auth_ids []int `bson:"auth_ids" json:"auth_ids"`
	Created_at string `bson:"created_at" json:"created_at"`
	Updated_at string `bson:"updated_at" json:"updated_at"`
	Deleted_at string `bson:"deleted_at" json:"deleted_at"`
}

