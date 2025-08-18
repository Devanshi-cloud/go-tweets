# Go Tweets

Go Tweets is a simple Twitter-like application built using Go. This project demonstrates the use of JWT for authentication, along with a structured approach to organizing Go code.

## Project Structure

The project is organized as follows:

```
go-tweets
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go        # Handler functions for processing requests
│   ├── middlewares
│   │   └── middlewares.go      # Middleware functions for request processing
│   ├── models
│   │   └── models.go          # Data models for tweets and users
│   └── routes
│       └── routes.go          # Application routes setup
├── pkg
│   └── jwt
│       └── jwt.go             # Functions for creating and validating JWT tokens
├── .gitignore                  # Files and directories to be ignored by Git
├── go.mod                      # Module definition and dependencies
├── go.sum                      # Checksums for module dependencies
└── README.md                   # Documentation for the project
```

## Features

- JWT-based authentication for secure access.
- Modular design with clear separation of concerns.
- Simple and intuitive API for managing tweets.

## Getting Started

To get started with the Go Tweets application, follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/Devanshi-cloud/go-tweets.git
   cd go-tweets
   ```

2. Install the dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage

Once the application is running, you can interact with the API endpoints to create and manage tweets. Refer to the API documentation for details on available endpoints and their usage.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.