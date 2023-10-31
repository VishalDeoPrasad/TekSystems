package main

import (
	"fmt"
	"golang/1-Prog/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b    int
		actural int
	}{
		{1, 2, 3},
		{4, 6, 5},
		{2, 4, 6},
		{1, 3, 4},
		{2, 4, 4},
	}

	for _, tc := range testCases {
		actural := tc.actural
		ans := calculator.Add(tc.a, tc.b)
		if actural != ans {
			t.Errorf("calculator.Add(%d, %d) = %d; want %d", tc.a, tc.b, ans, actural)
			//t.Fatalf("calculator.Add(%d, %d) = %d; want %d", tc.a, tc.b, ans, actural)
			//t.Log(fmt.Printf("calculator.Add(%d, %d) = %d; want %d", tc.a, tc.b, ans, actural))
			//t.Fail()
			//t.FailNow()

		}
	}
}

func TestSub(t *testing.T) {
	// Test cases
	testCases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
		{8, -5, 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Sub(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			result := calculator.Sub(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("calculator.Sub(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestMulti(t *testing.T) {
	t.Run("PositiveNumbers", func(t *testing.T) {
		testCases := []struct {
			a, b, expected int
		}{
			{1, 2, 2},
			{0, 0, 0},
			{10, 5, 30},
			{8, 5, 40},
		}

		for _, tc := range testCases {
			result := calculator.Multi(tc.a, tc.b)
			expected := tc.expected
			if result != expected {
				t.Errorf("Mult(%d, %d) returned %d; expected %d", tc.a, tc.b, result, expected)
			}
		}
	})

	t.Run("NegativeNumbers", func(t *testing.T) {
		testCases := []struct {
			a, b, expected int
		}{
			{-1, -2, -2},
			{0, 0, 0},
			{-10, 5, -50},
			{-8, -5, 40},
		}

		for _, tc := range testCases {
			result := calculator.Multi(tc.a, tc.b)
			expected := tc.expected
			if result != expected {
				t.Errorf("Mult(%d, %d) returned %d; expected %d", tc.a, tc.b, result, expected)
			}
		}
	})
}
