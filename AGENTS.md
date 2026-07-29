# Agent Instructions for Issue Tracker

## Project Overview

The Issue Tracker is a RESTful API for managing workplace incidents and safety reports built with Go, Gin, and PostgreSQL.

**Code Metrics:**
- Total Go code: 3646 lines
- 28 Go source files
- Architecture: Clean layered (presentation → application → data → infrastructure)

## Development Commands

```bash
# Start development server with live reload
air

# Run application directly
go run ./cmd/

# Run tests
go test ./...
# Or for verbose output:
./scripts/runtests.sh

# Format code
go fmt ./...

# Run linter
go vet ./...
```

## Docker Commands

```bash
# Start all services (API at localhost:3002)
docker compose up -d

# Stop services
docker compose down

# Remove volumes (fresh database)
docker compose down -v

# View logs
docker compose logs -f
```

## Database Access

```bash
# Access PostgreSQL shell
./scripts/login.sh
```

## API Testing

```bash
# Health check
curl http://localhost:3002/api/v1/ping

# Login (save token)
TOKEN=$(curl -s -X POST http://localhost:3002/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"yourpassword"}' | jq -r '.token')

# Register a new user (requires superadmin token)
curl -X POST http://localhost:3002/api/v1/auth/register \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"email":"newuser@example.com","name":"New User","password":"password123","role":"admin","department":"IT"}'

# Report incident (no auth required)
curl -X POST http://localhost:3002/api/v1/incidents \
  -H "Content-Type: application/json" \
  -d '{
    "principalName": "John Doe",
    "principalGender": "Male",
    "principalDob": "1990-01-15",
    "principalType": "patient",
    "patientId": "P12345",
    "patientWardDept": "Ward A",
    "peopleInvolved": "Nurse Smith",
    "dateOfIncident": "2026-06-09",
    "timeOfIncident": "14:00",
    "locationOfIncident": "Ward A, Room 3",
    "incidentWardDept": "Ward A",
    "witnesses": "Dr. Brown",
    "witnessType": "Staff",
    "witnessWardDept": "Ward A",
    "witnessJobTitle": "Doctor",
    "witenssPhone": "555-0100",
    "isNearMiss": false,
    "causeGroup": "Fall",
    "causes": "Wet floor",
    "prescribingDoctor": "Dr. Brown",
    "treatmentReceived": "First Aid",
    "equipmentInvolved": "No",
    "equipmentSentForRepair": false,
    "equipmentWithdrawn": false,
    "equipmentRetained": false,
    "isMedicalDevice": "No",
    "reporterName": "Jane Reporter",
    "reporterDesignation": "Nurse",
    "signature": true,
    "reporterInfo": "jane@example.com",
    "date": "2026-06-09",
    "severityLevel": "minor"
  }'

# Add comment to incident (requires manager or admin)
curl -X POST http://localhost:3002/api/v1/incidents/comments \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"incidentId": 1, "userId": 2, "comment": "Follow up needed"}'

# Get incidents (requires auth)
curl http://localhost:3002/api/v1/incidents -H "Authorization: Bearer $TOKEN"

# Update incident status (requires auth; reporter/supervisor/manager roles forbidden)
curl -X PATCH http://localhost:3002/api/v1/incidents/1/status \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"resolved"}'

# Get user info (requires superadmin role)
curl "http://localhost:3002/api/v1/user?email=test@example.com" -H "Authorization: Bearer $TOKEN"

# Get comments for incident (requires admin or manager)
curl "http://localhost:3002/api/v1/incidents/comments?incidentId=1" -H "Authorization: Bearer $TOKEN"

# Get incident management logs (requires admin or manager role)
curl "http://localhost:3002/api/v1/incidents/1/managementlogs" -H "Authorization: Bearer $TOKEN"

# Report a death (no auth required)
curl -X POST http://localhost:3002/api/v1/deathreport \
  -H "Content-Type: application/json" \
  -d '{
    "ref": "DR-001",
    "reportedDate": "2026-06-09",
    "incidentDate": "2026-06-09",
    "incidentTime": "14:00",
    "department": "IT",
    "location": "Ward A",
    "category": "Category",
    "subCategory": "SubCategory",
    "description": "Description",
    "actionTaken": "Action taken",
    "openedDate": "2026-06-09",
    "submittedTime": "14:00",
    "handler": "Handler",
    "manager": "Manager",
    "specialty": "Specialty",
    "exactLocation": "Exact Location",
    "coding": "Coding",
    "type": "Type",
    "riskGrading": "High",
    "result": "Result",
    "actualHarm": "Actual Harm",
    "potentialHarm": "Potential Harm",
    "details": "Details",
    "patientInvolved": true,
    "patientTold": true,
    "familyTold": true,
    "whatFamilyTold": "What family told",
    "incidentInvestigation": "Investigation",
    "reviewMeetingDate": "2026-06-09",
    "qualityAssuranceLead": "QA Lead",
    "doctorNotified": true,
    "meetingDiscussionPoints": "Discussion points",
    "meetingActionPoints": "Action points",
    "levelOfInvestigation": "Level"
  }'

# Update a death report (no auth required)
curl -X PUT http://localhost:3002/api/v1/deathreport \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "ref": "DR-001",
    "reportedDate": "2026-06-09",
    "incidentDate": "2026-06-09",
    "incidentTime": "14:00",
    "department": "IT",
    "location": "Ward A",
    "category": "Category",
    "subCategory": "SubCategory",
    "description": "Updated description",
    "actionTaken": "Updated action",
    "openedDate": "2026-06-09",
    "submittedTime": "14:00",
    "handler": "Handler",
    "manager": "Manager",
    "specialty": "Specialty",
    "exactLocation": "Exact Location",
    "coding": "Coding",
    "type": "Type",
    "riskGrading": "High",
    "result": "Result",
    "actualHarm": "Actual Harm",
    "potentialHarm": "Potential Harm",
    "details": "Details",
    "patientInvolved": true,
    "patientTold": true,
    "familyTold": true,
    "whatFamilyTold": "What family told",
    "incidentInvestigation": "Investigation",
    "reviewMeetingDate": "2026-06-09",
    "qualityAssuranceLead": "QA Lead",
    "doctorNotified": true,
    "meetingDiscussionPoints": "Discussion points",
    "meetingActionPoints": "Action points",
    "levelOfInvestigation": "Level"
  }'

# Get all death reports (no auth required)
curl "http://localhost:3002/api/v1/deathreports?page=1&limit=10"

# Search death reports (no auth required)
curl "http://localhost:3002/api/v1/searchDeathReport?searchQuery=DR-001"
```

