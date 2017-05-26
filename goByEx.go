package main

import (
	"fmt"
	"time"
)

func channelsSelect() {
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

func main() {

}
