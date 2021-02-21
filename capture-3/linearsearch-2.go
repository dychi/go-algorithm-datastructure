package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	var N int
	N = nextInt()
	v := nextInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] += nextInt()
	}
	// 線形探索
	foundID := -1
	for i := 0; i < N; i++ {
		if a[i] == v {
			foundID = i
			break
		}
	}
	fmt.Println(foundID)
}
