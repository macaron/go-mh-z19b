package main

import (
	"flag"
	"fmt"
	"github.com/macaron/go-mh-z19b"
)

func main() {
	device := flag.String("device", "/dev/serial0", "specific MH-Z19B")
	flag.Parse()

	ppm, _ := mhz19b.Read(*device)
	fmt.Printf("%d\n", ppm)
}
