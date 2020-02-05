package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Event struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Sessions   []time.Time   `json:"sessions" bson:"sessions"`
	Place      string        `json:"place" bson:"place"`
	Tags       []string      `json:"tags" bson:"tags"`
	Interested int64         `json:"interested" bson:"interested"`
}
