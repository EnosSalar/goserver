# GoServer

A simple HTTP web server written in Go that demonstrates basic web server functionality.

## Overview

GoServer is a minimal HTTP server implementation using Go's standard library. It serves as a starting point for learning Go web development or as a foundation for building more complex web applications.

## Features

- Lightweight HTTP server
- Simple "Hello World" endpoint
- No external dependencies
- Easy to understand and extend

## Prerequisites

- Go 1.23 or higher
- Basic understanding of Go programming

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd goserver
```

2. Download dependencies (if any):
```bash
go mod download
```

## Usage

### Running the Server

Start the server with:

```bash
go run main.go
```

The server will start on port 8082 and display:
```
Server is running on http://localhost:8082
Visit http://localhost:8082/hello
```

### Building the Server

To build a binary executable:

```bash
go build -o goserver
```

Then run the compiled binary:

```bash
./goserver
```

## API Endpoints

### GET /hello

Returns a simple "Hello World" message.

**Request:**
```bash
curl http://localhost:8082/hello
```

**Response:**
```
Hello World
```

**Status Code:** 200 OK

## Project Structure

```
goserver/
├── main.go          # Main application entry point
├── go.mod           # Go module definition
└── README.md        # Project documentation
```

## Code Overview

### main.go

The main application file contains:

- **`helloHandler()`** - HTTP handler function that responds with "Hello World"
- **`main()`** - Entry point that:
  - Registers the `/hello` route
  - Starts the HTTP server on port 8082
  - Logs server activity

## Configuration

The server is currently configured to run on port `8082`. To change the port, modify line 19 in `main.go`:

```go
log.Fatal(http.ListenAndServe(":8082", nil))
```

Replace `8082` with your desired port number.

## Development

### Adding New Endpoints

To add a new endpoint:

1. Create a handler function:
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "My Response")
}
```

2. Register it in the `main()` function:
```go
http.HandleFunc("/myroute", myHandler)
```

### Testing

Test the server manually:

```bash
# Start the server
go run main.go

# In another terminal, test the endpoint
curl http://localhost:8082/hello
```

## Dependencies

This project uses only the Go standard library:

- `fmt` - Formatted I/O operations
- `log` - Logging functionality
- `net/http` - HTTP server implementation

## Troubleshooting

### Port Already in Use

If you see an error like `bind: address already in use`, either:
- Stop the process using port 8082
- Change the port in `main.go`

### Go Version Issues

Ensure you have Go 1.23 or higher installed:

```bash
go version
```

## Future Enhancements

Possible improvements for this project:

- [ ] Add more endpoints (e.g., JSON API)
- [ ] Implement request logging middleware
- [ ] Add environment-based configuration
- [ ] Support for different HTTP methods (POST, PUT, DELETE)
- [ ] Add unit tests
- [ ] Implement graceful shutdown
- [ ] Add Docker support

## License

[Specify your license here]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

[Your name/organization]

## Acknowledgments

Built with Go's standard `net/http` package.
