package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// Input the coefficients from Python to compare
var expectedCoefficients = map[string]float64{
	"coef1": 0.5001,
	"coef2": 0.5000,
	"coef3": 0.4997,
	"coef4": 0.4999,
}

// Anscombe Quartet data
var c1 = []stats.Coordinate{
	{X: 10, Y: 8.04},
	{X: 8, Y: 6.95},
	{X: 13, Y: 7.58},
	{X: 9, Y: 8.81},
	{X: 11, Y: 8.33},
	{X: 14, Y: 9.96},
	{X: 6, Y: 7.24},
	{X: 4, Y: 4.26},
	{X: 12, Y: 10.84},
	{X: 7, Y: 4.82},
	{X: 5, Y: 5.68},
}

var c2 = []stats.Coordinate{
	{X: 10, Y: 9.14},
	{X: 8, Y: 8.14},
	{X: 13, Y: 8.74},
	{X: 9, Y: 8.77},
	{X: 11, Y: 9.26},
	{X: 14, Y: 8.10},
	{X: 6, Y: 6.13},
	{X: 4, Y: 3.10},
	{X: 12, Y: 9.13},
	{X: 7, Y: 7.26},
	{X: 5, Y: 4.74},
}

var c3 = []stats.Coordinate{
	{X: 10, Y: 7.46},
	{X: 8, Y: 6.77},
	{X: 13, Y: 12.74},
	{X: 9, Y: 7.11},
	{X: 11, Y: 7.81},
	{X: 14, Y: 8.84},
	{X: 6, Y: 6.08},
	{X: 4, Y: 5.39},
	{X: 12, Y: 8.15},
	{X: 7, Y: 6.42},
	{X: 5, Y: 5.73},
}

var c4 = []stats.Coordinate{
	{X: 8, Y: 6.58},
	{X: 8, Y: 5.76},
	{X: 8, Y: 7.71},
	{X: 8, Y: 8.84},
	{X: 8, Y: 8.47},
	{X: 8, Y: 7.04},
	{X: 8, Y: 5.25},
	{X: 19, Y: 12.50},
	{X: 8, Y: 5.56},
	{X: 8, Y: 7.91},
	{X: 8, Y: 6.89},
}

// Test to compare the coefficients
func TestCalculateCoefficient(t *testing.T) {
	// Linear Regression
	r1, _ := stats.LinearRegression(c1)
	r2, _ := stats.LinearRegression(c2)
	r3, _ := stats.LinearRegression(c3)
	r4, _ := stats.LinearRegression(c4)

	// Calculate coefficients
	coefficient1, _ := CalculateCoefficient(r1)
	coefficient2, _ := CalculateCoefficient(r2)
	coefficient3, _ := CalculateCoefficient(r3)
	coefficient4, _ := CalculateCoefficient(r4)

	// Compare calculated coefficient with the
	// one from Python
	expected1 := expectedCoefficients["coef1"]
	if diff := coefficient1 - expected1; diff > 0.001 {
		t.Errorf("Coefficient mismatch with c1: got %v, want %v", coefficient1, expected1)
	}

	expected2 := expectedCoefficients["coef2"]
	if diff := coefficient2 - expected2; diff > 0.001 {
		t.Errorf("Coefficient mismatch with c2: got %v, want %v", coefficient2, expected2)
	}

	expected3 := expectedCoefficients["coef3"]
	if diff := coefficient3 - expected3; diff > 0.001 {
		t.Errorf("Coefficient mismatch with c3: got %v, want %v", coefficient3, expected3)
	}

	expected4 := expectedCoefficients["coef4"]
	if diff := coefficient4 - expected4; diff > 0.001 {
		t.Errorf("Coefficient mismatch with c4: got %v, want %v", coefficient1, expected4)
	}
}

// Run benchmark to get execution time of linear
// regression and calculate coefficient
func BenchmarkLinRegCoef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regData1, _ := stats.LinearRegression(c1)
		_, _ = CalculateCoefficient(regData1)

		regData2, _ := stats.LinearRegression(c2)
		_, _ = CalculateCoefficient(regData2)

		regData3, _ := stats.LinearRegression(c3)
		_, _ = CalculateCoefficient(regData3)

		regData4, _ := stats.LinearRegression(c4)
		_, _ = CalculateCoefficient(regData4)
	}
}
