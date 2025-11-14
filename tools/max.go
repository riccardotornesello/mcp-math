package tools

import (
	"context"
	"errors"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MaxInput struct {
	Numbers []float64 `json:"numbers" jsonschema:"array of numbers to find maximum"`
}

type MaxOutput struct {
	Max float64 `json:"max" jsonschema:"the maximum value in the array"`
}

func Max(ctx context.Context, req *mcp.CallToolRequest, input MaxInput) (
	*mcp.CallToolResult,
	MaxOutput,
	error,
) {
	if len(input.Numbers) == 0 {
		return nil, MaxOutput{}, errors.New("cannot find maximum of empty array")
	}
	max := math.Inf(-1)
	for _, num := range input.Numbers {
		if num > max {
			max = num
		}
	}
	return nil, MaxOutput{Max: max}, nil
}
