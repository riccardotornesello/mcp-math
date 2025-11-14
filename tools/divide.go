package tools

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type DivideInput struct {
	A float64 `json:"a" jsonschema:"dividend (number to be divided)"`
	B float64 `json:"b" jsonschema:"divisor (number to divide by)"`
}

type DivideOutput struct {
	Result float64 `json:"result" jsonschema:"the result of a / b"`
}

func Divide(ctx context.Context, req *mcp.CallToolRequest, input DivideInput) (
	*mcp.CallToolResult,
	DivideOutput,
	error,
) {
	if input.B == 0 {
		return nil, DivideOutput{}, errors.New("division by zero")
	}
	return nil, DivideOutput{Result: input.A / input.B}, nil
}
