package main

import (
	"issueTracking/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(a *application) reportIncident(c *gin.Context) {
	context := c.Request.Context()
	var input IncidentReport
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !input.SeverityLevel.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid severity level provided"})
		return
	}

	dbIncident := &db.Incident{
		ReporterName:                input.ReporterName,
		Department:                  input.Department,
		Position:                    input.Position,
		ContactInfo:                 input.ContactInfo,
		DateOfIncident:              input.DateOfIncident,
		TimeOfIncident:              input.TimeOfIncident,
		LocationOfIncident:          input.LocationOfIncident,
		TypeOfIncident:              input.TypeOfIncident,
		PeopleInvolved:              input.PeopleInvolved,
		DescriptionOfIncident:       input.DescriptionOfIncident,
		ImmediateActionTaken:        input.ImmediateActionTaken,
		InjuryOrDamage:              input.InjuryOrDamage,
		SeverityLevel:               db.SeverityLevel(input.SeverityLevel),
		SupervisorNotified:          input.SupervisorNotified,
		RecommendedPreventiveAction: input.RecommendedPreventiveAction,
	}

	savedIncident, err := a.models.Incidents.Insert(context, dbIncident)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform database query"})
		return
	}
	c.JSON(http.StatusOK, savedIncident)
}