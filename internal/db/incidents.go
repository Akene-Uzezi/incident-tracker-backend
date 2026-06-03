package db

import "github.com/jackc/pgx/v5/pgxpool"

type IncidentsModel struct {
	DB *pgxpool.Pool
}

type SeverityLevel string

const (
	NearMiss SeverityLevel = "Near Miss"
	Minor SeverityLevel = "Minor"
	Major SeverityLevel = "Major"
	Critical SeverityLevel = "Critical"
)

func (s SeverityLevel) IsValid() bool {
	switch s{
		case NearMiss, Minor, Major, Critical:
			return true
	}
	return  false
}

type Incident struct {
	Id int `json:"id"`
	ReporterName string `json:"reporterName"`
	Department string `json:"department"`
	Position string `json:"position"`
	ContactInfo string `json:"contactInfo"`
	DateOfIncident string `json:"dateOfIncident"`
	TimeOfIncident string `json:"timeOfIncident"`
	LocationOfIncident string `json:"locationOfIncident"`
	TypeOfIncident string `json:"typeOfIncident"`
	PeopleInvolved string `json:"peopleInvolved"`
	DescriptionOfIncident string `json:"descriptionOfIncident"`
	ImmediateActionTaken string `json:"immediateActionTaken"`
	InjuryOrDamage string `json:"injuryOrDamage"` 
	SeverityLevel SeverityLevel `json:"severityLevel"`
	SupervisorNotified string `json:"supervisorNotified"`
	RecommendedPreventiveAction string `json:"recommendedPreventiveAction"` 
}