package db

import (
	"context"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DeathReportModel struct {
	DB *pgxpool.Pool
}

type DeathReport struct {
	ID                      int    `json:"id"`
	Ref                     string `json:"ref" binding:"required"`
	ReportedDate            string `json:"reportedDate" binding:"required"`
	IncidentDate            string `json:"incidentDate" binding:"required"`
	IncidentTime            string `json:"incidentTime" binding:"required"`
	Department              string `json:"department" binding:"required"`
	Location                string `json:"location" binding:"required"`
	Category                string `json:"category" binding:"required"`
	SubCategory             string `json:"subCategory" binding:"required"`
	Description             string `json:"description" binding:"required"`
	ActionTaken             string `json:"actionTaken" binding:"required"`
	OpenedDate              string `json:"openedDate"`
	SubmittedTime           string `json:"submittedTime"`
	Handler                 string `json:"handler"`
	Manager                 string `json:"manager"`
	Specialty               string `json:"specialty"`
	ExactLocation           string `json:"exactLocation"`
	Coding                  string `json:"coding"`
	Type                    string `json:"type"`
	RiskGrading             string `json:"riskGrading"`
	Result                  string `json:"result"`
	ActualHarm              string `json:"actualHarm"`
	PotentialHarm           string `json:"potentialHarm"`
	Details                 string `json:"details"`
	PatientInvolved         bool   `json:"patientInvolved"`
	PatientTold             bool   `json:"patientTold"`
	FamilyTold              bool   `json:"familyTold"`
	WhatFamilyTold          string `json:"whatFamilyTold"`
	IncidentInvestigation   string `json:"incidentInvestigation"`
	ReviewMeetingDate       string `json:"reviewMeetingDate"`
	QualityAssuranceLead    string `json:"qualityAssuranceLead"`
	DoctorNotified          bool   `json:"doctorNotified"`
	MeetingDiscussionPoints string `json:"meetingDiscussionPoints"`
	MeetingActionPoints     string `json:"meetingActionPoints"`
	LevelOfInvestigation    string `json:"levelOfInvestigation"`
}

func (m *DeathReportModel) InsertDeathReport(ctx context.Context, deathReport *DeathReport) error {
	query := `
    INSERT INTO death_reports (
        ref, reported_date, incident_date, incident_time, department, location,
        category, sub_category, description, action_taken, opened_date, submitted_time,
        handler, manager, specialty, exact_location, coding, type, risk_grading,
        result, actual_harm, potential_harm, details, patient_involved, patient_told,
        family_told, what_family_told, incident_investigation, review_meeting_date,
        quality_assurance_lead, doctor_notified, meeting_discussion_points,
        meeting_action_points, level_of_investigation
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
        $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
        $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
        $31, $32, $33, $34
    ) RETURNING id
`
	// Use QueryRow instead of Exec to capture RETURNING id
	err := m.DB.QueryRow(
		ctx, query,
		deathReport.Ref, deathReport.ReportedDate, deathReport.IncidentDate, deathReport.IncidentTime,
		deathReport.Department, deathReport.Location, deathReport.Category, deathReport.SubCategory,
		deathReport.Description, deathReport.ActionTaken, deathReport.OpenedDate, deathReport.SubmittedTime,
		deathReport.Handler, deathReport.Manager, deathReport.Specialty, deathReport.ExactLocation,
		deathReport.Coding, deathReport.Type, deathReport.RiskGrading, deathReport.Result,
		deathReport.ActualHarm, deathReport.PotentialHarm, deathReport.Details, deathReport.PatientInvolved,
		deathReport.PatientTold, deathReport.FamilyTold, deathReport.WhatFamilyTold,
		deathReport.IncidentInvestigation, deathReport.ReviewMeetingDate, deathReport.QualityAssuranceLead,
		deathReport.DoctorNotified, deathReport.MeetingDiscussionPoints, deathReport.MeetingActionPoints,
		deathReport.LevelOfInvestigation,
	).Scan(&deathReport.ID)
	if err != nil {
		return fmt.Errorf("database query error: %w", err)
	}

	return nil
}

func (m *DeathReportModel) SearchByID(ctx context.Context, id int) (*DeathReport, error) {
	var deathReport DeathReport
	query := `
	SELECT 
		id, ref, reported_date, incident_date, incident_time, department, location,
		category, sub_category, description, action_taken, opened_date, submitted_time,
		handler, manager, specialty, exact_location, coding, type, risk_grading,
		result, actual_harm, potential_harm, details, patient_involved, patient_told,
		family_told, what_family_told, incident_investigation, review_meeting_date,
		quality_assurance_lead, doctor_notified, meeting_discussion_points,
		meeting_action_points, level_of_investigation
	FROM death_reports
	WHERE id = $1
	`
	err := m.DB.QueryRow(ctx, query, id).Scan(
		&deathReport.ID,
		&deathReport.Ref,
		&deathReport.ReportedDate,
		&deathReport.IncidentDate,
		&deathReport.IncidentTime,
		&deathReport.Department,
		&deathReport.Location,
		&deathReport.Category,
		&deathReport.SubCategory,
		&deathReport.Description,
		&deathReport.ActionTaken,
		&deathReport.OpenedDate,
		&deathReport.SubmittedTime,
		&deathReport.Handler,
		&deathReport.Manager,
		&deathReport.Specialty,
		&deathReport.ExactLocation,
		&deathReport.Coding,
		&deathReport.Type,
		&deathReport.RiskGrading,
		&deathReport.Result,
		&deathReport.ActualHarm,
		&deathReport.PotentialHarm,
		&deathReport.Details,
		&deathReport.PatientInvolved,
		&deathReport.PatientTold,
		&deathReport.FamilyTold,
		&deathReport.WhatFamilyTold,
		&deathReport.IncidentInvestigation,
		&deathReport.ReviewMeetingDate,
		&deathReport.QualityAssuranceLead,
		&deathReport.DoctorNotified,
		&deathReport.MeetingDiscussionPoints,
		&deathReport.MeetingActionPoints,
		&deathReport.LevelOfInvestigation,
	)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}

	return &deathReport, nil
}

