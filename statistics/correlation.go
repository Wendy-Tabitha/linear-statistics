package statistics

import "math"

// Calculates the Pearson correlation coefficient between two datasets x and y.
func Correlation(x, y []float64, xMean, yMean float64) float64 {
	var num, denomX, denomY float64
	for i := 0; i < len(x); i++ {
		num += (x[i] - xMean) * (y[i] - yMean)
		denomX += (x[i] - xMean) * (x[i] - xMean)
		denomY += (y[i] - yMean) * (y[i] - yMean)
	}
	return num / math.Sqrt(denomX*denomY)
}
