package mhz19b

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

func Read(device string) (ppm int, err error) {
	c := &serial.Config{Name: device, Baud: 9600, Size: 8, Parity: 0, StopBits: 1}
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("Device error: %v", err)
		return 0, err
	}
	defer port.Close()

	// Request CO2 concentration
	_, err = port.Write([]byte{0xFF, 0x01, 0x86, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79})
	if err != nil {
		log.Printf("Write error: %v", err)
		return 0, err
	}

	time.Sleep(time.Second)

	// Response CO2 concentration
	buf := make([]byte, 9)
	_, err = port.Read(buf)
	if err != nil {
		log.Printf("Read error: %v", err)
		return 0, err
	}

	crc := getCheckSum(buf)
	if buf[8] != crc {
		log.Printf("CRC error: %d=%d", buf[8], crc)
		return 0, err
	}

	return int(buf[2])*256 + int(buf[3]), nil
}

func getCheckSum(packet []byte) byte {
	var checksum byte = 0
	for i := 1; i < 7; i++ {
		checksum += packet[i]
	}
	checksum = 255 - checksum
	checksum += 1
	return checksum
}
