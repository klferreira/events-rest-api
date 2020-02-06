package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Event struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name,omitempty"`
	Sessions   []time.Time   `json:"sessions" bson:"sessions,omitempty"`
	Place      string        `json:"place" bson:"place,omitempty"`
	Tags       []string      `json:"tags" bson:"tags,omitempty"`
	Interested int64         `json:"interested" bson:"interested,omitempty"`
}
