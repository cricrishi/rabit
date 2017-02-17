package service

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	conf "github.com/rabit/config"
	
)

type loginDetail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ListUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "my message",
		"nick":    "test",
		"url":     conf.ConfigData["config"].Get("Test_Url"),
	})

}

func AuthenticateUser(c *gin.Context) {
	var requestData loginDetail
	c.BindJSON(&requestData)
	fmt.Println(requestData)
	c.JSON(200, gin.H{
		"email":    requestData.Email,
		"password": requestData.Password,
	})

}
