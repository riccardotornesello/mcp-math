package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MultiplyInput struct {
	A float64 `json:"a" jsonschema:"first number to multiply"`
	B float64 `json:"b" jsonschema:"second number to multiply"`
}

type MultiplyOutput struct {
	Result float64 `json:"result" jsonschema:"the product of the two numbers"`
}

func Multiply(ctx context.Context, req *mcp.CallToolRequest, input MultiplyInput) (
	*mcp.CallToolResult,
	MultiplyOutput,
	error,
) {
	return nil, MultiplyOutput{Result: input.A * input.B}, nil
}
