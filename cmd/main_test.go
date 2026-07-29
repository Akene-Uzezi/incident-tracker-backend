package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"issueTracking/internal/db"

	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testPool *pgxpool.Pool

func mockAuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userRole", role)
		c.Next()
	}
}

func mockEmailMiddleware(email string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userEmail", email)
		c.Next()
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	pool, cleanup, err := db.SetupTestDBSuite()
	if err != nil {
		log.Fatalf("failed to initialize test container: %v", err)
	}
	testPool = pool

	exitCode := m.Run()

	cleanup()

	os.Exit(exitCode)
}

func insertIncident(payload *db.Incident, a *application, t *testing.T) error {
	if _, err := a.models.Incidents.Insert(context.Background(), payload); err != nil {
		return fmt.Errorf("error seeding test data into incidents")
	}

	return nil
}

func insertUser(a *application, t *testing.T) error {
	if _, err := a.models.Users.Insert(context.Background(), "testuser", "testuser@example.com", "testpassword", "admin", "it"); err != nil {
		return fmt.Errorf("error seeding testdata into users: %v", err)
	}
	return nil
}

func insertDeathReport(payload *db.DeathReport, a *application, t *testing.T) error {
	if err := a.models.DeathReport.InsertDeathReport(context.Background(), payload); err != nil {
		return fmt.Errorf("error seeding testdata into deathreport: %v", err)
	}
	return nil
}

func insertDeathReportWithPayload(a *application, t *testing.T) error {
	payload := &db.DeathReport{
		Ref:                     "DR-2026-001",
		ReportedDate:            "2026-07-29",
		IncidentDate:            "2026-07-28",
		IncidentTime:            "14:30",
		Department:              "Cardiology",
		Location:                "Building A",
		Category:                "Clinical Incident",
		SubCategory:             "Patient Care",
		Description:             "Test death report description for automated testing.",
		ActionTaken:             "Immediate review initiated.",
		OpenedDate:              "2026-07-29",
		SubmittedTime:           "09:00",
		Handler:                 "Dr. John Doe",
		Manager:                 "Jane Smith",
		Specialty:               "Internal Medicine",
		ExactLocation:           "Ward 3, Bed 12",
		Coding:                  "COD-101",
		Type:                    "Clinical Incident",
		RiskGrading:             "High",
		Result:                  "Under Review",
		ActualHarm:              "Severe",
		PotentialHarm:           "Critical",
		Details:                 "Additional test details regarding the event.",
		PatientInvolved:         true,
		PatientTold:             true,
		FamilyTold:              true,
		WhatFamilyTold:          "Family was informed by the attending physician.",
		IncidentInvestigation:   "Investigation ongoing by QA lead.",
		ReviewMeetingDate:       "2026-08-01",
		QualityAssuranceLead:    "Dr. Alice Johnson",
		DoctorNotified:          true,
		MeetingDiscussionPoints: "Discussed protocol adherence and timeline of events.",
		MeetingActionPoints:     "Update ward monitoring checklists.",
		LevelOfInvestigation:    "Comprehensive",
	}
	if err := a.models.DeathReport.InsertDeathReport(context.Background(), payload); err != nil {
		return fmt.Errorf("error inserting seed data: %w", err)
	}

	return nil
}
