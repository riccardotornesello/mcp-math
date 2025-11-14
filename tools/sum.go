package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SumInput struct {
	A float64 `json:"a" jsonschema:"first number to add"`
	B float64 `json:"b" jsonschema:"second number to add"`
}

type SumOutput struct {
	Sum float64 `json:"sum" jsonschema:"the sum of the two numbers"`
}

func Sum(ctx context.Context, req *mcp.CallToolRequest, input SumInput) (
	*mcp.CallToolResult,
	SumOutput,
	error,
) {
	return nil, SumOutput{Sum: input.A + input.B}, nil
}
