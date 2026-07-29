package main

import (
	"fmt"
	"net/http"

	"issueTracking/internal/db"

	"github.com/gin-gonic/gin"
)

func (a *application) deathReport(c *gin.Context) {
	var deathReport db.DeathReport
	if err := c.ShouldBindJSON(&deathReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A bad request was sent"})
		return
	}
	context := c.Request.Context()
	err := a.models.DeathReport.InsertDeathReport(context, &deathReport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "The death has been reported"})
}

func (a *application) updateDeathReport(c *gin.Context) {
	var updateRequest db.DeathReport
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("A bad request was passed: %v", err.Error())})
		return
	}
	ctx := c.Request.Context()
	existingReport, err := a.models.DeathReport.SearchByID(ctx, updateRequest.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if existingReport != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	existingReport = &updateRequest
	c.JSON(http.StatusOK, gin.H{"deathReport": existingReport})
}

func (a *application) searchDeathReport(c *gin.Context) {
}
