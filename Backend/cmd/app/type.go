package app

import "math/rand"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{Name: name, ID: rand.Intn(100)}
}
