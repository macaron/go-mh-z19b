package main

import (
	"flag"
	"fmt"

	mhz19b "github.com/macaron/go-mh-z19b"
)

var revision string

func main() {
	calibrate := flag.Bool("r", false, "before calibrate MH-Z19B")
	device := flag.String("device", "/dev/serial0", "specific MH-Z19B")
	rev := flag.Bool("revision", false, "show Revision")
	flag.Parse()

	m := mhz19b.New(*device)

	if *rev {
		fmt.Println("Revision:", revision)
		return
	}

	if *calibrate {
		_ = m.CalibrateDefault()
	}

	ppm, _ := mhz19b.Read(*device)
	fmt.Printf("%d\n", ppm)
}
