package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	id       bson.ObjectId `json:"id" bson:"_id"`
	username string        `json:"username" bson:"username"`
	fname    string        `json:"fname" bson:"fname"`
	lname    string        `json:"lname" bson:"lname"`
	role     string        `json:"role" bson:"role"`
	password []byte        `json:"password" bson:"password"`
}

type Session struct {
	username string
	lastact  time.Time
}
