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
	N := nextInt()
	W := nextInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}
	exist := false
	// bit は 2^N 通りの部分集合全体を動く
	for bit := 0; bit < (1 << N); bit++ {
		var sum int = 0
		// 配列の値を取り出す
		for i := 0; i < N; i++ {
			// bit の値で示される部分集合に含まれているのかどうか ex) 01010 & 00010 => 00010
			if bit>>i&1 == 1 {
				sum += a[i]
			}
		}
		// sum が W に一致するかどうか
		if sum == W {
			exist = true
		}
	}
	fmt.Printf("N: %d\n", N)
	fmt.Printf("W: %d\n", W)
	fmt.Printf("a: %d\n", a)
	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
