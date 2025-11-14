package main

import (
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"riccardotornesello.it/mcp-math/tools"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "math", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "sum", Description: "adds two numbers"}, tools.Sum)

	handler := mcp.NewStreamableHTTPHandler(
		func(*http.Request) *mcp.Server { return server },
		&mcp.StreamableHTTPOptions{},
	)
	http.HandleFunc("/mcp", handler.ServeHTTP)

	log.Println("Starting MCP math server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
