package dtos

import "github.com/google/uuid"

type ResponseUser struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
