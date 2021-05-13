package main

import (
	"flag"
	"fmt"
	"github.com/macaron/go-mh-z19b"
)

var revision string

func main() {
	device := flag.String("device", "/dev/serial0", "specific MH-Z19B")
	rev := flag.Bool("revision", false, "show Revision")
	flag.Parse()

	if *rev {
		fmt.Println("Revision:", revision)
		return
	}

	ppm, _ := mhz19b.Read(*device)
	fmt.Printf("%d\n", ppm)
}
