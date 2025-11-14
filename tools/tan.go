package tools

import (
	"context"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type TanInput struct {
	Angle float64 `json:"angle" jsonschema:"angle in radians"`
}

type TanOutput struct {
	Result float64 `json:"result" jsonschema:"tangent of the angle"`
}

func Tan(ctx context.Context, req *mcp.CallToolRequest, input TanInput) (
	*mcp.CallToolResult,
	TanOutput,
	error,
) {
	return nil, TanOutput{Result: math.Tan(input.Angle)}, nil
}
