package main

import (
	"net/http"
	v1 "todo-list/v1"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	v1.BuildGroup(router)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TODO List",
		})
	})

	router.Run()
}

func main() {
	RunServer()
}
