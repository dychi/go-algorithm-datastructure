package main

import "fmt"

func fibo(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fibo(N-1) + fibo(N-2)
}

func main() {
	fmt.Println(fibo(9))
}
