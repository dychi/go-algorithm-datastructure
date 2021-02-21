package main

import "fmt"

var memo [1000]int

func cfibo(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	if memo[N] > 0 {
		return memo[N]
	}
	memo[N] = cfibo(N-1) + cfibo(N-2)
	return memo[N]
}

func main() {
	fmt.Println(cfibo(40))
}
