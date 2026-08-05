# Agent Instructions for Issue Tracker

## Project Overview

The Issue Tracker is a RESTful API for managing workplace incidents and safety reports built with Go, Gin, and PostgreSQL.

**Code Metrics:**
- Total Go code: 4103 lines
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
# Start all services (API at localhost:3002, internally port 3001)
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
docker exec -it issuetracker_db psql -U tracker_user -d issuetracker

# Or via local script
./scripts/login.sh
```

## Scripts

| Script | Purpose |
|--------|---------|
| `./scripts/runtests.sh` | Run tests verbosely |
| `./scripts/login.sh` | Open PostgreSQL shell in Docker container |
| `./scripts/createtable.sh` | Create database tables |
| `./scripts/resetdb.sh` | Reset database |
| `./scripts/restart.sh` | Restart services |
| `./scripts/commit.sh` | Commit helper |
| `./scripts/format.sh` | Format code |

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
    "staffJobTitle": "Nurse",
    "staffPhone": "555-0100",
    "staffPlaceOfWork": "Ward A",
    "staffSite": "Main Hospital",
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
    "equipmentModel": "Model X",
    "equipmentSentForRepair": false,
    "equipmentWithdrawn": false,
    "equipmentRetained": false,
    "equipmentNumber": "EQ-123",
    "isMedicalDevice": "No",
    "reporterName": "Jane Reporter",
    "reporterDesignation": "Nurse",
    "signature": true,
    "reporterInfo": "jane@example.com",
    "date": "2026-06-09",
    "severityLevel": "minor",
    "incidentStatus": "unresolved"
  }'

# Get incidents (paginated, requires auth)
curl "http://localhost:3002/api/v1/incidents?page=1&limit=10" -H "Authorization: Bearer $TOKEN"

# Search incidents (requires auth)
curl "http://localhost:3002/api/v1/searchIncidents?searchQuery=Ward%20A" -H "Authorization: Bearer $TOKEN"

# Update incident status (requires admin or superadmin; body field is `incidentStatus`)
curl -X PATCH http://localhost:3002/api/v1/incidents/1/status \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"incidentStatus":"resolved"}'

# Get all users (paginated, requires superadmin)
curl "http://localhost:3002/api/v1/users?page=1&limit=10" -H "Authorization: Bearer $TOKEN"

# Get user by email (requires superadmin)
curl "http://localhost:3002/api/v1/user?email=test@example.com" -H "Authorization: Bearer $TOKEN"

# Search users (requires superadmin)
curl "http://localhost:3002/api/v1/searchUsers?searchQuery=John" -H "Authorization: Bearer $TOKEN"

# Update user (requires superadmin)
curl -X PUT http://localhost:3002/api/v1/auth/update \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","name":"New Name","role":"admin","department":"IT"}'

# Disable user (requires superadmin)
curl -X PUT http://localhost:3002/api/v1/auth/disable \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com"}'

# Enable user (requires superadmin)
curl -X PUT http://localhost:3002/api/v1/auth/enable \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com"}'

# Reset user password (requires superadmin)
curl -X PUT http://localhost:3002/api/v1/auth/resetpassword \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"newpassword123"}'

# Self-reset password (requires auth)
curl -X PUT http://localhost:3002/api/v1/auth/userResetPassword \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"email":"self@example.com","newPassword":"newpassword123"}'

# Submit incident management report (requires admin or manager)
curl -X POST http://localhost:3002/api/v1/incidents/1/management \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "impactOnService": "Minor delay",
    "contributoryFactors": "Understaffing",
    "actionsTakenOutcomes": "Patient stabilized",
    "recommendations": "Increase staffing",
    "lessonsLearned": "Need better protocols",
    "informedPatient": true,
    "informedRelative": false,
    "informedSeniorManager": true,
    "informedPharmacist": false,
    "policeIncidentNumber": "",
    "informedOther": "",
    "riskSeverity": 3,
    "riskLikelihood": 2,
    "riskRating": 6,
    "ohsAbsenceOver3Days": false,
    "ohsActOfViolenceOrDanger": false,
    "ohsHospitalizationOver24Hours": false,
    "ohsStaffName": "",
    "ohsStaffDob": "",
    "ohsStaffAddress": "",
    "managerName": "Jane Manager",
    "managerSignature": true,
    "managerDesignation": "Clinical Manager",
    "managerDate": "2026-06-09"
  }'

# Get incident management report (requires auth)
curl "http://localhost:3002/api/v1/incidents/1/management" -H "Authorization: Bearer $TOKEN"

# Update incident management report (requires admin or manager)
curl -X PUT http://localhost:3002/api/v1/incidents/1/management \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "impactOnService": "Minor delay",
    "contributoryFactors": "Understaffing",
    "actionsTakenOutcomes": "Patient stabilized",
    "recommendations": "Increase staffing",
    "lessonsLearned": "Need better protocols",
    "informedPatient": true,
    "informedRelative": false,
    "informedSeniorManager": true,
    "informedPharmacist": false,
    "riskSeverity": 3,
    "riskLikelihood": 2,
    "riskRating": 6,
    "ohsAbsenceOver3Days": false,
    "ohsActOfViolenceOrDanger": false,
    "ohsHospitalizationOver24Hours": false,
    "managerName": "Jane Manager",
    "managerSignature": true,
    "managerDesignation": "Clinical Manager",
    "managerDate": "2026-06-09"
  }'

# Get incident management logs (requires admin or manager)
curl "http://localhost:3002/api/v1/incidents/1/managementlogs" -H "Authorization: Bearer $TOKEN"

# Add comment to incident (requires admin or manager)
curl -X POST http://localhost:3002/api/v1/incidents/comments \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"incidentId": 1, "userId": 2, "comment": "Follow up needed"}'

# Get comments for incident (requires admin or manager)
curl "http://localhost:3002/api/v1/incidents/comments?incidentId=1" -H "Authorization: Bearer $TOKEN"

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

# Get all death reports (paginated, no auth required)
curl "http://localhost:3002/api/v1/deathreports?page=1&limit=10"

# Search death reports (no auth required; supports dateFrom/dateTo)
curl "http://localhost:3002/api/v1/searchDeathReport?searchQuery=DR-001"
```

