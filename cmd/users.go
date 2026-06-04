package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func(a *application) update(c *gin.Context) {
	userRole := c.GetString("userRole")
	if userRole != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized. Must be a superadmin"})
		return
	}
	context := c.Request.Context()
	var user UpdateRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingUser, err := a.models.Users.GetByEmail(context, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform database query"})
		return
	}
	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Role	= user.Role
	updatedUser, err := a.models.Users.Update(context, existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform database query"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func(a *application) disable(c *gin.Context) {

}