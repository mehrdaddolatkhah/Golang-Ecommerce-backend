package model

import "github.com/google/uuid"

type Cart struct {
	ID        uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"userid"`
	ProductID uuid.UUID `json:"productid"`
	Quantity  string    `json:"quantity"`
	Price     string    `json:"price"`
}
