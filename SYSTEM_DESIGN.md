# Issue Tracker - System Design

**Code Metrics:** 3646 lines of Go, 28 source files

## System Overview

The Issue Tracker is a stateless RESTful API built with Go that provides incident tracking capabilities with role-based access control. The system follows a layered architecture pattern with clear separation between presentation, application, and data layers.

Go version: 1.26+

## Architecture Diagram

```mermaid
flowchart TD
    Web[Web Client] --> LB["LOAD BALANCER<br/>Optional for production"]
    Mobile[Mobile App] --> LB
    Consumer[API Consumer] --> LB
    LB --> Gin[Gin Web Framework]
    
    subgraph Router["Router & Middleware"]
        CORS[CORS]
        JWT[JWT Auth]
        Rate[Rate Limiting]
    end
    
    subgraph Routes["Route Groups"]
        Ping["/api/v1/ping"]
        Auth["/api/v1/auth/*<br/>register, login, update, disable<br/>enable, resetpassword, userResetPassword"]
        Incidents["/api/v1/incidents<br/>POST public, GET auth"]
        IncidentsStatus["/api/v1/incidents/:id/status"]
        Management["/api/v1/incidents/:id/management<br/>POST, GET, PUT"]
        ManagementLogs["/api/v1/incidents/:id/managementlogs"]
        Comments["/api/v1/incidents/comments<br/>POST, GET"]
        Users["/api/v1/users"]
        User["/api/v1/user"]
        SearchU["/api/v1/searchUsers"]
        SearchI["/api/v1/searchIncidents"]
        DeathReport["/api/v1/deathreport"]
        DeathReportUpdate["/api/v1/deathreport/:id"]
        DeathSearch["/api/v1/searchDeathReport"]
    end
    
    subgraph Handlers["Handlers"]
        AuthGo[auth.go]
        IncidentsGo[incidents.go]
        UsersGo[users.go]
        CommentsGo[comments.go]
        MgmtGo[incidentmanagement.go]
        UtilsGo[utils.go]
    end
    
    App[Application Layer<br/>type application struct]
    
    Gin --> Router --> Routes --> Handlers --> App
    
    Models[db.Models]
    UsersM[users.go]
    IncidentsM[incidents.go]
    MgmtM[incidentmanagement.go]
    CommentsM[comments.go]
    DBInit[db.go]
    
    App --> Models
    Models --> UsersM
    Models --> IncidentsM
    Models --> MgmtM
    Models --> CommentsM
    Models --> DBInit
    
    PG[(PostgreSQL 16)] --> Env[env.go typed accessors]
    Docker[docker-compose.yml] --> Scripts[scripts/]
    Tables[tables.sql] --> PG
```

## Component Interaction Flow

```mermaid
flowchart LR
    Client[Client] --> API[API]
    API --> DB[Database]
    API --> Middleware[Middleware JWT Auth]
    Middleware --> Handler[Handler Business Logic]
    Handler --> Model[Model Data Access]
    Model --> DB
```

## Data Flow Sequence

```mermaid
sequenceDiagram
    participant C as Client
    participant G as Gin Router
    participant M as Middleware
    participant H as Handler login
    participant D as Model Layer
    participant P as PostgreSQL
    
    C->>G: HTTP Request<br/>POST /api/v1/auth/login
    G->>M: matches route
    M->>H: none for login
    H->>H: 1. Validate
    H->>D: 2. GetByEmail()
    D->>P: Query: SELECT WHERE email=$1
    P-->>D: Query Results
    D-->>H: User Model
    H->>H: 3. Verify pass
    H->>H: 4. Create JWT
    H-->>C: Response {token, user}
```

## Request-Response Flow

```mermaid
flowchart TD
    Req["REQUEST"] --> CORS["1. CORS Middleware"]
    CORS --> Route["2. Route Matching"]
    Route --> Auth["3. Auth Middleware<br/>- Extract Bearer token<br/>- Validate JWT signature<br/>- Verify expiration<br/>- Set user context"]
    Auth --> Handler["4. Handler Execution<br/>- Input validation<br/>- Role-based authorization<br/>- Business logic"]
    Handler --> Serialize["5. Response Serialization"]
    Serialize --> Pool["Connection Pool<br/>Min: 2, Max: 10"]
    Pool --> PGX["PGX Driver<br/>- Parameterized queries<br/>- Connection pooling<br/>- Row scanning to structs"]
    PGX --> Res["RESPONSE"]
```