## Role Permissions

| Role | Permissions |
|------|-------------|
| superadmin | All endpoints including user management (register, update, disable, enable, reset password, get user, search users), report incidents, view all incidents, update any incident status, submit incident management reports, update incident management reports, add comments, view comments |
| admin | Report incidents, view all incidents, update any incident status, submit incident management reports, update incident management reports, add comments, view comments |
| supervisor | Report incidents, view own department incidents (matched via `incident_ward_dept`, `patient_ward_dept`, or `staff_place_of_work`) |
| manager | Report incidents, view all incidents, view incident management reports and logs, add comments, view comments, submit incident management reports, update incident management reports |
| reporter | Report incidents via public endpoint only, view own department incidents |

**Notes:**
- Only `superadmin` and `admin` can update incident status (`PATCH /incidents/:id/status`). `supervisor`, `manager`, and `reporter` are forbidden.
- `GET /incidents` returns paginated results. `supervisor` and `reporter` are scoped to their department.
- `GET /incidents` supports `dateFrom` and `dateTo` query parameters for filtering.
- JWT tokens expire after 72 hours.

## Auth Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/v1/auth/register` | superadmin | Register new user |
| POST | `/api/v1/auth/login` | None | Login, returns JWT + user |
| PUT | `/api/v1/auth/update` | superadmin | Update user name/role/department |
| PUT | `/api/v1/auth/disable` | superadmin | Disable a user account |
| PUT | `/api/v1/auth/enable` | superadmin | Enable a disabled user account |
| PUT | `/api/v1/auth/resetpassword` | superadmin | Admin-reset any user password |
| PUT | `/api/v1/auth/userResetPassword` | authenticated | Self-service password change |

### Register Request

```go
type RegisterRequest struct {
    Name       string `json:"name" binding:"required"`
    Email      string `json:"email" binding:"required"`
    Password   string `json:"password" binding:"omitempty,min=8"`
    Role       string `json:"role" binding:"required"`
    Department string `json:"department" binding:"required"`
}
```

### Update Request

```go
type UpdateRequest struct {
    Name       string `json:"name" binding:"required"`
    Email      string `json:"email" binding:"required"`
    Role       string `json:"role" binding:"required"`
    Department string `json:"department" binding:"required"`
}
```

### Disable Request

