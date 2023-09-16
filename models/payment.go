package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Paymentid *string `json:"paymentid" gorm:"primary_key"`
	Userid    *string `json:"userid"`
	Lastname  *string `json:"lastname"`
	Firstname *string `json:"firstname"`
}
