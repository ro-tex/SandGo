package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func p(x interface{}) (n int, err error) {
	return fmt.Println(x)
}

func chanSelect() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go func() {
		for {
			c1 <- true
		}
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		c2 <- true
	}()

	for {
		select {
		case <-c1:
			fmt.Print(".")
		case <-c2:
			fmt.Printf("\n\nChannel 2\n")
			return
		default:
			fmt.Print("-")
		}
	}
}

func chanDetectClosed() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "job"
		}
		close(ch)
	}()

	for {
		if s, ok := <-ch; ok {
			fmt.Printf("s: '%v', ok: %v -> reading on\n", s, ok)
		} else { // not OK, so the channel is closed
			fmt.Printf("s: '%v', ok: %v -> closing\n", s, ok)
			return
		}
	}
}

// We can read from a closed channel and get the values stored there, as long as it's buffered.
// If it's not or if we run out of values we'll get the default channel value.
func chanReadFromClosed() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	n, ok := <-ch
	fmt.Println(n, ok)
	n, ok = <-ch
	fmt.Println(n, ok)
	n, ok = <-ch
	fmt.Println(n, ok)
}

func chanOperations() {
	// Functions are first class citizens, so they can be passed via channels.
	// This cahnnel will deliver the operation to be performed on the payload "s":
	chOps := make(chan func(string) int, 10)

	s := "asdf"

	go func() {
		for f := range chOps {
			fmt.Println(f(s))
		}
	}()

	chOps <- func(str string) int {
		return len(str)
	}

	chOps <- func(str string) int {
		return 2600 // something...
	}
}

func chanWorkerPool() {
	type payload struct {
		DataField string
	}

	// always buffer your channels unless you need to use them for synchro
	chWork := make(chan payload, 10)

	// spawning 5 workers
	for i := 0; i < 5; i++ {
		go func(workerID int) {
			for work := range chWork {
				time.Sleep(10 * time.Millisecond) // doing work
				log.Println("Payload processed by worker #" + strconv.Itoa(workerID) + "\t" + work.DataField)
			}
		}(i)
	}
	log.Println("workers are now running")

	// send 10 units of work
	for i := 0; i < 10; i++ {
		chWork <- payload{"work unit #" + strconv.Itoa(i)}
	}
	log.Println("closing channel - no more work")
	// A key point here is that when we close the channel that will terminate the range operations in the workers
	// and they will terminate. Thus, even if the function goes on to do something else that memory will be freed.
	// TL;DR: Close your channels when you're done with them!
	close(chWork)
	log.Println("done.")
}

type kid struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// set is a mutator, so it takes a pointer
func (k *kid) set(name string, age int) {
	k.Name = name
	k.Age = age
}

// not a mutator - doesn't need a pointer
// NOTE: it's good style to have all pointers or all non-pointers!
func (k kid) getAge() int {
	return k.Age
}

type kids []kid

// Implement the sort interface:
func (k kids) Len() int {
	return len(k)
}
func (k kids) Less(i, j int) bool {
	return k[i].getAge() < k[j].getAge()
}
func (k kids) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func sortingCustom() {
	group := kids{kid{"derek", 5}, kid{"jessica", 4}, kid{"mel", 3}, kid{"pete", 4}}
	sort.Sort(group)
	fmt.Println(group)
}

func jsons() {
	b := []byte(`{"a":123,"b":"asdf","c":["c1","c2"],"d":{"d1":1,"d2":"dibi"}}`)
	var j map[string]interface{} // this container can take any JSON

	if err := json.Unmarshal(b, &j); err != nil {
		panic(err)
	}
	// in order to use the parsed JSON value we need to cast it.
	// including if we want to use the indexing function of the map.
	// the good news is that we can cast in-line. clumsy but works.
	fmt.Println(j["d"].(map[string]interface{})["d2"].(string))

	// we'll unmarshall a specific type - 'kid'
	joeBytes := []byte(`{"name":"Joe","age":13,"extra":"won't be unmarshalled"}`)
	var joe kid
	if err := json.Unmarshal(joeBytes, &joe); err != nil {
		panic(err)
	}
	fmt.Println(joe) // the "extra" won't be there
	// accessing fields and calling methods
	// no need to type-cast!
	fmt.Printf("Name: %s, Age: %d", joe.Name, joe.getAge())
}

func timePlay() {
	p(time.Now().Format(time.RFC3339))

	t, err := time.Parse(time.RFC3339, "2017-05-29T16:31:33+02:00")
	if err != nil {
		panic(err)
	}
	p(t)
}

func main() {
	timePlay()
}
