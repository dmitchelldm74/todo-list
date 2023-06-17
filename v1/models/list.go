package models

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	TodoListId   uint `gorm:"primary_key;autoIncrement"`
	AccountId    uint
	TodoListName string
}