```go
type DisableRequest struct {
    Email string `json:"email" binding:"required"`
}
```

### Enable Request

```go
type EnableRequest struct {
    Email string `json:"email" binding:"required"`
}
```

### Reset Request

```go
type ResetRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required,min=8"`
}
```

### User Self Reset Password

```go
type UserResetPassword struct {
    Email       string `json:"email" binding:"required"`
    NewPassword string `json:"newPassword" binding:"required"`
}
```

## Incident Report

### API Endpoints

**POST /api/v1/incidents** - Create incident report
- Requires: No authentication
- Request body: `IncidentReport` struct
- Returns: Created `IncidentReport` with generated `id`

**GET /api/v1/incidents** - List incidents (paginated)
- Requires: Authentication
- Query Parameters: `page` (default 1), `limit` (default 10, max 50), `dateFrom`, `dateTo`
- Returns: `PaginatedIncidentResponse`

**GET /api/v1/searchIncidents** - Search incidents
- Requires: Authentication
- Query Parameters: `searchQuery`
- Returns: `{"incidents": [...]}`

**PATCH /api/v1/incidents/:id/status** - Update incident status
- Requires: `admin` or `superadmin` role
- Request body: `IncidentStatusUpdate`

### IncidentReport Struct

```go
type IncidentReport struct {
    Id                     int            `json:"id"`
    PrincipalName          string         `json:"principalName"`
    PrincipalGender        string         `json:"principalGender"`
    PrincipalDob           string         `json:"principalDob"`
    PrincipalType          string         `json:"principalType"`
    PatientId              string         `json:"patientId,omitempty"`
    PatientWardDept        string         `json:"patientWardDept,omitempty"`
    StaffJobTitle          string         `json:"staffJobTitle,omitempty"`
    StaffPhone             string         `json:"staffPhone,omitempty"`
    StaffPlaceOfWork       string         `json:"staffPlaceOfWork,omitempty"`
    StaffSite              string         `json:"staffSite,omitempty"`
    PeopleInvolved         string         `json:"peopleInvolved"`
    DateOfIncident         string         `json:"dateOfIncident"`
    TimeOfIncident         string         `json:"timeOfIncident"`
    LocationOfIncident     string         `json:"locationOfIncident"`
    IncidentWardDept       string         `json:"incidentWardDept"`
    Witnesses              string         `json:"witnesses,omitempty"`
    WitnessType            string         `json:"witnessType,omitempty"`
    WitnessWardDept        string         `json:"witnessWardDept,omitempty"`
    WitnessJobTitle        string         `json:"witnessJobTitle,omitempty"`
    WitnessPhone           string         `json:"witenssPhone,omitempty"`
    IsNearMiss             bool           `json:"isNearMiss"`
    CauseGroup             string         `json:"causeGroup"`
    Causes                 string         `json:"causes"`
    PrescribingDoctor      string         `json:"prescribingDoctor"`
    TreatmentReceived      string         `json:"treatmentReceived"`
    EquipmentInvolved      string         `json:"equipmentInvolved"`
    EquipmentModel         string         `json:"equipmentModel,omitempty"`
    EquipmentSentForRepair bool           `json:"equipmentSentForRepair"`
    EquipmentWithdrawn     bool           `json:"equipmentWithdrawn"`
    EquipmentRetained      bool           `json:"equipmentRetained"`
    EquipmentNumber        string         `json:"equipmentNumber,omitempty"`
    IsMedicalDevice        string         `json:"isMedicalDevice,omitempty"`
    ReporterName           string         `json:"reporterName" binding:"required"`
    ReporterDesignation    string         `json:"reporterDesignation" binding:"required"`
    Signature              bool           `json:"signature" binding:"required"`
    ReporterInfo           string         `json:"reporterInfo" binding:"required"`
    ReporterDate           string         `json:"date" binding:"required"`
    SeverityLevel          SeverityLevel  `json:"severityLevel"`
    IncidentStatus         IncidentStatus `json:"incidentStatus"`
}
```

### Severity Levels

| Value | Description |
|-------|-------------|
| `near miss` | Near miss |
| `minor` | Minor |
| `major` | Major |
| `critical` | Critical |

### Incident Status

| Value | Description |
|-------|-------------|
| `unresolved` | Unresolved |
| `inprogress` | In Progress |
| `resolved` | Resolved |

### Incident Status Update

```go
type IncidentStatusUpdate struct {
    Status string `json:"incidentStatus" binding:"required"`
}
```

### Paginated Incident Response

```go
type PaginatedIncidentResponse struct {
    Data       []db.IncidentReport `json:"data"`
    Pagination PaginationMeta      `json:"pagination"`
}

