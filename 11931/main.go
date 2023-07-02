package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}

	sort.Stable(sort.Reverse(sort.IntSlice(numbers)))

	// Print numbers
	for _, number := range numbers {
		fmt.Println(number)
	}
}
