package controllers

import (
	"../db"
	"../models"

	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(c echo.Context) (err error) {
	userCollection, session, err := getUserCollection()
	if err != nil {
		return
	}
	defer session.Close()
	var User models.User
	err = userCollection.FindId(bson.ObjectIdHex(c.Param("id"))).One(&User)
	if err != nil {
		return
	}
	fmt.Print(User)
	return c.JSON(http.StatusOK, User)
}

func getUserCollection() (userCollection *mgo.Collection, session *mgo.Session, err error) {
	session, err = db.GetSession()
	database := session.DB(db.GetDBName())
	if err != nil {
		return
	}
	userCollection = database.C("users")
	return
}
