package main

import (
	"fmt"
	"math/rand"
)

func is_sorted(sl []int) bool {
	if len(sl) <= 1 {
		return true
	}
	for i := 0; i < len(sl)-1; i++ {
		if sl[i] > sl[i+1] {
			return false
		}
	}
	return true
}

func get_perm(sl []int) []int {
	var out = make([]int, len(sl))
	var indices = rand.Perm(len(sl))
	for k, i := range indices {
		out[k] = sl[i]
	}
	return out
}

func bogosort(arr []int) []int {
	var out = get_perm(arr)
	for !is_sorted(out) {
		out = get_perm(arr)
	}
	return out
}

func main() {
	arr := []int{5, 0, 3, 4, 1}
	fmt.Println(bogosort(arr))
}
