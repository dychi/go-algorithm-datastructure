package main

import "fmt"

// 部分和問題を再帰関数で解く
func fullsearch(i int, w int, a []int) bool {
	if i == 0 {
		if w == 0 {
			return true
		}
		return false
	}
	// a[i-1]を選ばない場合
	if fullsearch(i-1, w, a) {
		return true
	}
	// a[i-1]を選ぶ場合
	if fullsearch(i-1, w-a[i-1], a) {
		return true
	}
	return false
}

func main() {
	var N, W int
	fmt.Scan(&N, &W)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	if fullsearch(N, W, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
