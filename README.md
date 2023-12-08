# GoJob Opportunities API

This project is a modern job opportunities API built using Golang. The API is powered by Go-Gin as a router, GoORM for database communication, SQLite as the database, Go Playground Validator and Swagger for API documentation and testing. The project follows a modern package structure to keep the codebase organized and maintainable.

---

## Features

- Introduction to Golang and building modern APIs
- Development environment setup for creating the API
- Using Go-Gin as a router for route management
- Validating data with Go Playground Validator
- Implementing SQLite as the database for the API
- Using GoORM for communication with the database
- Integrating Swagger for API documentation and testing
- Creating a modern package structure for organizing the project
- Implementing a complete job opportunities API with endpoints for searching, creating, editing, and deleting opportunities
- Implementing automated tests to ensure API quality

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/JulioZittei/goportunities.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `gopportunities`.
- `make test`: Run tests for all packages in the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `gopportunities` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [GoPlayground Validator](https://github.com/go-playground/validator) for validation
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

Made with ♥ by Julio Zittei 👋🏻
