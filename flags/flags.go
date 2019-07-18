package main

import (
	"flag"
	"fmt"
)

var flagvar int

func init() {
	// this binds the flag to an existing variable:
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}

func main() {
	var ip = flag.Int("ip", 1234, "help message for ip")
	flag.Parse()
	fmt.Println("ip is", *ip)
	fmt.Println("flagname is", flagvar)
}
