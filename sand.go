package main

import (
	"fmt"
	"math"
	"reflect"

	"./pack"
)

func hello() {
	fmt.Print(">>> ")
	fmt.Printf("Hello, %s!\n", "世界") // string formatting
}

// Returns multiple values as named results
func factor(a, b int) (sum, multi int) {
	sum = a + b
	multi = a * b
	return // returns both sum and multi as named results
}

func typesAndReturns() {

	var f = 3.14159
	fmt.Println(int(f)) // type casting, float -> int

	// you can also: sum, multi := factor(5, 2)
	fmt.Println(factor(5, 2))
}

func packages() {
	// import a local package: import "./pack"
	fmt.Println(pack.Expo())
}

func pointers() {
	f := 22 / 7.0
	var p *float64 = &f
	fmt.Println(*p)

	*p = 2.26 // dereferencing or indirecting
	fmt.Println(f)
}

func structs() {

	type Vertex struct {
		x int
		y int
	}

	var point Vertex = Vertex{3, 4}
	fmt.Printf("%T -> %+v\n", point, point) // %+v prints the field names, too!

	p := &point
	// both ways to access the struct through a pointer are permitted:
	fmt.Println((*p).x, p.y)
}

func arrays() {

	primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	firstFive := primes[0:5] // a named slice - a pointer to a region

	fmt.Println("The first 5 primes are:", firstFive)
	fmt.Println("The next 3 are:", primes[5:8]) // slice[startIndex:endIndex]
	fmt.Printf("len = %d cap = %d slice: %v\n", len(firstFive), cap(firstFive), firstFive)

	oddSlice := []int{1, 3, 5, 7, 9}
	fmt.Println(reflect.TypeOf(oddSlice), reflect.TypeOf(primes))

	dynSlice := make([]int, 1, cap(primes)+1) // make a new slice, len 1, cap 25
	dynSlice = dynSlice[:len(primes)]         // extending the slice within the capacity
	copy(dynSlice, primes[:])
	fmt.Printf("len = %d cap = %d slice: %v\n", len(dynSlice), cap(dynSlice), dynSlice)

	// append a slice beyond its capacity (auto copy + redirect):
	dynSlice = append(dynSlice, 31, 37, 41)
	fmt.Printf("len = %d cap = %d slice: %v\n", len(dynSlice), cap(dynSlice), dynSlice)

	fmt.Print("Looping with 'range' (i=>v): ")
	for i, v := range dynSlice[:5] {
		fmt.Printf("%d => %v | ", i, v)
	}
	fmt.Println()

	s := []struct { // a custom struct slice + init
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s)
}

func Pic(dx, dy int) [][]uint8 {

	var result [][]uint8 = make([][]uint8, dx, dx)

	for i := 0; i < dx; i++ {
		result[i] = make([]uint8, dy, dy)
		for j := 0; j < dy; j++ {
			result[i][j] = uint8(i ^ j)
		}
	}
	return result
}

func maps() {
	m := make(map[string]string) // init an empty map
	p := map[float64]string{     // init a const map
		math.Pow(1024, 1): "K",
		math.Pow(1024, 2): "M",
		math.Pow(1024, 3): "G",
	}
	m["hello"] = "world"
	m["del"] = "me"
	delete(m, "del") // delete(map, key)
	fmt.Println(m, p)

	for i := 1.0; i < 5000; i *= 1024 {
		e, present := p[i]
		fmt.Println(present)
		if present {
			fmt.Printf("%.0f is present: %v\n", i, e)
		} else {
			fmt.Printf("%.0f is NOT present\n", i)
		}
	}
}

func main() {
	// hello()

	// typesAndReturns()

	// packages()

	// pointers()

	// structs()

	// arrays()

	maps()

}
