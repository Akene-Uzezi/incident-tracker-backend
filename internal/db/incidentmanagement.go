package db

import "github.com/jackc/pgx/v5/pgxpool"

type IncidentManagementModel struct {
	DB *pgxpool.Pool
}

type IncidentManagement struct {
	IncidentId int `json:"incidentId" binding:"required"`

	ImpactOnService string `json:"impactOnService" binding:"required"`
	ContributoryFactors string `json:"contributoryFactors" binding:"required"`
	ActionsTakenOutcomes string `json:"actionsTakenOutcomes" binding:"required"`
	Recommendations string `json:"recommendations" binding:"required"`
	LessonsLearned string `json:"lessonsLearned" binding:"required"`

	InformedPatient bool `json:"informedPatient,omitempty"`
	InformedRelative bool `json:"informedRelative,omitempty"`
	InformedSeniorManager bool `json:"informedSeniorManager,omitempty"`
	InformedPharamacist bool `json:"informedPharmacist,omitempty"`
	PoliceIncidentNumber string `json:"policeIncidentNumber,omitempty"`
	InformedOther string `json:"informedOther,omitempty"`

	RiskSeverity int `json:"riskSeverity" binding:"required"`
	RiskLikelyhood int `json:"riskLikelyhood" binding:"required"`
	RiskRating int `json:"riskRating" binding:"required"`

	OhsAbscenceOver3Days bool `json:"ohsAbscenceOver3Days"`
	OhsActOfViolenceOrDanger bool `json:"ohsActOfViolenceOrDanger"`
	OhsHospitalisationOver24Hours bool `json:"ohsHospitalisationOver24Hours"`
	OhsStaffName string `json:"ohsStaffName"`
	OhsStaffDob string `json:"ohsStaffDob"`
	OhsStaffAddress string `json:"ohsStaffAddress"`

	ManagerName string `json:"managerName" binding:"required"`
	ManagerSignature bool `json:"managerSignature" binding:"required"`
	ManagerDesignation string `json:"managerDesignation" binding:"required"`
	ManagerDate string `json:"managerDate" binding:"required"` // date this was filled
}