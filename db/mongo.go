package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func GetSession() (session *mgo.Session, err error) {
	session, err = mgo.Dial("mongodb://testuser2:test@ds032319.mlab.com:32319/joinus_profiles")
	if err != nil {
		fmt.Print("error connecting to db")
		return
	}
	return
}

func GetDBName() string {
	return "joinus_profiles"
}
