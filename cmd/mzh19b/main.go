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

	m := mhz19b.New(*device)

	if *calibrate {
		_ = m.CalibrateDefault()
	}

	ppm, _ := m.Read()
	fmt.Printf("%d\n", ppm)
}
