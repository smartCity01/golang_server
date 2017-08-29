package controllers

import (
	"fmt"
	"net/http"

	"../db"
	"../models"
	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
)

func GetEvents(c echo.Context) (err error) {
	session, err := db.GetSession()
	database := session.DB("joinus_profiles")
	defer session.Close()
	if err != nil {
		return
	}

	eventCollection := database.C("events")
	var Events []models.Event
	err = eventCollection.Find(bson.M{}).All(&Events)
	fmt.Print(Events)
	return c.JSON(http.StatusOK, Events)
}
