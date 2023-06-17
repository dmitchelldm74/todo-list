package main

import (
	v1 "todo-list/v1"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	v1.BuildGroup(router)

	router.Run()
}

func main() {
	RunServer()
}
