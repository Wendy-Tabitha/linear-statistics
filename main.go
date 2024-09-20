package main

import (
	"fmt"
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
		os.Exit(1)
	}

	y := []float64{}
	g := strings.Split(data, "\n")
	for _, str := range g {

		str = strings.TrimSpace(str)

		if str == "" {
			continue
		}

		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error converting the string to float", err)
			os.Exit(1)
		}

		y = append(y, num)
	}
	if len(y) == 0 {
		return
	}
	n := float64(len(y))
	var x []float64
	for i := 0.0; i < n; i++ {
		x = append(x, i)
	}

	// Calculate means of x and y
	xMean, yMean := Mean(x), Mean(y)

	m, b := statistics.Regression(x, y, xMean, yMean)

	r := statistics.Correlation(x, y, xMean, yMean)

	// Output results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

func Mean(num []float64) float64 {
	total := 0.0
	n := len(num)
	for _, nb := range num {
		total += nb
	}
	return total / float64(n)
}
