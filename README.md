# REST API - Event Management System

A robust RESTful API built with Go for managing events and user authentication. This project demonstrates clean architecture, JWT authentication, and RESTful principles.

## Features

- **User Authentication** (Register/Login with JWT)
- **Event Management** (Create, Read, Update, Delete events)
- **RESTful API Design** with proper HTTP status codes
- **JWT-based Authentication** with secure token handling
- **Clean Architecture** with separation of concerns
- **SQLite Database** for data persistence

## API Endpoints

### Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/auth/register` | Register a new user |
| POST | `/api/v1/auth/login` | Login and get JWT token |

### Events
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/events` | Create a new event (Auth required) |
| GET | `/api/v1/events` | Get all events |
| GET | `/api/v1/events/:id` | Get event by ID |
| PUT | `/api/v1/events/:id` | Update event (Auth required) |
| DELETE | `/api/v1/events/:id` | Delete event (Auth required) |

## Technologies Used

- **Go** (Golang) - Backend language
- **Gin** - HTTP web framework
- **GORM** - ORM library for database operations
- **SQLite** - Database
- **JWT** - JSON Web Tokens for authentication

## Installation

### Prerequisites
- Go 1.19 or higher
- Git

### Steps
1. Clone the repository:
```bash
git clone https://github.com/yourusername/rest-api.git
cd rest-api
Install dependencies:

bash
go mod download
Run the application:

bash
go run main.go
The API will start on http://localhost:3030

Usage Examples
Register a User
bash
curl -X POST http://localhost:3030/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
Login
bash
curl -X POST http://localhost:3030/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
Create Event (with JWT)
bash
curl -X POST http://localhost:3030/api/v1/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Go Conference",
    "description": "A conference about Go programming",
    "date": "2025-05-20T10:00:00Z",
    "location": "San Francisco"
  }'
Project Structure
text
rest-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── database/
│   └── auth/
├── go.mod
├── go.sum
└── README.md
Contributing
Fork the project

Create your feature branch (git checkout -b feature/AmazingFeature)

Commit your changes (git commit -m 'Add some AmazingFeature')

Push to the branch (git push origin feature/AmazingFeature)

Open a Pull Request

License
This project is licensed under the MIT License.

Author
Rohit
