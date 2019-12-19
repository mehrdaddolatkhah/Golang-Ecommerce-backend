package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	Email    string    `db:"unique;not null";json:"email"`
	Mobile   string    `db:"unique;json:"mobile"`
	BirthDay string    `json:"birthday"`
	Identify string    `json:"identify"`
	Cart     string    `json:"cart"`
	Credit   string    `json:"credit"`
	Type     string    `json:"type"`
	Password string    `json:"password"`
}