## System Components

### 1. Presentation Layer (`cmd/`)
- **Routes**: HTTP endpoint definitions with middleware chains
- **Handlers**: Request processing and response formatting
- **Types**: Request/response DTOs and domain types

### 2. Application Layer
- **Business Logic**: Implemented in handlers
- **Authorization**: Role-based access control
- **Validation**: Input validation using Gin binding

### 3. Data Access Layer (`internal/db/`)
- **Models**: Database interaction logic
- **Connection Pool**: PGX connection management
- **Queries**: Parameterized SQL operations

### 4. Infrastructure Layer
- **Database**: PostgreSQL with connection pooling
- **Configuration**: Environment variables
- **Deployment**: Docker Compose (PostgreSQL + Server containers)
- **Scripts**:
  - `login.sh` → Access DB shell
  - `commit.sh` → Git helper

## Security Architecture

```mermaid
flowchart LR
    Client[Client Password] --> Plain[Password plain]
    Plain --> Bcrypt[Bcrypt Hash]
    Bcrypt --> Login[Login Request]
    Login --> Claims[JWT Claims]
    Claims --> Token[Signed Token]
    Token --> Extract[Extract Token]
    Extract --> Validate[Validate Signature]
    Validate --> Verify[Verify Claims]
    Verify --> Check[Check Role]
    Check --> Access{Access Granted?}
```

## Role Hierarchy

| Role | Permissions |
|------|-------------|
| **superadmin** | User management (register, update, disable/enable, reset password, get user), report incidents, view all incidents, update any incident status, add comments, view comments, submit and update incident management reports, view incident management reports and logs |
| **admin** | Report incidents, view all incidents, update any incident status, add comments, view comments, submit and update incident management reports, view incident management reports and logs |
| **supervisor** | Report incidents, view own department incidents (via `incident_ward_dept`) |
| **manager** | Report incidents, view all incidents, view incident management reports and logs, add comments, view comments, submit incident management reports, update incident management reports |
| **reporter** | Report incidents via public endpoint only, view own department incidents |

## Deployment Architecture

### Development

```mermaid
flowchart TD
    Local[Local Machine] --> DockerCompose[Docker Compose]
    DockerCompose --> App[Go Application<br/>- Port: 3001<br/>- Hot reload via Air<br/>- Volume: .:/app<br/>- Excludes: scripts/ directory]
```

### Production

```mermaid
flowchart TD
    Internet[Internet Traffic] --> LB[Load Balancer<br/>NGINX/HAProxy]
    LB --> API1[API Instance 1<br/>Go + Gin]
    LB --> API2[API Instance 2<br/>Go + Gin]
    API1 --> Primary[Primary<br/>RW]
    API2 --> Primary
    Primary --> Replica1[Replica 1<br/>RO]
    Primary --> Replicas[Replicas ...]
```

## Performance Characteristics

| Metric | Value |
|--------|-------|
| Max Connections | 10 |
| JWT Expiration | 72 hours |
| Request Timeout | 1s read, 5s write |
| Idle Timeout | 30s |
| Pagination Limit | Max 50 items |

## Current Implementation Status

**Implemented:**
- ✅ Clean layered architecture (presentation → application → data → infrastructure)
- ✅ JWT authentication with 72-hour expiry
- ✅ Role-based access control (superadmin, admin, supervisor, manager, reporter)
- ✅ Department-scoped incident access
- ✅ Incident management follow-up reports
- ✅ Incident comments
- ✅ Structured logging (`internal/logger`)
- ✅ Health check endpoint (`/api/v1/ping`)
- ✅ CORS configuration
- ✅ Database connection pooling
- ✅ Docker Compose setup
- ✅ Unit tests for routes and handlers

**Pending:**
- ❌ Unit tests (increase coverage beyond current partial implementation)
- ❌ Rate limiting
- ❌ Error handling package
- ❌ Configuration validation

## Scalability Considerations

1. **Horizontal Scaling**: API instances can be scaled behind a load balancer
2. **Database Scaling**: Connection pooling, read replicas for reporting
3. **Caching**: Redis layer can be added for frequently accessed data
4. **Stateless**: JWT enables horizontal scaling without session storage