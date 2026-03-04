package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// as prohibited goto and loop ---> RECURSION FUNC

func parseInts(fields []string, index int, result []int) []int {
	if index >= len(fields) {
		return result // if we've processed all fields, return the result
	}
	num, _ := strconv.Atoi(fields[index]) // parse to INT 
	return parseInts(fields, index+1, append(result, num))  
}

func sumNonPositivePow4(yn []int, index int) int {
	if index >= len(yn) { // if we've processed all elements, return 0
		return 0
	}
	rest := sumNonPositivePow4(yn, index+1) 
	if yn[index] <= 0 { // end := non positive ? -> calc pow4 + rest : just rest
		v := yn[index]
		return v*v*v*v + rest 
	}
	return rest
}

func processTestCases(scanner *bufio.Scanner, n int, results []int) []int {
	if n <= 0 {
		return results
	}
	scanner.Scan()
	x, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	if x != len(fields) {
		return processTestCases(scanner, n-1, append(results, -1))
	}

	yn := parseInts(fields, 0, []int{})
	sum := sumNonPositivePow4(yn, 0)
	return processTestCases(scanner, n-1, append(results, sum))
}

func printResults(results []int, index int) {
	if index >= len(results) {
		return
	}
	fmt.Println(results[index])
	printResults(results, index+1) 
}

func main() { // driver func 
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	results := processTestCases(scanner, n, []int{})
	printResults(results, 0)
}
