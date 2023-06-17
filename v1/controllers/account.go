package controllers

import (
	"net/http"

	db "todo-list/database"
	. "todo-list/v1/models"

	"github.com/gin-gonic/gin"
)

func BuildAccountGroup(router *gin.RouterGroup) *gin.RouterGroup {
	account := router.Group("/account")
	{
		account.GET("/get-id", getAccountId)
		account.GET("/list", listAccounts)
	}
	return account
}

// GET /v1/account/get-id
// Get a unique user id
func getAccountId(c *gin.Context) {
	account := Account{}
	db.ConnectDB().Create(&account)

	c.JSON(http.StatusOK, account)
}

// GET /v1/account/list
// List all accounts
func listAccounts(c *gin.Context) {
	var accounts []Account
	db.ConnectDB().Find(&accounts)

	c.JSON(http.StatusOK, accounts)
}
