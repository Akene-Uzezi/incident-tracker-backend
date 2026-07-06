package main

import (
	"net/http"

	"issueTracking/internal/db"

	"github.com/gin-gonic/gin"
)

func (a *application) addComment(c *gin.Context) {
	userRole := c.GetString("userRole")
	if userRole != "manager" && userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed add comments to this incident"})
		return
	}
	var comment *db.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context := c.Request.Context()
	if err := a.models.Comments.InsertComment(context, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "comment added"})
}

func (a *application) getComments(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.GetString("userRole")
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access incident comments"})
		return
	}

	var commentRequest *CommentRequest
	if err := c.ShouldBindJSON(commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments, err := a.models.Comments.GetComments(ctx, commentRequest.IncidentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
