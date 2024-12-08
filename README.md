# Go Echo Boilerplate

This boilerplate is designed to provide a quick start for production ready web application development using the Echo framework in Go. It includes implementations for CRUD operations, JWT authentication, user roles, and more.

<p align="center">
    <img src="https://img.shields.io/badge/golang-v1.23-lightblue" height="25"/>
    <img src="https://img.shields.io/badge/echo-v4.13-blue" height="25"/>
    <img src="https://img.shields.io/badge/gorm-v1.25-green" height="25"/>
    <img src="https://img.shields.io/badge/echo--jwt-v1.25-blue" height="25"/>
    <img src="https://img.shields.io/badge/conf-viper-%66bc67" height="25"/>
    <img src="https://img.shields.io/badge/db-postgres-%23336791" height="25"/>
</p>

## Table of Contents

- [Features](#features)
  - [CRUD Operations](#crud-operations)
  - [Router-Controller-Model Structure](#router-controller-model-structure)
  - [Authorization using Login Password](#authorization-using-login-password)
  - [JWT Authentication using Middleware](#jwt-authentication-using-middleware)
  - [Roles for Users](#roles-for-users)
  - [JSON Response Wrapper](#json-response-wrapper)
  - [Environment Variables](#environment-variables)
  - [Postman Collection with Environment Setup](#postman-collection-with-environment-setup)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)

## Features

### CRUD Operations

This boilerplate supports Create, Read, Update, and Delete operations for various models. The CRUD functionality is implemented using a clean architecture approach with separate layers for routing, controllers, and models.

### Router-Controller-Model Structure

- **Router**: Defines all the application routes and middleware.
- **Controller**: Handles the business logic and communicates between the router and models.
- **Model**: Represents the data structure and provides functions to interact with the database.

### Authorization using Login Password

Users can authenticate themselves using a login endpoint that requires a username and password. Passwords are securely hashed using bcrypt before being stored in the database.

### JWT Authentication using Middleware

JWT (JSON Web Token) is used to authenticate API requests. Middleware is employed to protect routes and ensure that only authenticated users can access certain endpoints.

```json
{
  "name": "test",
  "id": 2,
  "role": "USER",
  "exp": 1716461438
}
```

### Roles for Users

User roles (e.g., Admin, User) are implemented to manage permissions and access control within the application. Middleware checks user roles to authorize access to specific routes.

```golang
const (
	Admin     	UserRole = "ADMIN"
	Moderator 	UserRole = "MODERATOR"
	User 		UserRole = "USER"
)
```

### JSON Response Wrapper

All API responses are wrapped in a standard JSON format, including status codes, messages, and data payloads, ensuring consistency and ease of use for frontend applications.

```json
{
  "Success": true,
  "Message": "Success",
  "Data": [
    {
      "blogs": [
        {
          "ID": 1,
          "CreatedAt": "2024-05-22T22:41:11.616688Z",
          "UpdatedAt": "2024-05-22T22:41:11.616688Z",
          "DeletedAt": null,
          "UserId": 2,
          "Title": "test1",
          "Content": "content1"
        },
        {
          "ID": 2,
          "CreatedAt": "2024-05-23T12:50:46.833646Z",
          "UpdatedAt": "2024-05-23T12:50:46.833646Z",
          "DeletedAt": null,
          "UserId": 2,
          "Title": "test2",
          "Content": "content2"
        }
      ]
    }
  ]
}
```

### Environment Variables

Environment variables are used to configure the application, such as database connection strings, JWT secret keys, and other settings. The `config.json` file is utilized to manage these variables.

Environment variables

```json
{
  "Db": {
    "Host": "localhost",
    "Port": "5432",
    "Name": "test",
    "User": "postgres",
    "Password": "123",
    "SslMode": "disable"
  },
  "Jwt": {
    "SecretKey": "secret_key",
    "ExpTime": 3
  }
}
```

implement like this using `viper` in `main.go`

```golang
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//... somewherre in your code
// example loading database config envs
db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
    viper.GetString("Db.Host"),
    viper.GetString("Db.Port"),
    viper.GetString("Db.Name"),
    viper.GetString("Db.User"),
    viper.GetString("Db.Password"),
    viper.GetString("Db.SslMode"),
  )}), &gorm.Config{})
```

### Postman Collection with Environment Setup

A Postman collection (`_postman` folder) is provided to facilitate testing of the API endpoints. It includes pre-configured requests for all routes and an environment setup to manage variables like base URL and authentication tokens.

## Getting Started

### Prerequisites

- Go (version 1.23+)
- Postman (optional, for testing API endpoints)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/altynboy/go-echo-boilerplate.git
   cd go-echo-boilerplate
   ```
2. **Install dependecies**
   ```bash
   go mod tidy
   ```
3. **Setup env variables** in `config.json`
   ```json
   {
     "Db": {
       "Host": "localhost",
       "Port": "5432",
       "Name": "test",
       "User": "postgres",
       "Password": "123",
       "SslMode": "disable"
     },
     "Jwt": {
       "SecretKey": "secret_key",
       "ExpTime": 3
     }
   }
   ```
4. **Run the application**
   ```bash
   go run main.go
   ```
