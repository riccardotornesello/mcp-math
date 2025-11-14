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

	mcp.AddTool(server, &mcp.Tool{Name: "sum", Description: "adds two numbers"}, tools.Sum)

	fmt.Println("Starting MCP Math server...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server stopped.")
}
