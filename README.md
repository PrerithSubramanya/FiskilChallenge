# Go HTTP Sequence Service

This is a simple HTTP service written in Go that generates a unique sequence number for every request it receives.

## Features

- Listens on a configurable port and responds to HTTP GET requests.
- Returns a unique sequence number in the response body for each request.
- Handles multiple requests concurrently and is thread-safe.

## Usage

To run the service, simply use the `go run` command:

```bash
go run main.go
```

The service will start and listen on port 8080. You can test it by sending a GET request to `http://localhost:8080`. Each request will receive a unique sequence number in the response.

## Testing

To run the tests for the service, use the `go test` command:

```bash
go test
```

This will run a series of tests that check the functionality of the HTTP handler.
