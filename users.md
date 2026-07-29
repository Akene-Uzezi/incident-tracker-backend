# User Routes Documentation

## Overview

This document catalogs every API route that deals with users in the Issue Tracker backend. All endpoints are prefixed with `/api/v1` unless otherwise noted.

---

## Routes

### 1. Register a New User

| Attribute | Value |
|-----------|-------|
| **Method** | `POST` |
| **Path** | `/api/v1/auth/register` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/auth.go:13` (`register`) |

**Input Body** (`RegisterRequest`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `name` | string | Yes | Required |
| `email` | string | Yes | Required |
| `password` | string | No | Minimum 8 characters (optional; defaults to `redeemershealthvillage`) |
| `role` | string | Yes | Required; must be one of: `reporter`, `supervisor`, `admin`, `superadmin`, `manager` |
| `department` | string | Yes | Required |

**Output**:

| Status | Response |
|--------|----------|
| `201 Created` | `{"user": {id, name, email, role, department}}` |
| `400 Bad Request` | `{"error": "..."}` — validation or role error |
| `403 Forbidden` | `{"error": "Unauthorized. Must be a superadmin"}` |
| `409 Conflict` | `{"error": "user already exists"}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 2. Login

| Attribute | Value |
|-----------|-------|
| **Method** | `POST` |
| **Path** | `/api/v1/auth/login` |
| **Auth Required** | No |
| **Role Required** | None (public endpoint) |
| **Handler** | `cmd/auth.go:84` (`login`) |

