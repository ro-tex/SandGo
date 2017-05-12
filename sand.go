package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"

	"github.com/SandGo/pack"
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

type Vertex struct {
	X int
	Y int
}

func structs() {
	var point Vertex = Vertex{3, 4}
	fmt.Printf("%T -> %+v\n", point, point) // %+v prints the field names, too!

	p := &point
	// both ways to access the struct through a pointer are permitted:
	fmt.Println((*p).X, p.Y)
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

func concurrency() {
	ch := make(chan interface{}) // define a channel that takes any kind of message

	go func() { // spin off a lambda into a separate goroutine
		ch <- struct { // send a custom struct to the channel
			x int
			s string
		}{2, "hello"}
	}() // end of lambda/goroutine

	// timeout := time.After(1 * time.Second) // returns a channel of type time.Time
	// _ = <-timeout                          // an implicit "sleep" because "<-" blocks

	fmt.Println(<-ch) // read from the channel and print the message out

	chBool := make(chan bool, 2) // add a buffer of two messages

	fmt.Println("Working...")

	go func() { chBool <- true }()
	go func() {
		_ = <-time.After(1 * time.Second) // do some "work"
		chBool <- true                    // all done
	}()

	if <-chBool && <-chBool { // wait for all goroutines to finish
		fmt.Println("Done")
	}
}

/*
	Channels aren't like files; you don't usually need to close them. Closing
	is only necessary when the receiver must be told there are no more values
	coming, such as to terminate a range loop.
*/
func producersAndConsumers() {

	// This is a producer thread that fills the pipeline
	producer := func(p chan int) {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			p <- rand.Intn(100)
		}
		close(p) // no more production is coming
	}

	pipeline := make(chan int)
	go producer(pipeline)

	// This range will read from the channel until it's closed
	for p := range pipeline {
		fmt.Println(p)
	}
	fmt.Println("Done.\n----------------------------")

	nums := make(chan int)
	stop := make(chan bool)
	go func() { producer(nums) }()
	go func() {
		time.Sleep(2 * time.Second)
		stop <- true // send a command on the stop channel
	}()

	for {
		// This select will read a message from one of the channels and act accordingly.
		// Blocks until one of the cases can run (i.e. there's a message).
		// If there are multiple messages, it chooses one at random.
		select {
		case n := <-nums:
			fmt.Println(n)
		case <-stop:
			fmt.Printf("\nWe're done.\n")
			return
		default: // doesn't block
			fmt.Println("            waiting for data...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func functionsCanBeTypes() {
	type Formatter func(text *string)
	var f Formatter = func(t *string) { *t = fmt.Sprintf(">>>%s<<<", *t) }
	s := "hi"
	f(&s)                       // mutates s
	fmt.Printf("%T %s\n", f, s) // “main.Formatter >>>hi<<<”

	arr := []interface{}{1, "a", 3.14, 0, 0}
	fmt.Println(arr...) // array unpacking
}

// DeferredPrinter demonstrates that 'defer' uses a stack (LIFO)
func DeferredPrinter() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
}

func breakOnLabel() {
FOR:
	for c := 0; c < 100; c++ {
		fmt.Println(c)
		switch {
		case c >= 10:
			break FOR
		default:
			// do work
		}
	}
}

func gobber() {

	type speck struct {
		X int // matches Vertex.X
		Y int // matches Vertex.Y
		S int // doesn't match - will be empty
	}

	var buffer bytes.Buffer        // can be network
	enc := gob.NewEncoder(&buffer) // will encode and send a Vertex
	dec := gob.NewDecoder(&buffer) // wil receive and decode a Speck

	err := enc.Encode(Vertex{1, 2})
	if err != nil {
		fmt.Println(err)
	}

	// This will match Vertex even with the different name
	// because the exported values are the same
	var s speck

	err = dec.Decode(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Speck:", s)
}

func fn(m *map[int]int) {
	*m = make(map[int]int)
	(*m)[3] = 4
}
func fnmain() {
	var m map[int]int
	fn(&m)
	fmt.Println(m)
}

func main() {

	// hello()

	// typesAndReturns()

	// packages()

	// pointers()

	// structs()

	// arrays()

	// maps()

	// calling an anon func:
	// func(s string) { fmt.Printf("Printing from a lambda: %s", s) }("heya")

	// concurrency()
	// producersAndConsumers()

	// functionsCanBeTypes()

	// var e error = fmt.Errorf("Hey, I'm an error!")

	// DeferredPrinter()

	// breakOnLabel()

	// gobber() // packaging and unpackaging data

	// fnmain() // passing maps by reference (explicitly)

	fmt.Println()
}
