package main

import (
	"api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.DBConnection()
	router := gin.Default()
	router.Static("/public/images/users/", "./public/images/users/")
	router.Run(":8080")
}
