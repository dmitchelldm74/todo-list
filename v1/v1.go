package v1

import (
	db "todo-list/database"
	"todo-list/v1/controllers"
	. "todo-list/v1/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildGroup(router *gin.Engine) *gin.RouterGroup {
	db := db.ConnectDB()

	MigrateModels(db)
	v1 := router.Group("/v1")
	{
		controllers.BuildAccountGroup(v1)
		controllers.BuildListGroup(v1)
	}

	return v1
}

func MigrateModels(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&TodoList{})
	db.AutoMigrate(&TodoListItem{})
}
