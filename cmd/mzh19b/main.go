package main

import (
	"flag"
	"fmt"

	mhz19b "github.com/macaron/go-mh-z19b"
)

func main() {
	calibrate := flag.Bool("r", false, "before calibrate MH-Z19B")
	device := flag.String("device", "/dev/serial0", "specific MH-Z19B")
	flag.Parse()

	mhz19b := mhz19b.New(*device)

	if *calibrate {
		mhz19b.CalibrateDefault()
	}

	ppm, _ := mhz19b.Read()
	fmt.Printf("%d\n", ppm)
}