## Role Permissions

| Role | Permissions |
|------|-------------|
| superadmin | All endpoints including user management (register, update, disable, enable, reset password, get user), report incidents, view all incidents, update any incident status, submit incident management reports, update incident management reports, add comments, view comments |
| admin | Report incidents, view all incidents, update any incident status, submit incident management reports, update incident management reports, add comments, view comments |
| supervisor | Report incidents, view own department incidents (matched via `incident_ward_dept`, `patient_ward_dept`, or `staff_place_of_work`) |
| manager | Report incidents, view all incidents, view incident management reports and logs, add comments, view comments, submit incident management reports, update incident management reports |
| reporter | Report incidents via public endpoint only, view own department incidents |

## Death Report

The death report feature captures detailed information about workplace deaths. These endpoints are public and do not require authentication.

### API Endpoints

**POST /api/v1/deathreport** - Create death report
- Requires: No authentication
- Request body: `DeathReport` struct

**PUT /api/v1/deathreport** - Update death report
- Requires: No authentication
- Request body: `DeathReport` struct (must include `id`)

**GET /api/v1/deathreports** - Get all death reports (paginated)
- Requires: No authentication
- Query Parameters: `page` (default 1), `limit` (default 10, max 50)

**GET /api/v1/searchDeathReport** - Search death reports
- Requires: No authentication
- Query Parameters: `searchQuery` (if empty, returns all reports)

