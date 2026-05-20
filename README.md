# DoFocus Backend

A scalable productivity and focus management backend built using Go, Gin, PostgreSQL, JWT Authentication, and GORM.

DoFocus combines:

* Task Management
* Focus/Pomodoro Sessions
* User Authentication
* Productivity Tracking

The backend is designed using a clean layered architecture to make the project scalable, maintainable, and production-ready.

---

# Project Goal

The goal of DoFocus is to help users:

* Manage tasks efficiently
* Track focus sessions
* Improve productivity
* Analyze work patterns in the future

This backend is designed to support:

* Web applications
* Mobile applications
* Future microservices/extensions

---

# Tech Stack

## Backend

* Go
* Gin
* GORM
* PostgreSQL (NeonDB)
* JWT Authentication
* Gmail SMTP

## Frontend

* React

---

# Features Implemented

## Authentication

* User Registration
* OTP Email Verification
* Login using JWT
* Password Hashing using bcrypt
* Protected APIs using JWT Middleware

---

## Task Management

* Create Task
* Get User Tasks
* Update Task
* Delete Task

---

## Security Features

* JWT Token Authentication
* Password Hashing
* Protected Routes
* User-specific Task Isolation

---

# Future Planned Features

* Pomodoro Timer
* Focus Sessions Tracking
* Productivity Analytics
* Task Categories
* Daily/Weekly Reports
* Notifications
* Refresh Tokens
* Docker Deployment
* CI/CD Pipeline

---

# Project Architecture

The backend follows a layered architecture.

```txt
Request
   ↓
Routes
   ↓
Handler
   ↓
Service
   ↓
Repository
   ↓
Database
```

---

## Why This Architecture?

This architecture is used because it:

* Separates responsibilities properly
* Makes code maintainable
* Makes debugging easier
* Allows scaling the application
* Makes testing easier
* Follows industry backend standards

---

# Folder Responsibilities

## cmd/api

Contains application entry point.

```txt
main.go
```

Responsible for:

* Starting server
* Initializing database
* Registering routes
* Applying middleware

---

## internal/routes

Responsible for:

* Registering API routes
* Grouping routes
* Applying middleware

Example:

```txt
/api/v1/auth
/api/v1/task
```

---

## internal/handler

Responsible for:

* Handling HTTP requests
* Validating request body
* Sending HTTP responses

Handlers should NOT contain database logic.

---

## internal/service

Responsible for:

* Business logic
* Application rules
* Workflow handling

Example:

* Register user flow
* Login flow
* Task update logic

---

## internal/repository

Responsible for:

* Database queries
* Database operations
* CRUD queries

Repository layer directly communicates with PostgreSQL.

---

## internal/models

Responsible for:

* Database table structures
* GORM models
* Relationships

---

## internal/middleware

Responsible for:

* JWT authentication
* Protected route verification
* Request filtering

---

## internal/database

Responsible for:

* Database connection setup
* Auto migrations

---

## internal/util

Responsible for:

* Reusable helper functions
* JWT generation
* Email sending
* Password utilities

---

# Current Project Structure

```txt
dofocus-backend/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   │
│   ├── database/
│   │   └── database.go
│   │
│   ├── handler/
│   │   ├── auth_handler.go
│   │   └── task_handler.go
│   │
│   ├── middleware/
│   │   └── auth_middleware.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   ├── otp_verification.go
│   │   └── task.go
│   │
│   ├── repository/
│   │   ├── auth_repository.go
│   │   └── task_repository.go
│   │
│   ├── routes/
│   │   └── auth_routes.go
│   │
│   ├── service/
│   │   ├── auth_service.go
│   │   └── task_service.go
│   │
│   └── util/
│       ├── email.go
│       ├── jwt.go
│       └── password.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

# Database Design

Currently the backend contains 3 tables.

---

## users

Stores permanent user information.

### Fields

* user_id
* name
* email
* password
* created_at
* updated_at

---

## otp_verifications

Stores temporary OTP verification data.

### Fields

* id
* email
* otp
* verified
* expires_at
* created_at
* updated_at

### Purpose

Used for:

* Email verification
* OTP validation
* Registration security

After successful registration:

```txt
OTP verification row is deleted
```

---

## tasks

Stores user tasks.

### Fields

* id
* title
* completed
* user_id
* created_at
* updated_at

### Purpose

Each task belongs to a specific user.

This is implemented using:

```txt
user_id
```

---

# Future Database Table

## focus_sessions (planned)

This table will track:

* Pomodoro sessions
* Focus duration
* Productivity analytics

### Planned Fields

* id
* user_id
* task_id
* start_time
* end_time
* duration
* completed
* created_at

---

# Authentication Flow

## Registration Flow

```txt
User enters email
    ↓
