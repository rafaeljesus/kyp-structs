package structs

import (
	"time"
)

type Todo struct {
	Id        uint       `json:"id", sql:"primary_key"`
	Title     string     `json:"title", sql:"not null"`
	Done      bool       `json:"done"`
	UserId    uint       `json:"user_id, sql:"not null`
	CreatedAt time.Time  `json:"created_at", sql:"not null"`
	UpdatedAt time.Time  `json:"updated_at", sql:"not null`
	DeletedAt *time.Time `json:"-" "created_at"`
}
