package main

import (
	"fmt"
)

func main() {
	slot := make([]bool, 2000002)

	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var number int
		fmt.Scan(&number)
		slot[number + 1000000] = true
	}

	for i := 2000001; i >= 0; i-- {
		if slot[i] {
			fmt.Println(i - 1000000)
		}
	}
}
