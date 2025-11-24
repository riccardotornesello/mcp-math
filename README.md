# MCP Math Server

A Model Context Protocol (MCP) server that provides a comprehensive set of mathematical operations and functions.

## Features

This MCP server provides 14 mathematical tools organized into the following categories:

### Basic Arithmetic Operations

- **sum** - Add two numbers
- **subtract** - Subtract two numbers
- **multiply** - Multiply two numbers
- **divide** - Divide two numbers
- **modulo** - Calculate the remainder of division

### Array Operations

- **sum_array** - Calculate the sum of all numbers in an array
- **average** - Calculate the average of all numbers in an array
- **min** - Find the minimum value in an array
- **max** - Find the maximum value in an array

### Advanced Operations

- **power** - Raise a number to a power
- **root** - Calculate the nth root of a number

### Trigonometric Functions

- **sin** - Calculate sine of an angle (in radians)
- **cos** - Calculate cosine of an angle (in radians)
- **tan** - Calculate tangent of an angle (in radians)

## Installation

### Docker (Recommended)

Pull and run the latest image:

```bash
docker pull riccardotornesello/mcp-math:latest
docker run -d -p 8080:8080 riccardotornesello/mcp-math:latest
```

The server will be accessible at `http://localhost:8080/mcp`

### From Binary

Download the latest binary from the [releases page](https://github.com/riccardotornesello/mcp-math/releases), then run:

```bash
chmod +x mcp-math
./mcp-math
```

### Build from Source

```bash
git clone https://github.com/riccardotornesello/mcp-math.git
cd mcp-math
go build
./mcp-math
```

## Development

### Project Structure

```
mcp-math/
├── main.go           # Server entry point and tool registration
├── tools/            # Individual tool implementations
├── Dockerfile        # Docker image configuration
├── .github/
│   └── workflows/    # CI/CD pipelines
└── README.md
```

### Building and Testing

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...

# Build
go build

# Test Docker build
docker build -t mcp-math .
```

## Release Process

Releases are automated via GitHub Actions. When a new release is published:

1. **Binary Build**: The Go binary is built and attached to the release
2. **Docker Images**: Multi-platform images (amd64, arm64) are built and pushed to:
   - GitHub Container Registry: `ghcr.io/riccardotornesello/mcp-math`
   - Docker Hub: `riccardotornesello/mcp-math`
3. **Tags**: Images are tagged with version numbers (e.g., `v1.0.0`, `latest`)

## License

This project is licensed under the terms specified in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
