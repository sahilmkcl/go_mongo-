package main

import (
	"Go_server/database"
	"Go_server/model"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func registerRoutes(saveToken map[string]string) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/registerUser", func(c *gin.Context) {
		var user model.User
		err := c.Bind(&user)
		database.CheckError(err)
		database.CreateUser(user)
	})
	r.GET("/getUser", func(c *gin.Context) {
		token := strings.Join(c.Request.Header["Authorization"], "")
		value, ok := saveToken[token]
		if ok == false {
			c.JSON(404, "user not loged in")
		} else {
			// res := database.GetUser()
			c.JSON(200, value)
		}
	})
	r.POST("/update", func(c *gin.Context) {
		var update model.Update
		err := c.Bind(&update)
		database.CheckError(err)
		err = database.UpdateUser(update)
		if err != nil {
			c.String(200, "not updated")
		} else {
			c.String(200, "user updated")
		}

	})
	r.POST("/login", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		var user model.User
		err := c.Bind(&user)
		database.CheckError(err)
		var dataUser model.User
		dataUser, err = database.FindUser(user.Name)
		if err != nil {
			c.String(200, "user Not Found")
		} else if dataUser.Password == user.Password {
			//TODO generate token
			var token string
			token, err = createToken(user.Name)
			saveToken[token] = user.Name
			c.Header("auth", token)
			database.CheckError(err)
		} else {

			c.String(200, "Password invalid")
		}
	})
	return r
}

func createToken(username string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
