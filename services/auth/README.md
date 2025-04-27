# Template Service Developer Guidelines

## Overview

The Template Service is a microservice template that provides a foundation for building new services. It includes standard health check endpoints, version information, and a basic configuration system.

## Features

- **Health Check Endpoints**: Liveness and readiness probes for Kubernetes compatibility
- **Version Information**: Endpoint to retrieve build and version details
- **Configuration**: Environment variable-based configuration
- **Graceful Shutdown**: Proper signal handling for clean shutdowns
- **Docker Support**: Ready-to-use Docker build and run commands
- **CI/CD Integration**: GoReleaser configuration for automated builds and releases

## Project Structure

```
service/
├── bin/                  # Binary output directory
├── cmd/
│   └── service/          # Main application entry point
├── internal/
│   ├── app/
│   │   └── service/      # Service implementation
│   └── config/           # Configuration
│       └── build/        # Build information
├── .goreleaser.yaml      # GoReleaser configuration
├── Makefile              # Main Makefile
├── mk_build.mk           # Build-related make targets
├── mk_docker.mk          # Docker-related make targets
├── mk_lint.mk            # Linting-related make targets
└── mk_test.mk            # Testing-related make targets
```

## Getting Started

### Prerequisites

- Go 1.23 or later
- Docker (for containerized builds and runs)
- Make

### Building the Service

To build the service:

```bash
make build
```

This will create a binary in the `bin/` directory.

### Running the Service

To run the service locally:

```bash
make run
```

The service will start on port 8080 by default.

### Running with Docker

To build and run the service using Docker:

```bash
make build-image
make docker-run
```

To stop the Docker container:

```bash
make docker-stop
```

## Configuration

The service is configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| HOST     | Host to bind the server to | 0.0.0.0 |
| PORT     | Port to listen on | 8080 |

## API Endpoints

The service exposes the following endpoints:

- `GET /v1/liveness`: Health check endpoint for liveness probes
- `GET /v1/readiness`: Health check endpoint for readiness probes
- `GET /v1/version`: Returns version information about the service

## Development Workflow

### Code Style and Linting

To run the linter:

```bash
make lint
```

To format the code:

```bash
make fmt
```

### Running Tests

To run tests:

```bash
make test
```

To run tests with coverage:

```bash
make ci-test-with-coverage
```

### Extending Configuration

To add new configuration options:

1. Add the field to the appropriate struct in `internal/config/config.go`
2. Use the `env` tag to specify the environment variable name and default value

Example:

```
// In internal/config/config.go
type App struct {
    Host string `env:"HOST" envDefault:"0.0.0.0"`
    Port string `env:"PORT" envDefault:"8080"`
    LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}
```

## Building for Release

The service uses GoReleaser for building release artifacts. The configuration is in `.goreleaser.yaml`.

To create a release build:

```bash
make build
```

## Best Practices

1. **Error Handling**: Always check and handle errors appropriately
2. **Logging**: Use structured logging with appropriate log levels
3. **Configuration**: Make service configurable via environment variables
4. **Testing**: Write unit tests for all new functionality
5. **Documentation**: Update this README when adding new features or changing existing ones

## Troubleshooting

### Common Issues

- **Port already in use**: Change the port using the PORT environment variable
- **Permission denied**: Ensure you have the necessary permissions to bind to the specified port

## Contributing

1. Create a feature branch from main
2. Make your changes
3. Run tests and linting
4. Submit a pull request

## License

[Specify the license here]
