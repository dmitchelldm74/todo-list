package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID uint `json: "account_id" gorm:"primary_key;autoIncrement:1"`
}
