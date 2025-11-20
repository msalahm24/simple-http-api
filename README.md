# simple-http-api

A simple Go HTTP API with alphabet-based name validation.

## Features

- Validates names starting with letters A-M

## Quick Start

```bash
# Run the application
make run

# Run tests
make test

# Build
make build
```

## API Endpoints

### GET /hello-world

Greets names that start with A-M.

**Example:**
```bash
curl "http://localhost:8080/hello-world?name=Alice"
# {"message":"Hello Alice"}

curl "http://localhost:8080/hello-world?name=Nancy"
# {"error":"name must start with a letter between A and M"}
```

### GET /health

Health check endpoint.

```bash
curl "http://localhost:8080/health"
# {"status":"ok"}
```

## Configuration

Set the `PORT` environment variable to change the default port (8080):
```bash
PORT=3000 make run
```
