package main

import (
	"golang/1-Prog/calculator"
	"testing"
)

func TestMathFunctions(t *testing.T) {
	testCases := []struct {
		name           string
		operation      func(a, b int) int
		a, b, expected int
	}{
		{"Addition", calculator.Addition, 3, 5, 8},
		{"Addition", calculator.Addition, 0, 7, 8},
		{"Addition", calculator.Addition, 0, 2, 8},
		{"Subtract", calculator.Subtract, 8, 3, 5},
		{"Subtract", calculator.Subtract, -5, -2, 3},
		{"Multiplication", calculator.Multiplication, 8, 3, 24},
		{"Multiplication", calculator.Multiplication, -5, -2, 3},
		{"Multiplication", calculator.Multiplication, 5, -2, 3},
		{"Multiplication", calculator.Multiplication, -5, 2, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.operation(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("%s(%d, %d) returned %d; expected %d", tc.name, tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
