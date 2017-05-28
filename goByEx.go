package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

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

func main() {

}
