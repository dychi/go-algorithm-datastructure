package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	h := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&h[i])
	}
	dp := make([]float64, N)
	// 初期条件
	dp[0] = 0
	for i := 1; i < N; i++ {
		if i == 1 {
			dp[i] = math.Abs(h[i] - h[i-1])
		} else {
			dp[i] = math.Min(dp[i-1]+math.Abs(h[i]-h[i-1]), dp[i-2]+math.Abs(h[i]-h[i-2]))
		}
	}
	fmt.Println(dp[N-1])
}
