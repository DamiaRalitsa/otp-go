# OTP Service API

A technical test project for OTP (One-Time Password) service, built with Go. This API provides OTP generation and verification functionality using PostgreSQL and Gorilla Mux.

## Features
- OTP Generation
- OTP Verification
- RESTful endpoints
- PostgreSQL database integration
- Dockerized for easy deployment
- Unit tests with high coverage

## Project Structure
```
config.json
Dockerfile
go.mod
go.sum
Makefile
README.md
cmd/
  app/
    main.go
deployments/
  docker-compose.yml
internal/
  delivery/
    http/
      otp_controller.go
      routes/
        route.go
  domain/
    domain.go
    otp.go
  presenters/
    otp.go
    response.go
  repositories/
    otp/
      otp.go
  usecases/
    otp/
      otp.go
      otp_test.go
pkg/
  postgres/
    connection.go
```

## Getting Started

### Prerequisites
- Go 1.20+
- Docker
- PostgreSQL

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/DamiaRalitsa/otp-go.git
   cd otp-go
   ```
2. Configure your database connection in `config.json`.
3. Build and run with Docker:
   ```bash
   docker build -t otp-service .
   docker run -it --rm -p 8334:8334 -v $(pwd)/config.json:/config.json otp-service

   or simply run "make run" on terminal
   ```
4. Or run locally:
   ```bash
   go run cmd/app/main.go
   ```

### Testing
Run the unit tests:
```bash
make test
```

Generate test coverage report:
```bash
make test-coverage
```

### API Endpoints
#### OTP Service
- `POST /send-otp` - Generate and send OTP
  ```json
  {
    "user_id": "user123"
  }
  ```
  Response:
  ```json
  {
    "user_id": "user123",
    "otp": "123456"
  }
  ```

- `POST /verify-otp` - Verify OTP
  ```json
  {
    "user_id": "user123",
    "otp": "123456"
  }
  ```
  Response:
  ```json
  {
    "valid": true
  }
  ```

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
MIT