### DeathReport Struct

```go
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
```

## Incident Management Form

The incident management report captures follow-up documentation after an incident occurs. The form includes:

### Form Sections

1. **Operational Evaluation Metrics**
   - Impact on Service
   - Contributory Factors
   - Actions Taken / Outcomes
   - Recommendations
   - Lessons Learned

2. **Stakeholder Notifications Log**
   - Patient Informed
   - Relative Informed
   - Senior Manager Notified
   - Pharmacist Informed
   - Police Incident Number
   - Other Informed Parties

3. **Risk Factor Assessment**
   - Risk Severity Score (1-5)
   - Risk Likelihood Score (1-5)
   - Risk Rating (submitted value; typically Severity × Likelihood)

4. **Occupational Health & Safety Compliance**
   - Staff Absence Over 3 Days
   - Act of Violence or Peril Danger
   - Hospitalization Over 24 Hours
   - OHS Impacted Staff Name
   - Staff Date of Birth
   - Staff Home Address

5. **Executive Authorization Sign-Off**
   - Manager Name
   - Corporate Designation
   - Authorization Date
   - Manager Signature (required, legally binding)

### API Endpoints

**POST /api/v1/incidents/:id/management** - Create management report
- Requires: admin or manager role
- Request body: `IncidentManagement` struct

**GET /api/v1/incidents/:id/management** - Retrieve management report
- Requires: Authentication (any authenticated user)
- Response: `IncidentManagement` struct

**PUT /api/v1/incidents/:id/management** - Update existing report
- Requires: manager or admin role
- Request body: `IncidentManagement` struct

### IncidentManagement Struct

```go
type IncidentManagement struct {
    Id                              int    `json:"id"`
    IncidentId                      int    `json:"incidentId"`
    ImpactOnService                 string `json:"impactOnService" binding:"required"`
    ContributoryFactors             string `json:"contributoryFactors" binding:"required"`
    ActionsTakenOutcomes            string `json:"actionsTakenOutcomes" binding:"required"`
    Recommendations                 string `json:"recommendations" binding:"required"`
    LessonsLearned                  string `json:"lessonsLearned" binding:"required"`
    InformedPatient                 bool   `json:"informedPatient"`
    InformedRelative                bool   `json:"informedRelative"`
    InformedSeniorManager           bool   `json:"informedSeniorManager"`
    InformedPharmacist              bool   `json:"informedPharmacist"`
    PoliceIncidentNumber            string `json:"policeIncidentNumber,omitempty"`
    InformedOther                   string `json:"informedOther,omitempty"`
    RiskSeverity                    int    `json:"riskSeverity" binding:"required"`
    RiskLikelihood                  int    `json:"riskLikelihood" binding:"required"`
    RiskRating                      int    `json:"riskRating" binding:"required"`
    OhsAbsenceOver3Days             bool   `json:"ohsAbsenceOver3Days"`
    OhsActOfViolenceOrDanger        bool   `json:"ohsActOfViolenceOrDanger"`
    OhsHospitalizationOver24Hours   bool   `json:"ohsHospitalizationOver24Hours"`
    OhsStaffName                    string `json:"ohsStaffName"`
    OhsStaffDob                     string `json:"ohsStaffDob"`
    OhsStaffAddress                 string `json:"ohsStaffAddress"`
    ManagerName                     string `json:"managerName" binding:"required"`
    ManagerSignature                bool   `json:"managerSignature" binding:"required"`
    ManagerDesignation              string `json:"managerDesignation" binding:"required"`
    ManagerDate                     string `json:"managerDate" binding:"required"` // date this was filled
}
```

## Default Credentials

A superadmin user is created by default:
- Email: `admin@example.com`
- Password: The default password is hashed with bcrypt and stored in `tables.sql`. Check the database or reset it via code to set a known password.

**Note:** New users registered via `/api/v1/auth/register` are assigned a default password of `redeemershealthvillage` if none is provided. This is separate from the pre-seeded superadmin.
