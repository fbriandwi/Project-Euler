package main

import (
	"fmt"
)

func sumEvenFibonacciTerms(limit int) int {
	prevTerm, currentTerm := 1, 2
	evenSum := 0

	for currentTerm <= limit {
		if currentTerm%2 == 0 {
			evenSum += currentTerm
		}
		prevTerm, currentTerm = currentTerm, prevTerm+currentTerm
	}

	return evenSum
}

func main() {
	var numTestCases int
	fmt.Scan(&numTestCases)

	for i := 0; i < numTestCases; i++ {
		var limit int
		fmt.Scan(&limit)
		result := sumEvenFibonacciTerms(limit)
		fmt.Println(result)
	}
}
