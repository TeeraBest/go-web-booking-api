# Go Web API

A modern, production-ready RESTful API built with Go, Gin, and Docker. Includes user authentication, booking management, Prometheus metrics, Swagger documentation, and more.

## Features
- User registration and login (JWT-based authentication)
- Booking management endpoints
- Prometheus metrics endpoint
- Swagger UI for API documentation
- Custom validation (mobile, password)
- Dockerized for easy deployment

## Prerequisites
- [Go](https://golang.org/dl/) 1.18+
- [Docker](https://www.docker.com/get-started)
- [Make](https://www.gnu.org/software/make/) (optional, for convenience)

## Getting Started

### 1. Clone the repository
```sh
git clone https://github.com/TeeraBest/go-web-booking-api.git
cd go-web-api
```

### 2. Initialize and run with Docker
```sh
docker compose -p go-booking-api -f "docker/docker-compose.yml" up -d postgres pgadmin redis
```
This will build and start the API server and Redis (as defined in `docker-compose.yml`).

### 3. Install swagger and run app
```sh
cd src
go install github.com/swaggo/swag/cmd/swag@latest
cd src/cmd
go run main.go
```

## API Documentation
- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Prometheus metrics: [http://localhost:8080/metrics](http://localhost:8080/metrics)

## Example API Usage

### Login (by username)
```sh
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username": "your_username", "password": "your_password"}'
```

### Booking (requires JWT token)
```sh
curl -X POST http://localhost:8080/api/v1/bookings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{"event_id": 123, "seat": "A1"}'
```

## Environment Configuration
- Edit `src/config/config-docker.yml` or `src/config/config-development.yml` for environment-specific settings.
- Redis password and other secrets should be managed securely (do not commit real secrets to public repos).

---
**Note:** Replace `<your_jwt_token>` with your actual values. For more API endpoints, see the Swagger UI.
