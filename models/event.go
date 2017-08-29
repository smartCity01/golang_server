package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title,omitempty"`
	Host        string        `json:"host" bson:"host,omitempty"`
	HostName    string        `json:"hostName" bson:"hostName,omitempty"`
	Location    location      `json:"location" bson:"location,omitempty"`
	Description string        `json:"description" bson:"description,omitempty"`
	ImageURL    string        `json:"imageUrl" bson:"imageUrl,omitempty"`
	Venue       string        `json:"venue" bson:"venue,omitempty"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
