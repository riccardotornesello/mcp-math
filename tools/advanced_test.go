package tools

import (
	"context"
	"math"
	"testing"
)

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		input    PowerInput
		expected float64
	}{
		{"positive base and exponent", PowerInput{Base: 2, Exponent: 3}, 8},
		{"negative base, even exponent", PowerInput{Base: -2, Exponent: 2}, 4},
		{"negative base, odd exponent", PowerInput{Base: -2, Exponent: 3}, -8},
		{"zero exponent", PowerInput{Base: 5, Exponent: 0}, 1},
		{"zero base", PowerInput{Base: 0, Exponent: 5}, 0},
		{"negative exponent", PowerInput{Base: 2, Exponent: -2}, 0.25},
		{"decimal base", PowerInput{Base: 2.5, Exponent: 2}, 6.25},
		{"decimal exponent", PowerInput{Base: 4, Exponent: 0.5}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Power(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// Use a small epsilon for floating point comparison
			if math.Abs(result.Result-tt.expected) > 1e-9 {
				t.Errorf("Power(%v, %v) = %v, want %v", tt.input.Base, tt.input.Exponent, result.Result, tt.expected)
			}
		})
	}
}

func TestRoot(t *testing.T) {
	tests := []struct {
		name      string
		input     RootInput
		expected  float64
		expectErr bool
	}{
		{"square root", RootInput{Number: 16, N: 2}, 4, false},
		{"cube root", RootInput{Number: 27, N: 3}, 3, false},
		{"fourth root", RootInput{Number: 81, N: 4}, 3, false},
		{"root of 1", RootInput{Number: 1, N: 5}, 1, false},
		{"root of 0", RootInput{Number: 0, N: 2}, 0, false},
		{"zero root degree", RootInput{Number: 16, N: 0}, 0, true},
		{"decimal number", RootInput{Number: 6.25, N: 2}, 2.5, false},
		{"decimal root degree", RootInput{Number: 16, N: 2.5}, 0, false}, // specific value check below
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Root(context.Background(), nil, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				// For the decimal root degree test, we need special handling
				if tt.name == "decimal root degree" {
					// 16^(1/2.5) = 16^0.4 ≈ 3.0314...
					if math.Abs(result.Result-math.Pow(16, 1.0/2.5)) > 1e-9 {
						t.Errorf("Root(%v, %v) = %v, want %v", tt.input.Number, tt.input.N, result.Result, math.Pow(16, 1.0/2.5))
					}
				} else {
					// Use a small epsilon for floating point comparison
					if math.Abs(result.Result-tt.expected) > 1e-9 {
						t.Errorf("Root(%v, %v) = %v, want %v", tt.input.Number, tt.input.N, result.Result, tt.expected)
					}
				}
			}
		})
	}
}
