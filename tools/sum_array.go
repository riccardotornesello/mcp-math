package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SumArrayInput struct {
	Numbers []float64 `json:"numbers" jsonschema:"array of numbers to sum"`
}

type SumArrayOutput struct {
	Sum float64 `json:"sum" jsonschema:"the sum of all numbers in the array"`
}

func SumArray(ctx context.Context, req *mcp.CallToolRequest, input SumArrayInput) (
	*mcp.CallToolResult,
	SumArrayOutput,
	error,
) {
	sum := 0.0
	for _, num := range input.Numbers {
		sum += num
	}
	return nil, SumArrayOutput{Sum: sum}, nil
}
