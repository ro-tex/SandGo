package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/ro-tex/SandGo/hello"
)

func helloWorld() {
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
	// import and use a local package:
	fmt.Println(hello.Hello())
}

func pointers() {
	f := 22 / 7.0
	var p *float64 = &f
	fmt.Println(*p)

	*p = 2.26 // dereferencing or indirecting
	fmt.Println(f)

	n := 42
	pn := &n
	ppn := &pn
	fmt.Printf("var(n): %d, pointer to var(*pn): %d, pointer to pointer to var(**ppn): %d\n", n, *pn, **ppn)

	v := Vertex{2, 3}

	fmt.Println(v)
	v.Rotate()
	fmt.Println(v)

	fmt.Println("Objects:")

	var o interface{}
	fmt.Println(o) // nil
	o = v
	fmt.Println(o) // {4, 2}
	o = &v
	fmt.Println(o) // &{4, 2}

	vvv, err := o.(Vertex) // this is basically o.typeof(vertex)

	if !err {
		fmt.Println("It's a vertex:", vvv)
	} else {
		fmt.Println("Not a vertex")
	}
}

type Vertex struct {
	X int `json:"a"` // needs to be public in order to be exported, as in "X int"
	Y int `json:"b"`
}

// Implement the Stringer interface
func (v Vertex) String() string {
	return fmt.Sprintf("{X: %d, Y: %d}", v.X, v.Y)
}

/*
	So this is basically what an object method looks like in Go. You can only define methods on types defined
	in the current package (so no built-in types).

	The v that we pass in the method signature defines how we see 'this' in the scope of the method - it's
	up to us to name it.
*/
func (v Vertex) Quadrant() int {
	switch {
	case v.X >= 0 && v.Y >= 0:
		return 1
	case v.X < 0 && v.Y >= 0:
		return 2
	case v.X < 0 && v.Y < 0:
		return 3
	default:
		return 4
	}
}

// Mutator: This one takes a pointer to an object, so it can modify it.
func (v *Vertex) Rotate() {
	v.X, v.Y = v.Y, v.X
}

// ### DEFINING ERRORS START ###
type SimpleFloatError float64

func (e SimpleFloatError) Error() string {
	return fmt.Sprintf("negative input: %f", float64(e))
}

type FloatErrorWithMessage struct {
	val float64
	msg string
}

func (e FloatErrorWithMessage) Error() string {
	return fmt.Sprintf("%s value: %f", e.msg, e.val)
}

func TestError(val float64) (float64, error) {
	// return 0, errors.New(fmt.Sprintf("bad input: %f", val)) // the built-in way
	//return 0, SimpleFloatError(val) // simple custom error
	return 0, FloatErrorWithMessage{val, "Can't work with this!"} // more functional custom error
}

// ### DEFINING ERRORS END ###

