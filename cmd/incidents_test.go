package main

import (
	"bytes"
	"encoding/json"
	"issueTracking/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReportIncident(t *testing.T) {
	db.TruncateTables(t, testPool)

	gin.SetMode(gin.TestMode)

	a := &application{
		origins: "*",
		models:  db.NewModels(testPool),
	}

	r := gin.Default()

	r.POST("/api/v1/incidents", a.reportIncident)

	payload := map[string]any{
		"principalName":       "testName",
		"princidpalGender":    "Male",
		"principalDob":        "today",
		"principalType":       "Patient",
		"patientId":           "iajdaj232",
		"patientWardDept":     "icu",
		"peopleInvolved":      "peopleInvolved",
		"dateOfIncident":      "today",
		"timeOfIncident":      "now",
		"locationOfIncident":  "here",
		"incidentWardDept":    "here?",
		"isNearMiss":          false,
		"causeGroup":          "causeGroup",
		"reporterName":        "Akene Uzezi",
		"reporterDesignation": "???",
		"signature":           true,
		"reporterInfo":        "some info",
		"date":                "today",
		"severityLevel":       "minor",
		"incidentStatus":      "unresolved",
	}
	jsonBody, _ := json.Marshal(&payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/incidents", bytes.NewBuffer(jsonBody))

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "testName", response["principalName"])
}
