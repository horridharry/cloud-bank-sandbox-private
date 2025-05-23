# Cloud Bank Sandbox

A containerized, cloud-native banking platform built in Go. This project simulates core banking services—like accounts and transactions—via isolated microservices running on Docker, with PostgreSQL persistence and real-world development practices.

## Features

- Modular microservices: Separate `accounts` and `transactions` services
- Fully containerized: Uses Docker Compose for orchestration
- PostgreSQL-backed: Data persistence across services
- REST APIs: JSON endpoints for service interaction
- Disk-efficient setup: Docker volume stored on external D: drive

## Stack

- Go 1.21
- PostgreSQL 15
- Docker + Docker Compose
- RESTful HTTP APIs
- UUIDs for entity tracking

## Project Structure

```
cloud-bank-sandbox/
├── docker-compose.yml
├── services/
│   ├── accounts/
│   │   ├── main.go
│   │   ├── handler.go
│   │   └── ...
│   └── transactions/
│       ├── main.go
│       ├── handler.go
│       └── ...
```

## Usage

Start all services:

```bash
docker compose up --build
```

Create an account:

```bash
curl -X POST http://localhost:8080/accounts \
  -H "Content-Type: application/json" \
  -d '{"name":"Harry"}'
```

Transfer money:

```bash
curl -X POST http://localhost:8081/transactions \
  -H "Content-Type: application/json" \
  -d '{"from_account_id":"<UUID>", "to_account_id":"<UUID>", "amount":100}'
```

Check all accounts:

```bash
curl http://localhost:8080/accounts
```

## Future Improvements

- API Gateway with request tracing
- Proper transaction rollback handling
- Service discovery (e.g. Consul)
- gRPC endpoints
- JWT authentication for account access

## License

MIT
