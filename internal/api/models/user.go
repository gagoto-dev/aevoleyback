package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `json="id"`
	Name      string    `json="name"`
	Password  string    `json="password"`
	Email     string    `json="email"`
	CreatedAt time.Time `json="createdAt"`
}

func NewUser(name, password, email string) *User {
	return &User{
		Id:        uuid.New(),
		Name:      name,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
	}
}
