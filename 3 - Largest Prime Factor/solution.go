package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(number int) int {
	largest := 2

	for number%2 == 0 {
		number /= 2
	}

	factor := 3
	maxFactor := int(math.Sqrt(float64(number)))

	for factor <= maxFactor {
		if number%factor == 0 {
			largest = factor
			number /= factor

			for number%factor == 0 {
				number /= factor
			}

			if number == 1 {
				return largest
			}

			maxFactor = int(math.Sqrt(float64(number)))
		}

		factor += 2
	}

	return number
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		result := largestPrimeFactor(N)
		fmt.Println(result)
	}
}
