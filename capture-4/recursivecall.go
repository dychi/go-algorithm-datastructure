package main

import (
	"fmt"
)

func recursive(N int) int {
	fmt.Printf("recursive(%d) を呼び出しました\n", N)
	if N == 0 {
		return 0
	}
	result := N + recursive(N-1)
	fmt.Printf("%d までの和: %d\n", N, result)
	return result
}

func main() {
	recursive(5)
}
