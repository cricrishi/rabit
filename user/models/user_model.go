package models

import(
	//"gopkg.in/gin-gonic/gin.v1"
)

type LoginDetail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status string
	Error string
	UniqueKey string
}

