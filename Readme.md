# Bank API

This is a simple bank API built with Go, using the Gorilla Mux router and MySQL database. The project demonstrates a basic structure for a RESTful API with account and transaction management functionalities.

## Project Structure

├── controller
│ ├── account_controller.go
│ ├── transaction_controller.go
├── models
│ ├── account.go
│ ├── transaction.go
│ ├── details.go
├── errors
│ ├── errors.go
├── http
│ ├── http.go
├── repository
│ ├── account_repository.go
│ ├── transaction_repository.go
├── routes
│ ├── routes.go
├── service
│ ├── account_service.go
│ ├── transaction_service.go
├── main.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── README.md

bash
Copy code

## Prerequisites

- Go 1.16 or later
- Docker
- Docker Compose

## Setup and Running the Application

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/bank-api.git
cd bank-api
2. Build and Run with Docker
Make sure Docker and Docker Compose are installed on your machine. You can build and run the application using Docker Compose:

bash
Copy code
docker-compose up --build
3. Makefile Commands
A Makefile is included to simplify common tasks:

Build the application:

bash
Copy code
make build
Run the application:

bash
Copy code
make run



4. Accessing the Application
The application will be available at http://localhost:8082.

Swagger UI will be available at http://localhost:8082/swagger/index.html.

API Endpoints
Accounts
Create Account

URL: /accounts

Method: POST

Request Body:

json
Copy code
{
  "document_number": "12345678900"
}
Response:

json
Copy code
{
  "account_id": 1,
  "document_number": "12345678900"
}
Get Account

URL: /accounts/{id}

Method: GET

Response:

json
Copy code
{
  "account_id": 1,
  "document_number": "12345678900"
}
Transactions
Create Transaction

URL: /transactions

Method: POST

Request Body:

json
Copy code
{
  "account_id": 1,
  "operation_type_id": 4,
  "amount": 123.45
}
Response:

json
Copy code
{
  "transaction_id": 1,
  "account_id": 1,
  "operation_type_id": 4,
  "amount": 123.45,
  "event_date": "2023-06-23T10:00:00Z"
}
Project Details
Controller Layer
Account Controller: Handles account-related HTTP requests.
Transaction Controller: Handles transaction-related HTTP requests.
Service Layer
Account Service: Contains business logic for accounts.
Transaction Service: Contains business logic for transactions.
Repository Layer
Account Repository: Contains database operations for accounts.
Transaction Repository: Contains database operations for transactions.
Models Layer
Account: Represents account data.
Transaction: Represents transaction data.
Details: Represents detailed information about accounts.
Error Handling
Errors: Custom error handling and response formatting.
Routing
Routes: Defines all API routes and their corresponding handlers.
Database
MySQL: The application uses MySQL as the database. The schema is created and managed through raw SQL queries.