# Migrations and API Documentation Setup

## Prerequisites

- Go installed
- [Migrate CLI](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) installed

---

## Migration Commands

### Create a new migration

```bash
migrate create -ext sql -dir ./cmd/migrate/migrations -seq create_tablename_table
```

### Run migration and push to database

```bash
go run ./cmd/migrate/main.go up
```

---

## Swagger (API Documentation)

### Install Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Generate documentation

```bash
swag init -g cmd/api/main.go -o ./docs
```

---

## Development

### Run project with hot reload

```bash
air
```