package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email string `json:"email" binding:"required"`
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required min=8"`
}

func(a *application) register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
}