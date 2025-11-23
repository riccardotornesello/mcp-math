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

### Prerequisites
- Go 1.25.4 or higher (for building from source)
- Docker (for running with Docker)

### Using Docker (Recommended)

Pull the image from GitHub Container Registry:
```bash
docker pull ghcr.io/riccardotornesello/mcp-math:latest
```

Or from Docker Hub:
```bash
docker pull riccardotornesello/mcp-math:latest
```

Run the container:
```bash
docker run -d -p 8080:8080 ghcr.io/riccardotornesello/mcp-math:latest
```

### Build from Source
```bash
git clone https://github.com/riccardotornesello/mcp-math.git
cd mcp-math
go build
```

### Build Docker Image Locally
```bash
git clone https://github.com/riccardotornesello/mcp-math.git
cd mcp-math
docker build -t mcp-math .
docker run -d -p 8080:8080 mcp-math
```

## Usage

### Running with Docker
```bash
# Pull and run the latest image
docker run -d -p 8080:8080 ghcr.io/riccardotornesello/mcp-math:latest

# The server will start and listen on port 8080
# Access it at http://localhost:8080/
```

### Running from Binary
Start the MCP Math server:
```bash
./mcp-math
```

The server will start and listen for MCP protocol messages via HTTP on port 8080.

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
│   ├── ...
│   └── ...
├── go.mod
└── README.md
```

### Building
```bash
go build
```

### Running
```bash
./mcp-math
```

## License

This project is licensed under the terms specified in the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Docker Images

Docker images are automatically built and published to both GitHub Container Registry and Docker Hub on every release:

- **GitHub Container Registry**: `ghcr.io/riccardotornesello/mcp-math`
- **Docker Hub**: `riccardotornesello/mcp-math`

Images are tagged with:
- `latest` - Latest stable release
- `vX.Y.Z` - Specific version (e.g., `v1.0.0`)
