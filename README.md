# BookingToGo Customer Registration API

A technical test project for BookingToGo, built with Go. This API manages customers, families, and nationalities, using PostgreSQL and Gorilla Mux.

## Features
- CRUD for Customers
- CRUD for Families (linked to Customers)
- CRUD for Nationalities
- RESTful endpoints
- Dockerized for easy deployment

## Project Structure
```
config.json
Dockerfile
go.mod
go.sum
Makefile
cmd/app/main.go
deployments/docker-compose.yml
internal/
  delivery/http/
    customer.go
    family.go
    nationality.go
    routes/route.go
  domain/domain.go
  repositories/
    customer.go
    family.go
    nationality.go
  usecases/
    customer.go
    family.go
    nationality.go
pkg/postgres/connection.go
```

## Getting Started

### Prerequisites
- Go 1.20+
- Docker
- PostgreSQL

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/DamiaRalitsa/customer-registration-go.git
   cd customer-registration-go
   ```
2. Configure your database connection in `config.json`.
3. Build and run with Docker:
   ```bash
   docker build -t bookingtogo-app .
   docker run -it --rm -p 8334:8334 -v $(pwd)/config.json:/config.json bookingtogo-app

   or simply run "make run" on terminal
   ```
4. Or run locally:
   ```bash
   go run cmd/app/main.go
   ```

### API Endpoints
#### Customers
- `GET /customers` - List all customers
- `GET /customers/{id}` - Get customer by ID
- `POST /customers` - Create customer
- `PUT /customers/{id}` - Update customer
- `DELETE /customers/{id}` - Delete customer

#### Families
- `GET /customers/{cst_id}/family` - List families for a customer
- `GET /family/{ft_id}` - Get family member by ID
- `POST /customers/{cst_id}/family` - Add family member
- `PUT /family/{ft_id}` - Update family member
- `DELETE /family/{ft_id}` - Delete family member

#### Nationalities
- `GET /nationalities` - List all nationalities
- `GET /nationalities/{id}` - Get nationality by ID
- `POST /nationalities` - Create nationality
- `PUT /nationalities/{id}` - Update nationality
- `DELETE /nationalities/{id}` - Delete nationality

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
MIT
