package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sinsombat/practice/models"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := models.User{}

	if err := uc.session.DB("go_practice").C("users").Find(nil).All(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