func (m *DeathReportModel) UpdateDeathReport(ctx context.Context, reportUpdate *DeathReport) error {
	query := `
		UPDATE death_reports SET
			ref = $1,
			reported_date = $2,
			opened_date = $3,
			submitted_time = $4,
			handler = $5,
			manager = $6,
			location = $7,
			department = $8,
			specialty = $9,
			exact_location = $10,
			coding = $11,
			type = $12,
			category = $13,
			sub_category = $14,
			risk_grading = $15,
			result = $16,
			actual_harm = $17,
			potential_harm = $18,
			details = $19,
			incident_date = $20,
			incident_time = $21,
			description = $22,
			action_taken = $23,
			patient_involved = $24,
			patient_told = $25,
			family_told = $26,
			what_family_told = $27,
			incident_investigation = $28,
			review_meeting_date = $29,
			quality_assurance_lead = $30,
			doctor_notified = $31,
			meeting_discussion_points = $32,
			meeting_action_points = $33,
			level_of_investigation = $34
		WHERE id = $35
	`

	_, err := m.DB.Exec(
		ctx, query,
		reportUpdate.Ref,                     // $1
		reportUpdate.ReportedDate,            // $2
		reportUpdate.OpenedDate,              // $3
		reportUpdate.SubmittedTime,           // $4
		reportUpdate.Handler,                 // $5
		reportUpdate.Manager,                 // $6
		reportUpdate.Location,                // $7
		reportUpdate.Department,              // $8
		reportUpdate.Specialty,               // $9
		reportUpdate.ExactLocation,           // $10
		reportUpdate.Coding,                  // $11
		reportUpdate.Type,                    // $12
		reportUpdate.Category,                // $13
		reportUpdate.SubCategory,             // $14
		reportUpdate.RiskGrading,             // $15
		reportUpdate.Result,                  // $16
		reportUpdate.ActualHarm,              // $17
		reportUpdate.PotentialHarm,           // $18
		reportUpdate.Details,                 // $19
		reportUpdate.IncidentDate,            // $20
		reportUpdate.IncidentTime,            // $21
		reportUpdate.Description,             // $22
		reportUpdate.ActionTaken,             // $23
		reportUpdate.PatientInvolved,         // $24
		reportUpdate.PatientTold,             // $25
		reportUpdate.FamilyTold,              // $26
		reportUpdate.WhatFamilyTold,          // $27
		reportUpdate.IncidentInvestigation,   // $28
		reportUpdate.ReviewMeetingDate,       // $29
		reportUpdate.QualityAssuranceLead,    // $30
		reportUpdate.DoctorNotified,          // $31
		reportUpdate.MeetingDiscussionPoints, // $32
		reportUpdate.MeetingActionPoints,     // $33
		reportUpdate.LevelOfInvestigation,    // $34  <-- THIS WAS MISSING
		reportUpdate.ID,                      // $35
	)
	if err != nil {
		return fmt.Errorf("database query error: %w", err)
	}

	return nil
}

func (m *DeathReportModel) GetDeathReports(ctx context.Context, limit, offset int) (*[]DeathReport, int, int, error) {
	var totalItems int
	err := m.DB.QueryRow(ctx, "SELECT COUNT(*) FROM death_reports").Scan(&totalItems)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("database query error: %w", err)
	}
	query := `
		SELECT * FROM death_reports
		ORDER BY id DESC	
		LIMIT $1 OFFSET $2
	`
	rows, err := m.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("database query err: %w", err)
	}
	defer rows.Close()
	deathReport, err := pgx.CollectRows(rows, pgx.RowToStructByName[DeathReport])
	if err != nil {
		return nil, 0, 0, fmt.Errorf("database query error: %w", err)
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	if totalPages == 0 {
		totalPages = 1
	}
	return &deathReport, totalPages, totalItems, nil
}

func (m *DeathReportModel) SearchDeathReports(ctx context.Context, searchQuery string) (*[]DeathReport, error) {
	safeSearchQuery := fmt.Sprintf("%%%s%%", searchQuery)
	query := `
		SELECT * FROM death_reports
		WHERE (
			COALESCE(ref, '') || ' ' ||
  		COALESCE(department, '') || ' ' ||
  		COALESCE(location, '') || ' ' ||
  		COALESCE(exact_location, '') || ' ' ||
  		COALESCE(handler, '') || ' ' ||
  		COALESCE(manager, '') || ' ' ||
  		COALESCE(specialty, '') || ' ' ||
  		COALESCE(coding, '') || ' ' ||
  		COALESCE(category, '') || ' ' ||
  		COALESCE(sub_category, '') || ' ' ||
  		COALESCE(description, '') || ' ' ||
  		COALESCE(details, '') || ' ' ||
  		COALESCE(action_taken, '') || ' ' ||
  		COALESCE(quality_assurance_lead, '') || ' ' ||
  		COALESCE(incident_investigation, '')
	) ILIKE $1;
	`
	rows, err := m.DB.Query(ctx, query, safeSearchQuery)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	deathReports, err := pgx.CollectRows(rows, pgx.RowToStructByName[DeathReport])
	if err != nil {
		return nil, fmt.Errorf("failed to scan row to struct: %w", err)
	}

	return &deathReports, nil
}
