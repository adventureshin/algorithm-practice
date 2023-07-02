package main

import (
	"fmt"
)

func main() {
	slot := make([]bool, 1000001)

	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var number int
		fmt.Scan(&number)
		slot[number] = true
	}

	for i := 1000000; i >= 1; i-- {
		if slot[i] {
			fmt.Println(i)
		}
	}
}
