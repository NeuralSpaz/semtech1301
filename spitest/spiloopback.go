package spitest

import (
	"fmt"
	"sync"

	"golang.org/x/exp/io/spi/driver"
)

type Device struct{}
type DeviceConn struct {
	sync.Mutex
	cs pin
}

func (s *DeviceConn) Configure(k, v int) error { return nil }
func (s *DeviceConn) Tx(w, r []byte) error {
	fmt.Printf("TX: ")
	for k := range w {
		fmt.Printf("%02x ", w[k])
	}
	copy(r, w)
	fmt.Printf("\nRX: ")
	for k := range w {
		fmt.Printf("%02x ", r[k])
	}
	fmt.Printf("\n")
	return nil
}

func (s *DeviceConn) Close() error { return nil }

func (s *Device) Open() (driver.Conn, error) { return &DeviceConn{}, nil }

type pin uint8

const (
	chipdisable pin = 0
	chipenable  pin = 1
)

func (p *pin) chipEnable() {
	*p = chipenable
}
func (p *pin) chipDisable() {
	*p = chipdisable
}