**Input Body** (`loginRequest`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `email` | string | Yes | Required |
| `password` | string | Yes | Required |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{"token": "<jwt>", "user": {id, name, email, role, department}}` |
| `400 Bad Request` | `{"error": "..."}` — validation error |
| `404 Not Found` | `{"error": "user not found"}` |
| `403 Forbidden` | `{"error": "This account has been disabled"}` |
| `401 Unauthorized` | `{"error": "Invalid Credentials"}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 3. Update User

| Attribute | Value |
|-----------|-------|
| **Method** | `PUT` |
| **Path** | `/api/v1/auth/update` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/users.go:11` (`update`) |

**Input Body** (`UpdateRequest`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `name` | string | Yes | Required |
| `email` | string | Yes | Required |
| `role` | string | Yes | Required |
| `department` | string | Yes | Required |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{"user": {updated user}}` |
| `400 Bad Request` | `{"error": "..."}` — validation error |
| `403 Forbidden` | `{"error": "Unauthorized. Must be a superadmin"}` |
| `404 Not Found` | `{"error": "user not found"}` |
| `500 Internal Server Error` | `{"error": "Failed to perform database query"}` |

---

### 4. Disable User

| Attribute | Value |
|-----------|-------|
| **Method** | `PUT` |
| **Path** | `/api/v1/auth/disable` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/users.go:44` (`disable`) |

**Input Body** (`DisableRequest`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `email` | string | Yes | Required |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{user object with disabled: true}` |
| `400 Bad Request` | `{"error": "..."}` — validation error |
| `403 Forbidden` | `{"error": "Unauthorized. Must be a superadmin"}` |
| `404 Not Found` | `{"error": "user not found"}` |
| `500 Internal Server Error` | `{"error": "Failed to perform database query"}` |

---

### 5. Enable User

| Attribute | Value |
|-----------|-------|
| **Method** | `PUT` |
| **Path** | `/api/v1/auth/enable` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/users.go:74` (`enable`) |

**Input Body** (`EnableRequest`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `email` | string | Yes | Required |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{user object with disabled: false}` |
| `400 Bad Request` | `{"error": "..."}` — validation error |
| `403 Forbidden` | `{"error": "Unauthorized must be a super admin"}` |
| `404 Not Found` | `{"error": "user not found"}` |
| `500 Internal Server Error` | `{"error": "failed to perform database query"}` |

---

### 6. Reset User Password (Admin)

| Attribute | Value |
|-----------|-------|
| **Method** | `PUT` |
| **Path** | `/api/v1/auth/resetpassword` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/auth.go:132` (`resetPassword`) |

**Input Body** (`ResetRequest`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `email` | string | Yes | Required |
| `password` | string | Yes | Required, minimum 8 characters |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{updated user object}` |
| `400 Bad Request` | `{"error": "Invalid request"}` |
| `403 Forbidden` | `{"error": "Only super admins are allowed to update a user"}` |
| `404 Not Found` | `{"error": "user not found"}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 7. Reset Own Password (Self-Service)

| Attribute | Value |
|-----------|-------|
| **Method** | `PUT` |
| **Path** | `/api/v1/auth/userResetPassword` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | Any authenticated role (self-service only) |
| **Handler** | `cmd/users.go:156` (`userResetPassword`) |

**Input Body** (`UserResetPassword`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `email` | string | Yes | Required |
| `newPassword` | string | Yes | Required |

**Constraints**: The authenticated user's email (from JWT claims) must match the `email` field in the request body. A user cannot reset another user's password.

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{"message": "password changed successfully"}` |
| `400 Bad Request` | `{"error": "Invalid Request"}` |
| `403 Forbidden` | `{"error": "you are not allowed to reset another users password"}` |
| `500 Internal Server Error` | `{"error": "failed to hash new password"}` or database error |

---

### 8. Get All Users

| Attribute | Value |
|-----------|-------|
| **Method** | `GET` |
| **Path** | `/api/v1/users` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/users.go:124` (`getUsers`) |

**Query Parameters**:

| Parameter | Type | Required | Default | Validation |
|-----------|------|----------|---------|------------|
| `page` | int | No | `1` | Minimum 1 |
| `limit` | int | No | `10` | Minimum 1, maximum 50 |

**Output** (`PaginatedUserResponse`):

| Status | Response |
|--------|----------|
| `200 OK` | `{"data": [{users}], "pagination": {current_page, page_size, total_items, total_pages}}` |
| `403 Forbidden` | `{"error": "only super admins can fetch all users"}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 9. Get User by Email

| Attribute | Value |
|-----------|-------|
| **Method** | `GET` |
| **Path** | `/api/v1/user` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/users.go:104` (`getUser`) |

**Query Parameters**:

| Parameter | Type | Required | Validation |
|-----------|------|----------|------------|
| `email` | string | Yes | Non-empty string |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{user object}` |
| `400 Bad Request` | `{"error": "An email address must be passed into the query"}` |
| `403 Forbidden` | `{"error": "You are not allowed to access users"}` |
| `404 Not Found` | `{"error": "user not found"}` |
| `500 Internal Server Error` | `{"error": "Unable to get user"}` |

---

### 10. Search Users

| Attribute | Value |
|-----------|-------|
| **Method** | `GET` |
| **Path** | `/api/v1/searchUsers` |
| **Auth Required** | Yes (`authMiddleware`) |
| **Role Required** | `superadmin` |
| **Handler** | `cmd/users.go:184` (`searchUsers`) |

**Query Parameters**:

| Parameter | Type | Required | Validation |
|-----------|------|----------|------------|
| `searchQuery` | string | No | Non-empty string triggers search; empty string returns all users |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{"users": [{user objects}]}` |
| `403 Forbidden` | `{"error": "Only superadmins are allowed access to this route"}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

## Death Report Routes

The death report feature is public and does not require authentication.

### 1. Create Death Report

| Attribute | Value |
|-----------|-------|
| **Method** | `POST` |
| **Path** | `/api/v1/deathreport` |
| **Auth Required** | No |
| **Role Required** | None (public endpoint) |
| **Handler** | `cmd/deathreport.go:13` (`deathReport`) |

**Input Body** (`DeathReport`):

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| `ref` | string | Yes | Required |
| `reportedDate` | string | Yes | Required |
| `incidentDate` | string | Yes | Required |
| `incidentTime` | string | Yes | Required |
| `department` | string | Yes | Required |
| `location` | string | Yes | Required |
| `category` | string | Yes | Required |
| `subCategory` | string | Yes | Required |
| `description` | string | Yes | Required |
| `actionTaken` | string | Yes | Required |
| `openedDate` | string | No | Optional |
| `submittedTime` | string | No | Optional |
| `handler` | string | No | Optional |
| `manager` | string | No | Optional |
| `specialty` | string | No | Optional |
| `exactLocation` | string | No | Optional |
| `coding` | string | No | Optional |
| `type` | string | No | Optional |
| `riskGrading` | string | No | Optional |
| `result` | string | No | Optional |
| `actualHarm` | string | No | Optional |
| `potentialHarm` | string | No | Optional |
| `details` | string | No | Optional |
| `patientInvolved` | bool | No | Optional |
| `patientTold` | bool | No | Optional |
| `familyTold` | bool | No | Optional |
| `whatFamilyTold` | string | No | Optional |
| `incidentInvestigation` | string | No | Optional |
| `reviewMeetingDate` | string | No | Optional |
| `qualityAssuranceLead` | string | No | Optional |
| `doctorNotified` | bool | No | Optional |
| `meetingDiscussionPoints` | string | No | Optional |
| `meetingActionPoints` | string | No | Optional |
| `levelOfInvestigation` | string | No | Optional |

**Output**:

| Status | Response |
|--------|----------|
| `201 Created` | `{"message": "The death has been reported"}` |
| `400 Bad Request` | `{"error": "A bad request was sent"}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 2. Update Death Report

| Attribute | Value |
|-----------|-------|
| **Method** | `PUT` |
| **Path** | `/api/v1/deathreport` |
| **Auth Required** | No |
| **Role Required** | None (public endpoint) |
| **Handler** | `cmd/deathreport.go:29` (`updateDeathReport`) |

**Input Body** (`DeathReport` — must include `id`):

Same fields as Create Death Report, plus `id` (required) to identify the report to update.

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{"message": "The death report has been updated"}` |
| `400 Bad Request` | `{"error": "A bad request was passed: ..."}` |
| `404 Not Found` | `{"error": "..."}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 3. Get All Death Reports

| Attribute | Value |
|-----------|-------|
| **Method** | `GET` |
| **Path** | `/api/v1/deathreports` |
| **Auth Required** | No |
| **Role Required** | None (public endpoint) |
| **Handler** | `cmd/deathreport.go:53` (`getDeathReports`) |

**Query Parameters**:

| Parameter | Type | Required | Default | Validation |
|-----------|------|----------|---------|------------|
| `page` | int | No | `1` | Minimum 1 |
| `limit` | int | No | `10` | Minimum 1, maximum 50 |

**Output** (`PaginatedDeathReportResponse`):

| Status | Response |
|--------|----------|
| `200 OK` | `{"deathReports": {data: [...], pagination: {current_page, page_size, total_items, total_pages}}}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

### 4. Search Death Reports

| Attribute | Value |
|-----------|-------|
| **Method** | `GET` |
| **Path** | `/api/v1/searchDeathReport` |
| **Auth Required** | No |
| **Role Required** | None (public endpoint) |
| **Handler** | `cmd/deathreport.go:81` (`searchDeathReport`) |

**Query Parameters**:

| Parameter | Type | Required | Validation |
|-----------|------|----------|------------|
| `searchQuery` | string | No | Non-empty string triggers search; empty string returns all reports |

**Output**:

| Status | Response |
|--------|----------|
| `200 OK` | `{"deathReports": [{death report objects}]}` |
| `500 Internal Server Error` | `{"error": "..."}` |

---

## Authentication Middleware

All protected user routes (except login) are guarded by `authMiddleware` (`cmd/middleware.go:12`). It:

1. Extracts the `Authorization: Bearer <token>` header
2. Validates the JWT signature using the configured secret
3. Verifies token expiration
4. Sets user context values in the Gin context:
   - `userId`
   - `userRole`
   - `userEmail`
   - `userDepartment`

Handlers then check `userRole` to enforce role-based access.

---

## Related Files

| File | Purpose |
|------|---------|
| `cmd/routes.go` | Route registration |
| `cmd/auth.go` | Register and login handlers |
| `cmd/users.go` | Update, disable, enable, get, and reset password handlers |
| `cmd/middleware.go` | JWT authentication middleware |
| `cmd/types.go` | Request/response DTOs (`RegisterRequest`, `UpdateRequest`, `DisableRequest`, `EnableRequest`, `ResetRequest`, `loginRequest`, `Claims`, `UserResetPassword`, `PaginatedUserResponse`, `PaginatedDeathReportResponse`) |
| `cmd/deathreport.go` | Death report handlers (`deathReport`, `updateDeathReport`, `getDeathReports`, `searchDeathReport`) |
| `internal/db/deathreporting.go` | Death report database model and queries |
| `internal/db/users.go` | Database model and queries for users |
