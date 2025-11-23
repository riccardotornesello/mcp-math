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
docker pull ghcr.io/riccardotornesello/mcp-math:latest
docker run -d -p 8080:8080 ghcr.io/riccardotornesello/mcp-math:latest
```

The server will be accessible at `http://localhost:8080/`

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

## Tool Reference

### Basic Arithmetic Operations

| Tool | Description | Parameters |
|------|-------------|------------|
| **sum** | Adds two numbers together | `a` (number), `b` (number) |
| **subtract** | Subtracts the second number from the first number | `a` (number), `b` (number) |
| **multiply** | Multiplies two numbers | `a` (number), `b` (number) |
| **divide** | Divides the first number by the second number | `a` (number), `b` (number) |
| **modulo** | Calculates the remainder when dividing the first number by the second | `a` (number), `b` (number) |

### Array Operations

| Tool | Description | Parameters |
|------|-------------|------------|
| **sum_array** | Calculates the sum of all numbers in an array | `numbers` (array of numbers) |
| **average** | Calculates the average (mean) of all numbers in an array | `numbers` (array of numbers) |
| **min** | Finds the minimum value in an array of numbers | `numbers` (array of numbers) |
| **max** | Finds the maximum value in an array of numbers | `numbers` (array of numbers) |

### Advanced Operations

| Tool | Description | Parameters |
|------|-------------|------------|
| **power** | Raises a base number to the power of an exponent | `base` (number), `exponent` (number) |
| **root** | Calculates the nth root of a number | `number` (number), `n` (number) |

### Trigonometric Functions

All trigonometric functions use radians for angle measurements.

| Tool | Description | Parameters |
|------|-------------|------------|
| **sin** | Calculates the sine of an angle | `angle` (number in radians) |
| **cos** | Calculates the cosine of an angle | `angle` (number in radians) |
| **tan** | Calculates the tangent of an angle | `angle` (number in radians) |

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
3. **Tags**: Images are tagged with version numbers (e.g., `v1.0.0`, `1.0`, `1`, `latest`)

To create a release:
1. Create and push a new tag: `git tag v1.0.0 && git push origin v1.0.0`
2. Create a release from the tag on GitHub
3. The CI/CD pipeline will automatically build and publish artifacts

## License

This project is licensed under the terms specified in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
