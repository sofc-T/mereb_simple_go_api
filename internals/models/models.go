package models

import "github.com/google/uuid"

type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name" binding:"required"`
	Age     int      `json:"age" binding:"required"`
	Hobbies []string `json:"hobbies" binding:"required"`
}

func NewPerson(name string, age int, hobbies []string) *Person {
	return &Person{
		ID:      uuid.NewString(),
		Name:    name,
		Age:     age,
		Hobbies: hobbies,
	}
}
