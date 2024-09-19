package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

	m, b := Regression(x, y, xMean, yMean)

	r := Correlation(x, y, xMean, yMean)

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
