package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sinsombat/practice/controllers"
	"gopkg.in/mgo.v2"
)

var s *mgo.Session

func init() {
	s = controllers.NewInitController()
}

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(s)
	r.GET("/", uc.GetUsers)
}
