package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SubtractInput struct {
	A float64 `json:"a" jsonschema:"first number"`
	B float64 `json:"b" jsonschema:"second number to subtract from the first"`
}

type SubtractOutput struct {
	Result float64 `json:"result" jsonschema:"the result of a - b"`
}

func Subtract(ctx context.Context, req *mcp.CallToolRequest, input SubtractInput) (
	*mcp.CallToolResult,
	SubtractOutput,
	error,
) {
	return nil, SubtractOutput{Result: input.A - input.B}, nil
}
