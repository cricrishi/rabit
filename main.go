package main

import (
	user "github.com/rabit/user/service"
	conf "github.com/rabit/config"
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
)



func main() {
	conf.LoadConfig("config.yaml","config")
	fmt.Println(conf.ConfigData["config"].Get("Test_Url"))
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/users", user.ListUsers)
		v1.POST("/user/login", user.AuthenticateUser)
	}	
	router.Run(":3000")

}
