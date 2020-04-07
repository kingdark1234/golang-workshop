package entity

import (
	"time"
)

// People ...
type People struct { // Table People
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:255;not null;unique" json:"firstName"`
	LastName  string    `gorm:"size:255;not null;unique" json:"lastName"`
	Age       int       `gorm:"size:100;not null;" json:"age"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Type ...
type AddPeople struct {
	firstName string
	lastName  string
	age       int
}

type updatePeople struct {
	id        string
	firstName string
	lastName  string
	age       int
}

type deleteAndGetPeople struct {
	id string
}

type getAllPeoples []struct{ *AddPeople }
