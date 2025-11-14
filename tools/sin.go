package tools

import (
	"context"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SinInput struct {
	Angle float64 `json:"angle" jsonschema:"angle in radians"`
}

type SinOutput struct {
	Result float64 `json:"result" jsonschema:"sine of the angle"`
}

func Sin(ctx context.Context, req *mcp.CallToolRequest, input SinInput) (
	*mcp.CallToolResult,
	SinOutput,
	error,
) {
	return nil, SinOutput{Result: math.Sin(input.Angle)}, nil
}
