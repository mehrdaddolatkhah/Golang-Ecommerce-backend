package model

import "github.com/google/uuid"

type Category struct {
	Id     uuid.UUID `sql:"unique;type:uuid;default:uuid_generate_v4()";json:"id"`
	Parent uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()";json:"parent";json:"parent"`
	Title  string    `json:"title"`
	Level  string    `json:"level"`
}
