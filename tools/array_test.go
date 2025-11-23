package tools

import (
	"context"
	"testing"
)

func TestSumArray(t *testing.T) {
	tests := []struct {
		name     string
		input    SumArrayInput
		expected float64
	}{
		{"positive numbers", SumArrayInput{Numbers: []float64{1, 2, 3, 4, 5}}, 15},
		{"negative numbers", SumArrayInput{Numbers: []float64{-1, -2, -3}}, -6},
		{"mixed numbers", SumArrayInput{Numbers: []float64{-5, 5, -3, 3}}, 0},
		{"single element", SumArrayInput{Numbers: []float64{42}}, 42},
		{"empty array", SumArrayInput{Numbers: []float64{}}, 0},
		{"decimal numbers", SumArrayInput{Numbers: []float64{1.5, 2.5, 3.0}}, 7.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := SumArray(context.Background(), nil, tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result.Sum != tt.expected {
				t.Errorf("SumArray(%v) = %v, want %v", tt.input.Numbers, result.Sum, tt.expected)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name      string
		input     AverageInput
		expected  float64
		expectErr bool
	}{
		{"positive numbers", AverageInput{Numbers: []float64{2, 4, 6, 8}}, 5, false},
		{"negative numbers", AverageInput{Numbers: []float64{-2, -4, -6}}, -4, false},
		{"mixed numbers", AverageInput{Numbers: []float64{-5, 5, -10, 10}}, 0, false},
		{"single element", AverageInput{Numbers: []float64{42}}, 42, false},
		{"empty array", AverageInput{Numbers: []float64{}}, 0, true},
		{"decimal result", AverageInput{Numbers: []float64{1, 2, 3}}, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Average(context.Background(), nil, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result.Average != tt.expected {
					t.Errorf("Average(%v) = %v, want %v", tt.input.Numbers, result.Average, tt.expected)
				}
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name      string
		input     MinInput
		expected  float64
		expectErr bool
	}{
		{"positive numbers", MinInput{Numbers: []float64{5, 2, 8, 1, 9}}, 1, false},
		{"negative numbers", MinInput{Numbers: []float64{-2, -8, -1, -5}}, -8, false},
		{"mixed numbers", MinInput{Numbers: []float64{5, -3, 8, -10, 2}}, -10, false},
		{"single element", MinInput{Numbers: []float64{42}}, 42, false},
		{"empty array", MinInput{Numbers: []float64{}}, 0, true},
		{"all same values", MinInput{Numbers: []float64{7, 7, 7}}, 7, false},
		{"decimal numbers", MinInput{Numbers: []float64{1.5, 2.3, 0.9, 1.2}}, 0.9, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Min(context.Background(), nil, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result.Min != tt.expected {
					t.Errorf("Min(%v) = %v, want %v", tt.input.Numbers, result.Min, tt.expected)
				}
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name      string
		input     MaxInput
		expected  float64
		expectErr bool
	}{
		{"positive numbers", MaxInput{Numbers: []float64{5, 2, 8, 1, 9}}, 9, false},
		{"negative numbers", MaxInput{Numbers: []float64{-2, -8, -1, -5}}, -1, false},
		{"mixed numbers", MaxInput{Numbers: []float64{5, -3, 8, -10, 2}}, 8, false},
		{"single element", MaxInput{Numbers: []float64{42}}, 42, false},
		{"empty array", MaxInput{Numbers: []float64{}}, 0, true},
		{"all same values", MaxInput{Numbers: []float64{7, 7, 7}}, 7, false},
		{"decimal numbers", MaxInput{Numbers: []float64{1.5, 2.3, 0.9, 1.2}}, 2.3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, result, err := Max(context.Background(), nil, tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result.Max != tt.expected {
					t.Errorf("Max(%v) = %v, want %v", tt.input.Numbers, result.Max, tt.expected)
				}
			}
		})
	}
}
