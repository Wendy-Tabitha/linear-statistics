package statistics

// Calculates the mean (average) of a slice of float64 numbers.
func Mean(num []float64) float64 {
	total := 0.0
	n := len(num)
	for _, nb := range num {
		total += nb
	}
	return total / float64(n)
}
