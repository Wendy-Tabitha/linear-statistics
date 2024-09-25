# Linear-statistics

## Description

This program reads data from a file, where each line represents a y-value of a graph, and the corresponding x-values are the line numbers (0, 1, 2, ...). The program calculates two key statistics:
- Linear Regression Line: A formula representing the trend line of the data.
- Pearson Correlation Coefficient: A measure of the strength of the linear relationship between the x and y values.
The program reads the data file, computes the necessary values, and prints the results in a formatted output.

## Features

- [Linear Regression](https://en.wikipedia.org/wiki/Linear_regression)
- [Pearson Correlation Coefficient](https://en.wikipedia.org/wiki/Pearson_correlation_coefficient)

## Instructions

To clone this repository: use the command below on your terminal:
```bash
git clone https://github.com/Wendy-Tabitha/linear-statistics.git
```

The program should be able to read data from a file and compute each of the forementioned statistics. The input data will be provided in a file specified as a command-line argument. Each line in the file contains one value, representing a statistical population.

To execute the program, users will use a command below
```bash
go run . data.txt
```
OR:

```bash
go run main.go data.txt
```
After reading the file, your program should compute each of the requested statistics and print the results as follows (values are rounded integers):

The output should look like the example below:
```bash
Linear Regression Line: y = <value>x + <value>
Pearson Correlation Coefficient: <value>
```


## Testing

This project has test files which test the fuctionality of the functions that calculates the mean, regression and correlation of a data set. 

To run the test file:

```bash
go test -v
```
We run the above command in the console or terminal to test the files that are in the statistics directory.

## Contribution Guidelines (if applicable)

- Fork the Repository: If you intend to contribute changes, fork the repository from github to create your own copy.

- Branching Strategy: Create a new branch for your specific development work to avoid conflicts with the main branch.

- Pull Requests: When ready, submit a pull request to propose merging your changes back into the main repository.

## Learning Objectives

This project is designed to help learn about:

- Linear Regression.
- Pearson Correlation Coefficient.

This project was done by:
[Wendy Akinyi](https://github.com/Wendy-Tabitha/linear-statistics)