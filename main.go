package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"linear-stats/statistics"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run . data.txt")
		return
	}

	first, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	data := string(first)

	if !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("Use a txt file")
		return
	}

	y := []float64{}
	maxInt := math.MaxInt
	g := strings.Split(data, "\n")
	for _, str := range g {

		str = strings.TrimSpace(str)

		if str == "" {
			continue
		}

		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error converting the string to float", err)
			return
		}

		y = append(y, num)
	}
	if len(y) == 0 {
		return
	}
	if len(y) == 1 {
		fmt.Println("Only one value in data set")
		return
	}


	for _, nb := range y {
		if nb > float64(maxInt) {
			fmt.Println("Number is too large")
			return
		}
	}
	n := float64(len(y))
	var x []float64
	for i := 0.0; i < n; i++ {
		x = append(x, i)
	}

	// Calculate means of x and y
	xMean, yMean := statistics.Mean(x), statistics.Mean(y)

	m, b := statistics.Regression(x, y, xMean, yMean)

	r := statistics.Correlation(x, y, xMean, yMean)

	// Output results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
