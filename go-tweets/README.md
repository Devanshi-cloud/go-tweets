# Go Tweets

Go Tweets is a simple Go application that provides functionality for creating and validating JSON Web Tokens (JWT). This project is designed to help developers manage user authentication and authorization using JWTs.

## Overview

The `pkg/jwt` package contains functions for generating and validating JWT tokens. The main functions include:

- `CreateToken`: Generates a JWT token with user ID and username as claims.
- `ValidateToken`: Validates a given token and extracts the user ID and username from it.

## Setup Instructions

To get started with the Go Tweets project, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Devanshi-cloud/go-tweets.git
   cd go-tweets
   ```

2. **Install dependencies:**
   Ensure you have Go installed on your machine. Run the following command to download the necessary dependencies:
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   You can run your Go application using:
   ```bash
   go run .
   ```

## Usage Examples

### Creating a Token

To create a JWT token, use the `CreateToken` function:

```go
token, err := jwt.CreateToken(userID, username, secretKey)
```

### Validating a Token

To validate a JWT token, use the `ValidateToken` function:

```go
userID, username, err := jwt.ValidateToken(tokenStr, secretKey, true)
```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.