package main

import (
	conf "github.com/rabit/config"
	user "github.com/rabit/user/service"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	conf.LoadConfig("config.yaml", "config")
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/users", user.ListUsers)
		v1.POST("/user/login", user.AuthenticateUser)
	}
	router.Run(":3000")

}
