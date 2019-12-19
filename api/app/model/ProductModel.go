package model

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()";json:"-"`
	Name        string    `json:"name"`
	EnName      string    `json:"enname"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"categoryID"`
}