type PaginationMeta struct {
    CurrentPage int `json:"current_page"`
    PageSize    int `json:"page_size"`
    TotalItems  int `json:"total_items"`
    TotalPages  int `json:"total_pages"`
}
```

## Death Report

The death report feature captures detailed information about workplace deaths. These endpoints are public and do not require authentication.

### API Endpoints

**POST /api/v1/deathreport** - Create death report
- Requires: No authentication
- Request body: `DeathReport` struct
- Returns: `{"message": "The death has been reported"}`

**PUT /api/v1/deathreport** - Update death report
- Requires: No authentication
- Request body: `DeathReport` struct (must include `id`)
- Returns: `{"message": "The death report has been updated"}`

**GET /api/v1/deathreports** - Get all death reports (paginated)
- Requires: No authentication
- Query Parameters: `page` (default 1), `limit` (default 10, max 50), `dateFrom`, `dateTo`
- Returns: `{"deathReports": PaginatedDeathReportResponse}`

**GET /api/v1/searchDeathReport** - Search death reports
- Requires: No authentication
- Query Parameters: `searchQuery`, `dateFrom`, `dateTo`
- Returns: `{"deathReports": [...]}`

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
- Requires: `admin` or `manager` role
- Request body: `IncidentManagement` struct

**GET /api/v1/incidents/:id/management** - Retrieve management report
- Requires: Authentication (any authenticated user)
- Response: `IncidentManagement` struct

**PUT /api/v1/incidents/:id/management** - Update existing report
- Requires: `manager` or `admin` role
- Request body: `IncidentManagement` struct

**GET /api/v1/incidents/:id/managementlogs** - Retrieve change logs for a management report
- Requires: `admin` or `manager` role
- Response: `[]IncidentManagementLogs`

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

### IncidentManagementLogs Struct

```go
type IncidentManagementLogs struct {
    Id         int                `json:"id"`
    IncidentId int                `json:"incidentId"`
    ChangedBy  int                `json:"changedBy"`
    Action     string             `json:"action"`
    OldValue   IncidentManagement `json:"oldValue"`
    NewValue   IncidentManagement `json:"newValue"`
    CreatedAt  time.Time          `json:"createdAt"`
    UserName   string             `json:"userName"`
}
```

## Comments

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/api/v1/incidents/comments` | `admin` or `manager` | Add a comment |
| GET | `/api/v1/incidents/comments?incidentId=1` | `admin` or `manager` | Get comments for incident |

### Comment Struct

```go
type Comment struct {
    Id              int    `json:"id"`
    IncidentId      int    `json:"incidentId" binding:"required"`
    UserId          int    `json:"userId" binding:"required"`
    Comment         string `json:"comment" binding:"required"`
    CommentUserName string `json:"commentUserName"`
    CommentUserRole string `json:"commentUserRole"`
}
```

## Users

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/v1/users?page=1&limit=10` | `superadmin` | List all users (paginated) |
| GET | `/api/v1/user?email=test@example.com` | `superadmin` | Get single user by email |
| GET | `/api/v1/searchUsers?searchQuery=John` | `superadmin` | Search users |

### User Struct

```go
type User struct {
    Id         int    `json:"id"`
    Name       string `json:"name"`
    Email      string `json:"email"`
    Role       string `json:"role"`
    Department string `json:"department"`
    Disabled   bool   `json:"disabled"`
}
```

## Default Credentials

A superadmin user is created by default:
- Email: `admin@example.com`
- Password: `redeemershealthvillage`
- Bcrypt hash stored in `tables.sql`: `$2a$10$UQgnunKYIsM.hTWtjYooG.SPNKBqywEbOKddh1tU4tJuDiqfcn5Dm`

**Note:** New users registered via `/api/v1/auth/register` are assigned a default password of `redeemershealthvillage` if none is provided. This is separate from the pre-seeded superadmin.
