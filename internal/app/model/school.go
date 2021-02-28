package model

import uuid "github.com/satori/go.uuid"

type School struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	City string    `json:"city"`
}
