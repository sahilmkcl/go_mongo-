package main

import (
	"Go_server/database"
	"Go_server/model"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/registerUser", func(c *gin.Context) {
		var user model.User
		err := c.Bind(&user)
		database.CheckError(err)
		database.CreateUser(user)
	})
	r.GET("/getUser", func(c *gin.Context) {
		res := database.GetUser()
		c.JSON(200, res)
	})
	r.POST("/login", func(c *gin.Context) {
		var user model.User
		err := c.Bind(&user)
		database.CheckError(err)
		var dataUser model.User
		dataUser, err = database.FindUser(user.Name)
		if err != nil {
			c.String(404, "user Not Found")
		} else if dataUser.Password == user.Password {
			c.String(200, "Welcome")
		} else {
			c.String(200, "Password invalid")
		}
	})
	return r
}
