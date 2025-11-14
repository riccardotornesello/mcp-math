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
- Go 1.25.4 or higher

### Build from Source
```bash
git clone https://github.com/riccardotornesello/mcp-math.git
cd mcp-math
go build
```

## Usage

Start the MCP Math server:
```bash
./mcp-math
```

The server will start and listen for MCP protocol messages via standard input/output.

## Tool Reference

### Basic Arithmetic Operations

#### sum
Adds two numbers together.

**Input:**
```json
{
  "a": 5.0,
  "b": 3.0
}
```

**Output:**
```json
{
  "sum": 8.0
}
```

#### subtract
Subtracts the second number from the first number.

**Input:**
```json
{
  "a": 10.0,
  "b": 3.0
}
```

**Output:**
```json
{
  "result": 7.0
}
```

#### multiply
Multiplies two numbers.

**Input:**
```json
{
  "a": 4.0,
  "b": 5.0
}
```

**Output:**
```json
{
  "result": 20.0
}
```

#### divide
Divides the first number by the second number.

**Input:**
```json
{
  "a": 20.0,
  "b": 4.0
}
```

**Output:**
```json
{
  "result": 5.0
}
```

**Note:** Returns an error if dividing by zero.

#### modulo
Calculates the remainder when dividing the first number by the second.

**Input:**
```json
{
  "a": 17.0,
  "b": 5.0
}
```

**Output:**
```json
{
  "result": 2.0
}
```

**Note:** Returns an error if the divisor is zero.

### Array Operations

#### sum_array
Calculates the sum of all numbers in an array.

**Input:**
```json
{
  "numbers": [1.0, 2.0, 3.0, 4.0, 5.0]
}
```

**Output:**
```json
{
  "sum": 15.0
}
```

#### average
Calculates the average (mean) of all numbers in an array.

**Input:**
```json
{
  "numbers": [10.0, 20.0, 30.0, 40.0]
}
```

**Output:**
```json
{
  "average": 25.0
}
```

**Note:** Returns an error for empty arrays.

#### min
Finds the minimum value in an array of numbers.

**Input:**
```json
{
  "numbers": [5.0, 2.0, 8.0, 1.0, 9.0]
}
```

**Output:**
```json
{
  "min": 1.0
}
```

**Note:** Returns an error for empty arrays.

#### max
Finds the maximum value in an array of numbers.

**Input:**
```json
{
  "numbers": [5.0, 2.0, 8.0, 1.0, 9.0]
}
```

**Output:**
```json
{
  "max": 9.0
}
```

**Note:** Returns an error for empty arrays.

### Advanced Operations

#### power
Raises a base number to the power of an exponent.

**Input:**
```json
{
  "base": 2.0,
  "exponent": 3.0
}
```

**Output:**
```json
{
  "result": 8.0
}
```

#### root
Calculates the nth root of a number.

**Input:**
```json
{
  "number": 27.0,
  "n": 3.0
}
```

**Output:**
```json
{
  "result": 3.0
}
```

**Note:** Returns an error if n is zero.

### Trigonometric Functions

All trigonometric functions use radians for angle measurements.

#### sin
Calculates the sine of an angle.

**Input:**
```json
{
  "angle": 1.5707963267948966
}
```

**Output:**
```json
{
  "result": 1.0
}
```

**Note:** The example shows sin(π/2) = 1

#### cos
Calculates the cosine of an angle.

**Input:**
```json
{
  "angle": 0.0
}
```

**Output:**
```json
{
  "result": 1.0
}
```

**Note:** The example shows cos(0) = 1

#### tan
Calculates the tangent of an angle.

**Input:**
```json
{
  "angle": 0.7853981633974483
}
```

**Output:**
```json
{
  "result": 1.0
}
```

**Note:** The example shows tan(π/4) = 1

## Error Handling

The following operations will return errors in specific cases:

- **divide**: Returns an error when attempting to divide by zero
- **modulo**: Returns an error when the divisor is zero
- **root**: Returns an error when n (the root degree) is zero
- **average**, **min**, **max**: Return errors when given empty arrays

## Development

### Project Structure
```
mcp-math/
├── main.go           # Server entry point and tool registration
├── tools/            # Individual tool implementations
│   ├── sum.go
│   ├── subtract.go
│   ├── multiply.go
│   ├── divide.go
│   ├── modulo.go
│   ├── sum_array.go
│   ├── average.go
│   ├── min.go
│   ├── max.go
│   ├── power.go
│   ├── root.go
│   ├── sin.go
│   ├── cos.go
│   └── tan.go
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
