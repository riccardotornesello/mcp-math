package tools

import (
	"context"
	"errors"
	"math"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MinInput struct {
	Numbers []float64 `json:"numbers" jsonschema:"array of numbers to find minimum"`
}

type MinOutput struct {
	Min float64 `json:"min" jsonschema:"the minimum value in the array"`
}

func Min(ctx context.Context, req *mcp.CallToolRequest, input MinInput) (
	*mcp.CallToolResult,
	MinOutput,
	error,
) {
	if len(input.Numbers) == 0 {
		return nil, MinOutput{}, errors.New("cannot find minimum of empty array")
	}
	min := math.Inf(1)
	for _, num := range input.Numbers {
		if num < min {
			min = num
		}
	}
	return nil, MinOutput{Min: min}, nil
}
