package tools

import (
	"context"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type CosInput struct {
	Angle float64 `json:"angle" jsonschema:"angle in radians"`
}

type CosOutput struct {
	Result float64 `json:"result" jsonschema:"cosine of the angle"`
}

func Cos(ctx context.Context, req *mcp.CallToolRequest, input CosInput) (
	*mcp.CallToolResult,
	CosOutput,
	error,
) {
	return nil, CosOutput{Result: math.Cos(input.Angle)}, nil
}