Send OTP
    ↓
OTP stored in database
    ↓
OTP sent through Gmail
    ↓
User verifies OTP
    ↓
verified = true
    ↓
User registers
    ↓
Password hashed
    ↓
User stored in users table
    ↓
OTP row deleted
```

---

## Login Flow

```txt
User enters email + password
    ↓
Backend verifies credentials
    ↓
JWT token generated
    ↓
JWT token returned to frontend
```

---

# JWT Authentication

After login:

Frontend stores JWT token.

For protected requests:

```http
Authorization: Bearer <token>
```

Middleware verifies token before allowing access.

---

# Protected Routes

All task APIs are protected.

Example:

```txt
/api/v1/task
```

JWT middleware:

* Verifies token
* Extracts user_id
* Allows request

---

# Task Ownership Security

Users can ONLY access their own tasks.

Database queries always filter using:

```txt
task_id + user_id
```

This prevents unauthorized access.

Example:

```txt
User A cannot update/delete User B tasks
```

---

# API Endpoints

# Authentication APIs

---

## Send OTP

```http
POST /api/v1/auth/send-otp
```

### Request

```json
{
  "email": "user@gmail.com"
}
```

---

## Verify OTP

```http
POST /api/v1/auth/verify-otp
```

### Request

```json
{
  "email": "user@gmail.com",
  "otp": "123456"
}
```

---

## Register User

```http
POST /api/v1/auth/register
```

### Request

```json
{
  "name": "Rahul",
  "email": "user@gmail.com",
  "password": "password123"
}
```

---

## Login

```http
POST /api/v1/auth/login
```

### Request

```json
{
  "email": "user@gmail.com",
  "password": "password123"
}
```

### Response

```json
{
  "token": "jwt_token"
}
```

---

# Task APIs

All task APIs require JWT token.

---

## Create Task

```http
POST /api/v1/task
```

### Request

```json
{
  "title": "Complete backend"
}
```

---

## Get Tasks

```http
GET /api/v1/task
```

---

## Update Task

```http
PUT /api/v1/task/:id
```

### Request

```json
{
  "title": "Updated Task",
  "completed": true
}
```

---

## Delete Task

```http
DELETE /api/v1/task/:id
```

---

# Environment Variables

Create a `.env` file in project root.

---

## DATABASE_URL

```env
DATABASE_URL=
```

Used to connect Go backend with PostgreSQL database.

---

## JWT_SECRET

```env
JWT_SECRET=
```

Used for:

* JWT token signing
* JWT token verification

This should always remain secret.

---

## EMAIL_USER

```env
EMAIL_USER=
```

Gmail account used for sending OTP emails.

---

## EMAIL_PASSWORD

```env
EMAIL_PASSWORD=
```

Google App Password used for SMTP authentication.

This is NOT the normal Gmail password.

---

# Project Setup Instructions

## 1. Clone Repository

```bash
git clone <repository-url>
```

Downloads project source code to local machine.

---

## 2. Move Into Project

```bash
cd dofocus-backend
```

Moves terminal into project directory.

---

## 3. Install Dependencies

```bash
go mod tidy
```

Downloads and synchronizes all Go project dependencies.

Example:

* Gin
* GORM
* JWT library
* PostgreSQL driver

---

## 4. Create `.env`

Create environment file:

```env
DATABASE_URL=
JWT_SECRET=
EMAIL_USER=
EMAIL_PASSWORD=
```

This stores sensitive credentials securely.

---

## 5. Run Application

```bash
go run cmd/api/main.go
```

Starts backend server.

Default server:

```txt
http://localhost:8080
```

---

# Security Design

## Password Security

Passwords are hashed using:

```txt
bcrypt
```

Plain passwords are NEVER stored in database.

---

## JWT Security

Protected APIs require valid JWT token.

---

## Route Protection

Middleware blocks unauthorized requests.

---

## User Isolation

Users can only access their own data.

---

# Git Workflow

Recommended branch workflow:

```txt
main
feature/auth
feature/task-crud
```

Feature branches should be merged into main after testing.

---

# Future Improvements

* Refresh Tokens
* Redis OTP Storage
* WebSocket Timer Sync
* Focus Session Analytics
* Productivity Reports
* Docker Support
* CI/CD Pipelines
* Unit Testing
* Rate Limiting
* Email Templates
* Background Jobs

---

# Author

Developed as part of the DoFocus productivity platform backend system.


* temp-main has code till: 
    ```
    commit f77b45ff4754ae6b9e8edf09b517e6421fb70db3 (HEAD -> main, origin/main, temp-main)
    Author: Rahul Badachi <rahulsbadachi052@gmail.com>
    Date:   Tue May 19 15:01:04 2026 +0530

    Implement protected task CRUD APIs and project documentation
    ```
