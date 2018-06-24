package controllers

import (
	"gopkg.in/mgo.v2"
)

func NewInitController() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	var ci *mgo.CollectionInfo

	if err1 := s.DB("go_mongo").C("users").Create(ci); err1 != nil {
		panic(err1)
	}
	if err2 := s.DB("go_mongo").C("roles").Create(ci); err2 != nil {
		panic(err2)
	}
	if err3 := s.DB("go_mongo").C("transections").Create(ci); err3 != nil {
		panic(err3)
	}
	return s
}
