package service

import (
	"fmt"
	db "github.com/rabit/database"
	"gopkg.in/gin-gonic/gin.v1"
)

type loginDetail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ListUsers(c *gin.Context) {
	err,users := db.GetUserDetail()
	c.JSON(200, gin.H{
		"status":  "Listed",
		"message": "This Api list the users from user table",
		"nick":    "test",
		"error": err,
		"users":    users,
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
