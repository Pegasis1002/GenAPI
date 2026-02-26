# GenAPI

GenAPI is a Go-based API boilerplate designed for scalability and modularity. It features structured logging, a Gin-based router, and a modular architecture to facilitate easy extension.

> **Note:** This project is currently in the early stages of development. Much of the codebase consists of placeholders, with core infrastructure like the logger and basic routing already implemented.

## Features

- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin) for high-performance HTTP routing.
- **Logging**: Structured JSON logging using Go's standard `slog` library.
- **Database**: Integration-ready with `pgx` (PostgreSQL).
- **Architecture**: Clean, modular structure separating core logic from domain-specific modules.
- **Containerization**: Docker and Docker Compose support.

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) 1.25.0 or higher.
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized deployment).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Pegasis1002/GenAPI.git
   cd genapi
   ```

2. Download dependencies:
   ```bash
   go mod download
   ```

### Running the API

#### Locally

To run the API locally, use the following command:

```bash
DB_URL="postgres://user:password@localhost:5432/dbname" go run cmd/api/main.go
```

The server will start on `http://0.0.0.0:8080`.

#### Using Docker

Build and run the application using Docker Compose:

```bash
docker-compose up --build
```

## API Documentation

### Endpoints (Current Placeholders)

- `GET /users`: Returns a list of placeholder users.

### Logging

The API uses structured JSON logging. Each request is logged with its status, method, path, latency, and client IP.

Example log output:
```json
{"time":"2026-02-26T10:00:00Z","level":"INFO","msg":"request_completed","status":200,"method":"GET","path":"/users","latency":500000,"ip":"127.0.0.1"}
```

## Extending the API

The project is structured to be easily extendable:

### Directory Structure

- `cmd/api/`: Application entry point.
- `configs/`: Configuration files (e.g., `config.yml`).
- `internal/core/`: Essential infrastructure components (DB, Logger, Router).
- `internal/middleware/`: Custom Gin middlewares.
- `internal/modules/`: Domain-specific business logic. This is where you should add new features.
- `pkg/`: Publicly exportable utility packages.

### Adding a New Module

1. Create a new directory under `internal/modules/` (e.g., `internal/modules/my_feature`).
2. Implement your handlers, services, and models within that directory.
3. Register your routes in `internal/core/router/gin.go` (or create a registration mechanism in your module and call it from the core router).

### Modifying the Logger

The logger is initialized in `internal/core/logger/logger.go`. You can customize the output format or level there.
