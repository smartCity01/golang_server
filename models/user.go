package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserName       string        `json:"username" bson:"username,omitempty"`
	HashedPassword string        `json:"hashedPassword" bson:"hashedPassword,omitempty"`
	Salt           string        `json:"salt" bson:"salt,omitempty"`
}
