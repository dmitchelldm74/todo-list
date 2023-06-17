package models

import (
	"gorm.io/gorm"
)

type TodoListItem struct {
	gorm.Model
	TodoListItemId      uint `gorm:"primary_key;autoIncrement"`
	TodoListId          uint
	TodoListItemText    string
	TodoListItemChecked bool
}
