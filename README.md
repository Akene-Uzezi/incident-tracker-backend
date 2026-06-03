# Issue Tracker

A RESTful API for tracking incidents/issues built with Go, Gin, and PostgreSQL.

## Features

- User authentication (registration)
- Incident reporting and tracking
- Role-based access control (implied from User model)
- Docker Compose setup for easy development

## Technologies

- Go 1.22+ (or as per go.mod)
- Gin web framework
- PostgreSQL database
- PGX PostgreSQL driver

## Getting Started

### Prerequisites

- Docker and Docker Compose (for containerized setup)
- Go (if running locally)

### Running with Docker Compose

1. Clone the repository
2. Copy `.env.example` to `.env` (if exists) or create one based on the environment variables below
3. Run `docker-compose up -d` to start the PostgreSQL container
4. Run `go run ./cmd/main.go` to start the application

Alternatively, you can use the provided scripts:

- `./createtables.sh` - creates necessary database tables (if not already created)
- `./login.sh` - (purpose unknown, possibly for db access)

### Environment Variables

The following environment variables are used:

- `PORT`: The port on which the server runs (default: 3001)
- `dbConnStr`: PostgreSQL connection string (default: `postgres://tracker_user:tracker_password@localhost:5432/issuetracker`)

These can be set in a `.env` file or exported in the shell.

## API Endpoints

### Health Check

- `GET /api/v1/ping` - Returns a pong message

### Authentication

- `POST /api/v1/auth/register` - Register a new user
  - Request body: `{ "email": "string", "name": "string", "password": "string" }`
  - Password must be at least 8 characters

### Incidents

_(Endpoints not fully implemented in the provided code, but the model exists)_

## Database Schema

The application uses two main tables:

1. `users` - stores user information (id, name, email, password hash, role)
2. `incidents` - stores incident reports (see Incident struct in internal/db/incidents.go for fields)

## Scripts

- `commit.sh` - Helper script for committing changes (custom)
- `createtables.sh` - Creates database tables (likely runs SQL migrations)
- `login.sh` - (Unclear purpose, possibly for database access)

## Project Structure

```
.
├── cmd/                - Application entrypoint and route handlers
│   ├── auth.go         - Authentication handlers
│   ├── main.go         - Application initialization and server startup
│   ├── routes.go       - Route definitions
│   └── server.go       - Server configuration
├── internal/           - Private application libraries
│   ├── db/             - Database models and connection pooling
│   │   ├── db.go       - Database connection and model initialization
│   │   ├── incidents.go - Incident model and severity levels
│   │   └── users.go    - User model
│   └── env/            - Environment variable helpers
├── docker-compose.yml  - PostgreSQL service definition
├── go.mod              - Go module definition
├── go.sum              - Go module checksums
├── test.sql            - Sample SQL for table creation
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a pull request
