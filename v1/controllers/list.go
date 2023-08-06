package controllers

import (
	"net/http"
	"strconv"

	db "todo-list/database"
	. "todo-list/v1/models"

	"github.com/gin-gonic/gin"
)

func BuildListGroup(router *gin.RouterGroup) *gin.RouterGroup {
	list := router.Group("/list")
	{
		list.POST("/create", createTodoList)
		list.DELETE("/delete", removeTodoList)
		list.GET("/all", listAllTodoLists)
	}
	return list
}

type CreateTodoListInput struct {
	TodoListName string `json:"listName" binding:"required"`
	AccountId    uint   `json:"accountId" binding:"required"`
}

// POST /v1/list/create
// Create a todo list
func createTodoList(c *gin.Context) {
	// Validate input
	var input CreateTodoListInput
	accountId, hasAccountId := c.GetPostForm("accountId")
	listName, hasListName := c.GetPostForm("listName")

	if hasAccountId && hasListName {
		parsedAccountId, _ := strconv.Atoi(accountId)

		input.AccountId = uint(parsedAccountId)
		input.TodoListName = listName
	} else if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := TodoList{
		TodoListName: input.TodoListName,
		AccountId:    input.AccountId,
	}
	db.ConnectDB().Create(&list)

	c.JSON(http.StatusOK, list)
}

// DELETE /v1/list/remove
// Remove a todo list
func removeTodoList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"removed": true})
}

// GET /v1/list/all
// List all todo lists for an account
func listAllTodoLists(c *gin.Context) {
	var lists []TodoList
	var accountId = c.Query("accountId")

	if accountId == "" {
		db.ConnectDB().Find(&lists)
	} else {
		db.ConnectDB().Where("account_id = ?", accountId).Find(&lists)
	}

	c.JSON(http.StatusOK, lists)
}
