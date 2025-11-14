package tools

import (
	"context"
	"errors"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ModuloInput struct {
	A float64 `json:"a" jsonschema:"dividend"`
	B float64 `json:"b" jsonschema:"divisor"`
}

type ModuloOutput struct {
	Result float64 `json:"result" jsonschema:"the remainder of a divided by b"`
}

func Modulo(ctx context.Context, req *mcp.CallToolRequest, input ModuloInput) (
	*mcp.CallToolResult,
	ModuloOutput,
	error,
) {
	if input.B == 0 {
		return nil, ModuloOutput{}, errors.New("modulo by zero")
	}
	return nil, ModuloOutput{Result: math.Mod(input.A, input.B)}, nil
}
