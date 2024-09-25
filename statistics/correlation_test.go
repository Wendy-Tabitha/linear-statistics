package statistics

import "testing"

func TestCorrelation(t *testing.T) {
	// Example datasets
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []float64{23, 32, 67, 5, 453, 345, 3, 23423, 52, 5} // Mixed correlation

	// Positive correlation example
	xPos := []float64{1, 2, 3, 4, 5}
	yPos := []float64{2, 4, 6, 8, 10} // Perfect positive correlation

	// Negative correlation example
	xNeg := []float64{1, 2, 3, 4, 5}
	yNeg := []float64{10, 8, 6, 4, 2} // Perfect negative correlation

	xMean := Mean(x)
	yMean := Mean(y)

	xPosMean := Mean(xPos)
	yPosMean := Mean(yPos)

	xNegMean := Mean(xNeg)
	yNegMean := Mean(yNeg)

	type args struct {
		x     []float64
		y     []float64
		xMean float64
		yMean float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Mixed Correlation", args{x, y, xMean, yMean}, 0.29025232355215574},
		{"Positive Correlation", args{xPos, yPos, xPosMean, yPosMean}, 1.0},  // Perfect positive correlation
		{"Negative Correlation", args{xNeg, yNeg, xNegMean, yNegMean}, -1.0}, // Perfect negative correlation
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Correlation(tt.args.x, tt.args.y, tt.args.xMean, tt.args.yMean)
			if got != tt.want {
				t.Errorf("Correlation() = %v, want %v", got, tt.want)
			}
		})
	}
}
