package main

import (
	"testing"
)

// testCases define the expected slope and intercept values for each dataset.
// These values are based on the known regression results of Anscombe's quartet.
var testCases = []struct {
	label             string
	x, y              []float64
	expectedSlope     float64
	expectedIntercept float64
}{
	{"Dataset I", anscombe["x1"], anscombe["y1"], 0.5001, 3.0001},
	{"Dataset II", anscombe["x2"], anscombe["y2"], 0.5001, 3.0001},
	{"Dataset III", anscombe["x3"], anscombe["y3"], 0.5001, 3.0001},
	{"Dataset IV", anscombe["x4"], anscombe["y4"], 0.5001, 3.0001},
}

// TestLinearRegression verifies that linear regression calculations produce
// expected results for each dataset in Anscombe's quartet.
func TestLinearRegression(t *testing.T) {
	for _, tc := range testCases {
		slope, intercept, err := linearRegression(tc.x, tc.y)
		if err != nil {
			t.Fatalf("Error calculating regression for %s: %v", tc.label, err)
		}
		if !almostEqual(slope, tc.expectedSlope, 0.001) {
			t.Errorf("%s: Expected slope %.4f, got %.4f", tc.label, tc.expectedSlope, slope)
		}
		if !almostEqual(intercept, tc.expectedIntercept, 0.001) {
			t.Errorf("%s: Expected intercept %.4f, got %.4f", tc.label, tc.expectedIntercept, intercept)
		}
	}
}

// almostEqual checks if two float64 values are approximately equal within a
// specified tolerance. It is used to handle floating-point comparison inaccuracies.
func almostEqual(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < tolerance
}

// BenchmarkLinearRegression measures the performance of the fitAllModels function.
func BenchmarkLinearRegression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fitAllModels()
	}
}
