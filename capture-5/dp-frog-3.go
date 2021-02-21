package main

import (
	"fmt"
	"math"
)

func chmin(a *float64, b float64) {
	if *a > b {
		*a = b
	}
}

func main() {
	var INF float64 = 10 ^ 5
	var N int
	fmt.Scan(&N)
	h := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&h[i])
	}
	dp := make([]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	// ループ
	for i := 1; i < N; i++ {
		chmin(&dp[i], dp[i-1]+math.Abs(h[i]-h[i-1]))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+math.Abs(h[i]-h[i-2]))
		}
	}
	fmt.Println(dp[N-1])
}
