package structs

import (
	"time"
)

type User struct {
	Id                uint       `json:"id", sql:"primary_key"`
	Email             string     `json:"email" valid:"email"`
	Password          string     `json:"password,omitempty"`
	EncryptedPassword []byte     `json:"-" sql:"encrypted_password;not null"`
	CreatedAt         time.Time  `json:"created_at", sql:"not null"`
	UpdatedAt         time.Time  `json:"updated_at", sql:"not null`
	DeletedAt         *time.Time `json:"-" "created_at"`
}
