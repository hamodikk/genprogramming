package main

import (
	"fmt"
	"log"
	"time"

	"github.com/montanaflynn/stats"
)

// Anscombe data contains four datasets demonstrating
// identical linear regression properties despite differing values.
var anscombe = map[string][]float64{
	"x1": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x2": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x3": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x4": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	"y1": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	"y2": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	"y3": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	"y4": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
}

// linearRegression computes the slope and intercept of the linear regression
// line for given x and y data slices. It returns an error if the input data
// is invalid.
func linearRegression(x, y []float64) (float64, float64, error) {
	if len(x) != len(y) || len(x) == 0 {
		return 0, 0, fmt.Errorf("invalid input data: x and y must be non-empty and of equal length")
	}

	xMean, err := stats.Mean(x)
	if err != nil {
		return 0, 0, fmt.Errorf("error calculating x mean: %v", err)
	}
	yMean, err := stats.Mean(y)
	if err != nil {
		return 0, 0, fmt.Errorf("error calculating y mean: %v", err)
	}

	var numerator, denominator float64
	for i := range x {
		xDiff := x[i] - xMean
		numerator += xDiff * (y[i] - yMean)
		denominator += xDiff * xDiff
	}

	if denominator == 0 {
		return 0, 0, fmt.Errorf("error in regression: zero denominator")
	}

	slope := numerator / denominator
	intercept := yMean - slope*xMean
	return slope, intercept, nil
}

// fitAllModels performs linear regression on each dataset in Anscombe's quartet.
// It outputs the slope and intercept of each regression line.
func fitAllModels() {
	datasets := []struct {
		x, y  []float64
		label string
	}{
		{anscombe["x1"], anscombe["y1"], "Dataset I"},
		{anscombe["x2"], anscombe["y2"], "Dataset II"},
		{anscombe["x3"], anscombe["y3"], "Dataset III"},
		{anscombe["x4"], anscombe["y4"], "Dataset IV"},
	}

	for _, dataset := range datasets {
		slope, intercept, err := linearRegression(dataset.x, dataset.y)
		if err != nil {
			log.Printf("Error calculating regression for %s: %v\n", dataset.label, err)
			continue
		}
		fmt.Printf("%s: Slope = %.4f, Intercept = %.4f\n", dataset.label, slope, intercept)
	}
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		fitAllModels()
	}
	duration := time.Since(start).Seconds()
	fmt.Printf("Average execution time for fitting all models over 10000 runs: %.4f seconds\n", duration/10000)
}
