# Movies API

## Greeting

Welcome to the Movies API! This is a Go-based backend service that provides endpoints for managing movie reservations. The API is built with clean architecture principles and uses PostgreSQL for data persistence.

## Installation

### Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- PostgreSQL (or use the provided Docker setup)

### Setup Steps

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env
   ```

   Update the `.env` file with your database credentials and configuration.

4. **Start the database and build application (Docker)**

   ```bash
   docker-compose up -d --build
   ```

5. **Run database migrations**
   ```bash
   psql -U <db_user> -d <db_name> -f schema.sql
   ```

## Running the Application

### Using Docker Compose (Recommended)

```bash
docker-compose up -d --build
```

### Running Locally

```bash
go run cmd/main.go
```

The API server will start on `http://localhost:8080` (or the port specified in your `.env` file).

## API Documentation

Feel free to explore the following API endpoints:

[Postman Collection](https://www.postman.com/altimetry-architect-54647672/workspace/rodrigo-s-public-apis/collection/25405760-5b4968a6-320b-4ba9-96e2-589f16e7c632?action=share&creator=25405760)

## Project Structure

- **cmd/** - Application entry point
- **controller/** - HTTP request handlers
- **service/** - Business logic layer
- **usecase/** - Use case implementations
- **model/** - Data models and input structures
- **db/** - Database connection and queries
- **movie_reservation/** - Movie reservation specific logic
- **helpers/** - Utility functions

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Database

The project uses PostgreSQL. Schema is defined in [schema.sql](schema.sql) and queries are in [query.sql](query.sql).

SQL code generation is handled by [sqlc](https://sqlc.dev/) as configured in [sqlc.yaml](sqlc.yaml).

## License

This project is licensed under the MIT License.
