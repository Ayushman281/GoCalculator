# Go Calculator API

A simple REST API calculator built with Go that performs basic arithmetic operations (addition, subtraction, multiplication, and division) on both integers and floating-point numbers.

## Features

- ✅ Addition, subtraction, multiplication, and division operations
- ✅ Support for both integer and float operations
- ✅ Structured logging with request/response tracking
- ✅ Error handling with meaningful error messages
- ✅ Health check endpoint
- ✅ Division by zero protection
- ✅ Custom middleware for logging requests and responses

## Project Structure

```
GoCalculatorAPI/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/                # HTTP request handlers
│   │   ├── add.go
│   │   ├── subtract.go
│   │   ├── multiply.go
│   │   ├── divide.go
│   │   └── health.go
│   ├── middleware/              # HTTP middleware
│   │   ├── logging.go
│   │   └── response_writer.go
│   ├── models/                  # Request/response models
│   │   └── requests.go
│   ├── respond/                 # Response utilities
│   │   └── response.go
│   └── logger/                  # Logger configuration
│       └── logger.go
```

## Prerequisites

- Go 1.21 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Ayushman281/calculator-api.git
cd calculator-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
```
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

### Addition

#### Integer Addition
```
POST /add/int
```

**Request Body:**
```json
{
  "a": 5,
  "b": 3
}
```

**Response:**
```json
{
  "result": 8
}
```

#### Float Addition
```
POST /add/float
```

**Request Body:**
```json
{
  "a": 5.5,
  "b": 3.2
}
```

**Response:**
```json
{
  "result": 8.7
}
```

### Subtraction

#### Integer Subtraction
```
POST /sub/int
```

#### Float Subtraction
```
POST /sub/float
```

### Multiplication

#### Integer Multiplication
```
POST /multiply/int
```

#### Float Multiplication
```
POST /multiply/float
```

### Division

#### Integer Division
```
POST /divide/int
```

#### Float Division
```
POST /divide/float
```

**Note:** Division by zero returns a `400 Bad Request` error.

## Example Usage

Using `curl`:

```bash
# Integer addition
curl -X POST http://localhost:8080/add/int \
  -H "Content-Type: application/json" \
  -d '{"a": 10, "b": 5}'

# Float division
curl -X POST http://localhost:8080/divide/float \
  -H "Content-Type: application/json" \
  -d '{"a": 10.5, "b": 2.5}'

# Health check
curl http://localhost:8080/health
```

## Error Handling

The API returns structured error messages in JSON format:

```json
{
  "error": "error message description"
}
```

Common HTTP status codes:
- `200 OK` - Successful operation
- `400 Bad Request` - Invalid input or division by zero
- `405 Method Not Allowed` - Wrong HTTP method used
- `500 Internal Server Error` - Server error

## Logging

The application uses structured logging with the following information:
- Request method, path, and remote address
- Response status code
- Response size in bytes
- Request duration in milliseconds

## License

This project is open source and available under the MIT License.

## Author

Ayushman281
