package main

import "fmt"

func main() {
	var T, N, multiples_of_3, multiples_of_5, multiples_of_15, sum int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		multiples_of_3 = (N - 1) / 3
		multiples_of_5 = (N - 1) / 5
		multiples_of_15 = (N - 1) / 15
		sum = 3*multiples_of_3*(multiples_of_3+1)/2 +
			5*multiples_of_5*(multiples_of_5+1)/2 -
			15*multiples_of_15*(multiples_of_15+1)/2
		fmt.Println(sum)
	}
}
