package tools

import (
	"context"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type AverageInput struct {
	Numbers []float64 `json:"numbers" jsonschema:"array of numbers to calculate average"`
}

type AverageOutput struct {
	Average float64 `json:"average" jsonschema:"the average of all numbers in the array"`
}

func Average(ctx context.Context, req *mcp.CallToolRequest, input AverageInput) (
	*mcp.CallToolResult,
	AverageOutput,
	error,
) {
	if len(input.Numbers) == 0 {
		return nil, AverageOutput{}, errors.New("cannot calculate average of empty array")
	}
	sum := 0.0
	for _, num := range input.Numbers {
		sum += num
	}
	return nil, AverageOutput{Average: sum / float64(len(input.Numbers))}, nil
}
