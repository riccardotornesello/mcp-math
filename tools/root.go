package tools

import (
	"context"
	"errors"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type RootInput struct {
	Number float64 `json:"number" jsonschema:"the number to find the root of"`
	N      float64 `json:"n" jsonschema:"the degree of the root (e.g., 2 for square root, 3 for cube root)"`
}

type RootOutput struct {
	Result float64 `json:"result" jsonschema:"the nth root of the number"`
}

func Root(ctx context.Context, req *mcp.CallToolRequest, input RootInput) (
	*mcp.CallToolResult,
	RootOutput,
	error,
) {
	if input.N == 0 {
		return nil, RootOutput{}, errors.New("root degree cannot be zero")
	}
	result := math.Pow(input.Number, 1.0/input.N)
	return nil, RootOutput{Result: result}, nil
}
