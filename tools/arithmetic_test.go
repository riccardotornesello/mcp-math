package tools

import (
	"context"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    SumInput
		expected float64
	}{
		{"positive numbers", SumInput{A: 5, B: 3}, 8},
		{"negative numbers", SumInput{A: -5, B: -3}, -8},
		{"mixed numbers", SumInput{A: 5, B: -3}, 2},
		{"zero values", SumInput{A: 0, B: 0}, 0},
		{"decimal numbers", SumInput{A: 1.5, B: 2.5}, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Sum(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result.Sum != tt.expected {
				t.Errorf("Sum(%v, %v) = %v, want %v", tt.input.A, tt.input.B, result.Sum, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		input    SubtractInput
		expected float64
	}{
		{"positive numbers", SubtractInput{A: 5, B: 3}, 2},
		{"negative numbers", SubtractInput{A: -5, B: -3}, -2},
		{"mixed numbers", SubtractInput{A: 5, B: -3}, 8},
		{"zero values", SubtractInput{A: 0, B: 0}, 0},
		{"decimal numbers", SubtractInput{A: 5.5, B: 2.5}, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Subtract(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result.Result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.input.A, tt.input.B, result.Result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		input    MultiplyInput
		expected float64
	}{
		{"positive numbers", MultiplyInput{A: 5, B: 3}, 15},
		{"negative numbers", MultiplyInput{A: -5, B: -3}, 15},
		{"mixed numbers", MultiplyInput{A: 5, B: -3}, -15},
		{"zero multiplication", MultiplyInput{A: 5, B: 0}, 0},
		{"decimal numbers", MultiplyInput{A: 2.5, B: 4}, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Multiply(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result.Result != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.input.A, tt.input.B, result.Result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		input     DivideInput
		expected  float64
		expectErr bool
	}{
		{"positive numbers", DivideInput{A: 6, B: 3}, 2, false},
		{"negative numbers", DivideInput{A: -6, B: -3}, 2, false},
		{"mixed numbers", DivideInput{A: 6, B: -3}, -2, false},
		{"decimal result", DivideInput{A: 5, B: 2}, 2.5, false},
		{"division by zero", DivideInput{A: 5, B: 0}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Divide(context.Background(), nil, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result.Result != tt.expected {
					t.Errorf("Divide(%v, %v) = %v, want %v", tt.input.A, tt.input.B, result.Result, tt.expected)
				}
			}
		})
	}
}

func TestModulo(t *testing.T) {
	tests := []struct {
		name      string
		input     ModuloInput
		expected  float64
		expectErr bool
	}{
		{"positive numbers", ModuloInput{A: 7, B: 3}, 1, false},
		{"negative dividend", ModuloInput{A: -7, B: 3}, -1, false},
		{"decimal numbers", ModuloInput{A: 5.5, B: 2}, 1.5, false},
		{"modulo by zero", ModuloInput{A: 5, B: 0}, 0, true},
		{"exact division", ModuloInput{A: 6, B: 3}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Modulo(context.Background(), nil, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result.Result != tt.expected {
					t.Errorf("Modulo(%v, %v) = %v, want %v", tt.input.A, tt.input.B, result.Result, tt.expected)
				}
			}
		})
	}
}
