package model

import "time"

type Event struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	Sessions   []time.Time `json:"sessions"`
	Place      string      `json:"place"`
	Tags       []string    `json:"tags"`
	Interested int64       `json:"interested"`
}
