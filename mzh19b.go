package mhz19b

import (
	"errors"
	"time"

	"github.com/tarm/serial"
)

type MHZ19B struct {
	SerialConfig *serial.Config
}

func New(device string) *MHZ19B {
	serialConfig := &serial.Config{Name: device, Baud: 9600, Size: 8, Parity: 0, StopBits: 1}
	mhz19b := &MHZ19B{SerialConfig: serialConfig}
	return mhz19b
}

func (mhz19b *MHZ19B) CalibrateDefault() error {
	port, err := serial.OpenPort(mhz19b.SerialConfig)
	if err != nil {
		return err
	}
	defer port.Close()

	// Reset default point (400ppm)
	_, err = port.Write([]byte{0xFF, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x78})
	if err != nil {
		return err
	}

	time.Sleep(time.Second)

	return nil
}

func (mhz19b *MHZ19B) Read() (int, error) {
	port, err := serial.OpenPort(mhz19b.SerialConfig)
	if err != nil {
		return 0, err
	}
	defer port.Close()

	// Request CO2 concentration
	_, err = port.Write([]byte{0xFF, 0x01, 0x86, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79})
	if err != nil {
		return 0, err
	}

	time.Sleep(time.Second)

	// Response CO2 concentration
	buf := make([]byte, 9)
	_, err = port.Read(buf)
	if err != nil {
		return 0, err
	}

	crc := getCheckSum(buf)
	if buf[8] != crc {
		return 0, errors.New("CRC error")
	}

	return int(buf[2])*256 + int(buf[3]), nil
}

func Read(device string) (ppm int, err error) {
	mhz19b := New(device)
	return mhz19b.Read()
}

func getCheckSum(packet []byte) byte {
	var checksum byte = 0
	for i := 1; i < 8; i++ {
		checksum += packet[i]
	}
	checksum = 255 - checksum
	checksum += 1
	return checksum
}
