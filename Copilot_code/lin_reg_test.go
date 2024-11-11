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

func TestCalculateCoefficient(t *testing.T) {
	testCases := []struct {
		data     []stats.Coordinate
		expected float64
		label    string
	}{
		{c1, expectedCoefficients["coef1"], "c1"},
		{c2, expectedCoefficients["coef2"], "c2"},
		{c3, expectedCoefficients["coef3"], "c3"},
		{c4, expectedCoefficients["coef4"], "c4"},
	}

	for _, tc := range testCases {
		r, err := stats.LinearRegression(tc.data)
		if err != nil {
			t.Fatalf("Error calculating linear regression for %s: %v", tc.label, err)
		}
		coefficient, err := CalculateCoefficient(r)
		if err != nil {
			t.Fatalf("Error calculating coefficient for %s: %v", tc.label, err)
		}
		if coefficient != tc.expected {
			t.Errorf("Coefficient for %s: expected %f, got %f", tc.label, tc.expected, coefficient)
		}
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
