package statistics

import "testing"

func TestRegression(t *testing.T) {
	// Original dataset (mixed values)
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []float64{23, 32, 67, 5, 453, 345, 3, 23423, 52, 5}
	xMean := Mean(x)
	yMean := Mean(y)

	// Positive values
	xPos := []float64{1, 2, 3, 4, 5}
	yPos := []float64{2, 4, 6, 8, 10}
	xPosMean := Mean(xPos)
	yPosMean := Mean(yPos)

	// Negative values
	xNeg := []float64{-1, -2, -3, -4, -5}
	yNeg := []float64{-2, -4, -6, -8, -10}
	xNegMean := Mean(xNeg)
	yNegMean := Mean(yNeg)

	type args struct {
		x     []float64
		y     []float64
		xMean float64
		yMean float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// Original mixed dataset
		{"Mixed Regression", args{x, y, xMean, yMean}, 706.9333333333333, -740.3999999999996},

		// Positive correlation test
		{"Positive Regression", args{xPos, yPos, xPosMean, yPosMean}, 2.0, 0.0}, // y = 2x, b = 0

		// Negative correlation test
		{"Negative Regression", args{xNeg, yNeg, xNegMean, yNegMean}, 2.0, 0.0}, // y = 2x, b = 0 (for negative x and y)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Regression(tt.args.x, tt.args.y, tt.args.xMean, tt.args.yMean)
			if got != tt.want {
				t.Errorf("Regression() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Regression() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