// returns an added function that is a closure
func adder(base int) func(int) int {
	sum := base

	return func(x int) int {
		sum := sum + x
		return sum
	}
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

func interfaces(v Vertex) {
	var s fmt.Stringer

	s = &v

	stringer, ok := s.(error)
	if !ok {
		fmt.Printf("It can be a string! Value: %v, Type: %T, ok: %v\n", stringer, stringer, ok)
	} else {
		fmt.Printf("It CANNOT be a string! Value: %v, Type: %T, ok: %v\n", stringer, stringer, ok)
	}

	a := []interface{}{"1", "2", "3", 7}
	a = append([]interface{}{"dump"}, a...)
	fmt.Println(a...)
}

func typeSwitch(o interface{}) {
	// A type switch enables us to react to several possible incoming parameter types in a single structure.
	switch obj := o.(type) { // type is the keyword
	case int:
		fmt.Println("int", obj)
	case Vertex:
		fmt.Println("It's a vertex!", obj)
	default:
		fmt.Printf("We cannot process this thing! Type: %T\n", obj)
	}
}

// This is a function that gets any number of arguments of a given type. In this case - any type.
func variadic(args ...interface{}) {
	for _, v := range args {
		fmt.Print(v, " | ")
	}
}

func readers() {
	var stream string = "Hello, this is a stream of data to be read."
	r := strings.NewReader(stream)
	b := make([]byte, 8) // we'll read into this buffer

	for {
		n, err := r.Read(b) // n is number of bytes read

		/*
			What is VERY important here is the [:n] slicing when we are reading the buffer.
			The problem is that when the last chunk is read there will be garbage at the end of the buffer
			and we want to filter out that garbage by taking just the number of bytes we just read.
		*/
		fmt.Printf("n = %v err = %v b = %v s = %s\n", n, err, b[:n], string(b[:n]))

		if err == io.EOF {
			break
		}
	}
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

func writeToFile(text string, fileName string) {
	f, e := os.Create(fileName)

	// The file handle will be closed when we exit this function.
	// This is a good approach, so we don't forget to do it.
	defer f.Close()

	if e != nil {
		fmt.Println(e)
		return
	}
	f.Write([]byte(text))
}

// https://golang.org/pkg/runtime/#Caller
func getCallerFileAndLine() {
	// Caller param - how many levels of the stack trace to skip
	programCounter, fileName, lineNumber, ok := runtime.Caller(1)

	fmt.Printf("%+v %+v %+v %+v ", programCounter, fileName[strings.LastIndex(fileName, "/")+1:], lineNumber, ok)
	return
}

type tree struct {
	v int
	l *tree
	r *tree
}

func (t *tree) find(x int) bool {
	if t == nil {
		return false
	}
	return t.v == x || t.l.find(x) || t.r.find(x)
}

func EquableTriangle(a, b, c int) bool {
	p := float64(a + b + c)
	return p == math.Sqrt(p*(p-float64(2*a))*(p-float64(2*b))*(p-float64(2*c)))/4
}

func Accum(s string) string {
	segments := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		segments[i] = strings.ToUpper(string(s[i])) + strings.Repeat(string(s[i]), i)
	}
	return strings.Join(segments, "-")
}

func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	words = strings.ToUpper(strings.Replace(words, " ", "", -1))
	out := make([]string, 0, len(words))
	idxA, idxZ := int('A'), int('Z')
	for i := 0; i < len(words); i++ {
		if idxA <= int(words[i]) && int(words[i]) <= idxZ {
			idx := int(words[i]) - idxA
			out = append(out, nato[idx])
		} else if string(words[i]) == " " {
			continue
		} else {
			out = append(out, string(words[i]))
		}
	}
	return strings.Join(out, " ")
}

// ProductFib solves https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go
func ProductFib(prod uint64) [3]uint64 {
	fib := getFibonacciFunc()
	for i := uint64(0); ; i++ {
		f1, f2 := fib(i), fib(i+1)
		if f1*f2 == prod {
			return [3]uint64{f1, f2, 1}
		} else if f1*f2 > prod {
			return [3]uint64{f1, f2, 0}
		}
	}
}

// Returns a function that returns the Fibonacci numbers while utilising dynamic programming
// in order to optimise calculation. We use a closure in order to make the fibs slice permanent.
func getFibonacciFunc() func(uint64) uint64 {
	// we'll cache all Fibonacci numbers we calculate in memory:
	fibs := []uint64{0, 1, 1}

	return func(n uint64) uint64 {
		if n >= uint64(len(fibs)) {
			// we don't have that number calculated - we'll now calculate all
			// Fibonacci numbers leading up to (and incl.) the number:
			for i := uint64(len(fibs)); i <= n; i++ {
				fibs = append(fibs, fibs[i-2]+fibs[i-1])
			}
		}
		// we have that already calculated - we'll just return it:
		return fibs[n]
	}
}

func main() {

	p := fmt.Println

	// helloWorld()

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

	// writeToFile("hello", "hello.txt")

	// getCallerFileAndLine()

	// f := getFibonacciFunc()
	// p(f(3), f(5), f(9))

	p()
}
