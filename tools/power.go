package tools

import (
	"context"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type PowerInput struct {
	Base     float64 `json:"base" jsonschema:"the base number"`
	Exponent float64 `json:"exponent" jsonschema:"the exponent"`
}

type PowerOutput struct {
	Result float64 `json:"result" jsonschema:"base raised to the power of exponent"`
}

func Power(ctx context.Context, req *mcp.CallToolRequest, input PowerInput) (
	*mcp.CallToolResult,
	PowerOutput,
	error,
) {
	return nil, PowerOutput{Result: math.Pow(input.Base, input.Exponent)}, nil
}
