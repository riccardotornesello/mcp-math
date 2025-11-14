package main

import (
	"context"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"riccardotornesello.it/mcp-math/tools"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "math", Version: "v1.0.0"}, nil)

	// Basic arithmetic operations
	mcp.AddTool(server, &mcp.Tool{Name: "sum", Description: "adds two numbers"}, tools.Sum)
	mcp.AddTool(server, &mcp.Tool{Name: "subtract", Description: "subtracts second number from first number"}, tools.Subtract)
	mcp.AddTool(server, &mcp.Tool{Name: "multiply", Description: "multiplies two numbers"}, tools.Multiply)
	mcp.AddTool(server, &mcp.Tool{Name: "divide", Description: "divides first number by second number"}, tools.Divide)
	mcp.AddTool(server, &mcp.Tool{Name: "modulo", Description: "calculates remainder of division"}, tools.Modulo)

	// Array operations
	mcp.AddTool(server, &mcp.Tool{Name: "sum_array", Description: "calculates sum of all numbers in an array"}, tools.SumArray)
	mcp.AddTool(server, &mcp.Tool{Name: "average", Description: "calculates average of all numbers in an array"}, tools.Average)
	mcp.AddTool(server, &mcp.Tool{Name: "min", Description: "finds minimum value in an array"}, tools.Min)
	mcp.AddTool(server, &mcp.Tool{Name: "max", Description: "finds maximum value in an array"}, tools.Max)

	// Advanced operations
	mcp.AddTool(server, &mcp.Tool{Name: "power", Description: "raises base to the power of exponent"}, tools.Power)
	mcp.AddTool(server, &mcp.Tool{Name: "root", Description: "calculates nth root of a number"}, tools.Root)

	// Trigonometric functions
	mcp.AddTool(server, &mcp.Tool{Name: "sin", Description: "calculates sine of an angle in radians"}, tools.Sin)
	mcp.AddTool(server, &mcp.Tool{Name: "cos", Description: "calculates cosine of an angle in radians"}, tools.Cos)
	mcp.AddTool(server, &mcp.Tool{Name: "tan", Description: "calculates tangent of an angle in radians"}, tools.Tan)

	fmt.Println("Starting MCP Math server...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server stopped.")
}
