package main

import (
	"fmt"
	"math/rand"
)

func isSorted(sl []int) bool {
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

func getPerm(sl []int) []int {
	var out = make([]int, len(sl))
	var indices = rand.Perm(len(sl))
	for k, i := range indices {
		out[k] = sl[i]
	}
	return out
}

// Bogosort sorts a slice by generating random permutations until one of them happens to be sorted.
func Bogosort(arr []int) []int {
	var out = getPerm(arr)
	for !isSorted(out) {
		out = getPerm(arr)
	}
	return out
}

func main() {
	arr := []int{5, 0, 3, 4, 1}
	fmt.Println(Bogosort(arr))
}
