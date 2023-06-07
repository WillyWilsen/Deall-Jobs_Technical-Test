# Deall-Jobs_Technical-Test

## Structure

├── README.md

├── config.json

├── database

│   └── database.go

├── db.sql

├── go.mod

├── go.sum

├── handler

│   └── auth.go

├── model

│   └── user.go

├── repository

│   └── auth.go

├── server.go

└── utility

│   └── utility.go

- `config.json`: A configuration file that holds various configuration settings for the service, such as HTTP port and database connection details.
- `database`: A directory that contains the database-related code and files.
    - `database.go`: The database package that provides functionality for connecting to the database.
- `db.sql`: A SQL file that contains the database schema to set up the initial database structure or tables.
- `handler`: A directory that holds the HTTP request handlers.
    - `auth.go`: The handler package that contains the HTTP handlers for authentication-related endpoints, such as registration and login.
- `model`: A directory that contains the data models or structures used in the service.
    - `user.go`: The model package that defines the structure of the User model, including its properties and validation rules.
- `repository`: A directory that contains the repository layer responsible for data access and persistence.
    - `auth.go`: The repository package that provides functions to interact with the database and perform operations related to authentication, such as user creation and retrieval.
- `server.go`: The main entry point of the service. It initializes the server, sets up the necessary routes, and starts the server to listen for incoming HTTP requests.
- `utility`: A directory that holds utility or helper functions used throughout the service.
    - `utility.go`: The utility package that contains common functions or utilities that can be reused across different parts of the service.

## Getting Started

### Prerequisites

Before running the service, make sure you have the following prerequisites installed:

- Go (version 1.19+)
- MySQL

### Installation

1. Clone the repository
```
git clone https://github.com/WillyWilsen/Deall-Jobs_Technical-Test.git
```
2. Install the required dependencies
```
go mod download
```
3. Update the `config.json` file with your HTTP port and database connection details
4. Run the `db.sql` to your MySQL database to set up the initial database tables

### Usage

To start the service, run the following command
```
go run server.go
```

You can access it via `http://localhost:{http_port}`. Replace the `http_port` with your HTTP port in `config.json`

## API Endpoints

- `/api/register` (POST): Register

    payload: {
        "name": string,
        "email": string,
        "password": string
    }

- `/api/login` (POST): Login

    payload: {
        "email": string,
        "password": string
    }