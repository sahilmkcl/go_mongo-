package main

import (
	"Go_server/database"
	"Go_server/model"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/addUser", func(c *gin.Context) {
		var user model.User
		err := c.Bind(&user)
		database.CheckError(err)
		database.CreateUser(user)
	})
	r.GET("/getUser", func(c *gin.Context) {
		res := database.GetUser()
		c.JSON(200, res)
	})
	return r
}
