package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	userHelper "github.com/rabit/user/helper"
	userModel "github.com/rabit/user/models"
)



func ListUsers(c *gin.Context) {
	err,users := userHelper.ListUsers()
	c.JSON(200, gin.H{
		"status":  "Listed",
		"message": "This Api list the users from user table",
		"nick":    "test",
		"error": err,
		"users":    users,
	})

}

func AuthenticateUser(c *gin.Context) {
	var requestData userModel.LoginDetail
	c.BindJSON(&requestData)
	r := userHelper.AuthenticateUser(requestData)
	c.JSON(200, gin.H{
		"status":    r.Status,
		"error": r.Error,
		"uuid": r.UniqueKey,
	})

}
