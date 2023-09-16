package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Paymentid uuid.UUID `json:"paymentid" gorm:"primary_key"`
	Userid    string
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}
