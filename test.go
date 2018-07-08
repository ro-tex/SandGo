package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < math.MaxUint32; i++ {
	}
	fmt.Println(time.Since(start))

	start = time.Now()
	for i := 0; i < math.MaxUint32; i += 1 {
	}
	fmt.Println(time.Since(start))
}
