package statistics

import "math"

func Regression(x, y []float64, xMean, yMean float64) (float64, float64) {
	var num, denom float64
	for i := 0; i < len(x); i++ {
		num += (x[i] - xMean) * (y[i] - yMean)
		denom += (x[i] - xMean) * (x[i] - xMean)
	}
	m := num / denom
	b := yMean - m*xMean
	return m, b
}

func Correlation(x, y []float64, xMean, yMean float64) float64 {
	var num, denomX, denomY float64
	for i := 0; i < len(x); i++ {
		num += (x[i] - xMean) * (y[i] - yMean)
		denomX += (x[i] - xMean) * (x[i] - xMean)
		denomY += (y[i] - yMean) * (y[i] - yMean)
	}
	return num / math.Sqrt(denomX*denomY)
}
