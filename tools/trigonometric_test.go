package tools

import (
	"context"
	"math"
	"testing"
)

func TestSin(t *testing.T) {
	tests := []struct {
		name     string
		input    SinInput
		expected float64
	}{
		{"zero", SinInput{Angle: 0}, 0},
		{"pi/2", SinInput{Angle: math.Pi / 2}, 1},
		{"pi", SinInput{Angle: math.Pi}, 0},
		{"3*pi/2", SinInput{Angle: 3 * math.Pi / 2}, -1},
		{"2*pi", SinInput{Angle: 2 * math.Pi}, 0},
		{"negative angle", SinInput{Angle: -math.Pi / 2}, -1},
		{"pi/6", SinInput{Angle: math.Pi / 6}, 0.5},
		{"pi/4", SinInput{Angle: math.Pi / 4}, math.Sqrt(2) / 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Sin(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// Use a small epsilon for floating point comparison
			if math.Abs(result.Result-tt.expected) > 1e-9 {
				t.Errorf("Sin(%v) = %v, want %v", tt.input.Angle, result.Result, tt.expected)
			}
		})
	}
}

func TestCos(t *testing.T) {
	tests := []struct {
		name     string
		input    CosInput
		expected float64
	}{
		{"zero", CosInput{Angle: 0}, 1},
		{"pi/2", CosInput{Angle: math.Pi / 2}, 0},
		{"pi", CosInput{Angle: math.Pi}, -1},
		{"3*pi/2", CosInput{Angle: 3 * math.Pi / 2}, 0},
		{"2*pi", CosInput{Angle: 2 * math.Pi}, 1},
		{"negative angle", CosInput{Angle: -math.Pi}, -1},
		{"pi/3", CosInput{Angle: math.Pi / 3}, 0.5},
		{"pi/4", CosInput{Angle: math.Pi / 4}, math.Sqrt(2) / 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Cos(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// Use a small epsilon for floating point comparison
			if math.Abs(result.Result-tt.expected) > 1e-9 {
				t.Errorf("Cos(%v) = %v, want %v", tt.input.Angle, result.Result, tt.expected)
			}
		})
	}
}

func TestTan(t *testing.T) {
	tests := []struct {
		name     string
		input    TanInput
		expected float64
	}{
		{"zero", TanInput{Angle: 0}, 0},
		{"pi/4", TanInput{Angle: math.Pi / 4}, 1},
		{"pi", TanInput{Angle: math.Pi}, 0},
		{"negative angle", TanInput{Angle: -math.Pi / 4}, -1},
		{"small angle", TanInput{Angle: 0.1}, math.Tan(0.1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Tan(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// Use a small epsilon for floating point comparison
			if math.Abs(result.Result-tt.expected) > 1e-9 {
				t.Errorf("Tan(%v) = %v, want %v", tt.input.Angle, result.Result, tt.expected)
			}
		})
	}
}
