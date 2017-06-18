package spitest

import (
	"errors"
	"sync"

	"golang.org/x/exp/io/spi/driver"
)

type Device struct {
	MemoryMap bool
	MM        map[byte]byte          // address mapped
	PM        map[byte]map[byte]byte // page address mapped
	// Pin
}

type DeviceConn struct {
	Txbuf []byte
	Rxbuf []byte
	sync.Mutex
	// Pin
}

type MMDeviceConn struct {
	MM    map[byte]byte
	Txbuf []byte
	Rxbuf []byte
	sync.Mutex
	// Pin
}

type PMDeviceConn struct {
	PM    map[byte]map[byte]byte
	Txbuf []byte
	Rxbuf []byte
	sync.Mutex
	// Pin
}

var (
	UnequalBufferError  error = errors.New("sync read writer buffer must be same size")
	MemoryMapValueError error = errors.New("no value mapped to register")
	PagePointerError    error = errors.New("page pointer not initialised")
)

func (s *DeviceConn) Configure(k, v int) error { return nil }

func (s *DeviceConn) Tx(w, r []byte) error {
	if len(w) != len(r) {
		return UnequalBufferError
	}
	copy(r, w)
	return nil
}

func (s *DeviceConn) Close() error { return nil } //TODO

func (d *Device) Open() (driver.Conn, error) {
	if d.MM != nil {
		return &MMDeviceConn{MM: d.MM}, nil
	}
	if d.PM != nil {
		return &PMDeviceConn{PM: d.PM}, nil
	}
	return &DeviceConn{}, nil
}

func (s *MMDeviceConn) Configure(k, v int) error { return nil }
func (s *MMDeviceConn) Tx(w, r []byte) error {
	if len(w) != len(r) {
		return UnequalBufferError
	}
	address := clearBit(w[0], 7)
	value, ok := s.MM[address]
	if !ok {
		return MemoryMapValueError
	}
	if hasBit(w[0], 7) {
		s.MM[address] = w[1]
	} else {
		r[1] = value
	}
	return nil
}

func (s *MMDeviceConn) Close() error { return nil } //TODO

func (s *PMDeviceConn) Configure(k, v int) error { return nil }
func (s *PMDeviceConn) Tx(w, r []byte) error {
	if len(w) != len(r) {
		return UnequalBufferError
	}
	// log.Println("Length of tx", len(w))
	// get current page
	page, ok := s.PM[0x00][0x00] // check the current page pointer
	if !ok {
		return PagePointerError
	}
	pg := page & 0x03 // apply page mask

	// fmt.Printf("current page: %d\n", pg)

	address := clearBit(w[0], 7)
	if address < 33 {
		pg = 0x00 //all page -1 registers are all 32 or less
	}
	if hasBit(w[0], 7) { // WriteRegister
		// fmt.Println("Write request")
		if address == 0x00 { // change page
			// fmt.Println("page change request")
			s.PM[0x00][0x00] = w[1] & 0x03 // force int32
		} else {
			s.PM[pg][address] = w[1]
		}
	} else { // ReadRegister
		for i := 1; i < len(w); i++ {
			r[i] = s.PM[pg][w[i-1]]
		}
	}

	return nil
}

func (s *PMDeviceConn) Close() error { return nil } //TODO

type Pin uint8

const (
	chipdisable Pin = 0
	chipenable  Pin = 1
)

func (p *Pin) Low() {
	*p = 0
}
func (p *Pin) High() {
	*p = 1
}

func (p *Pin) State() uint8 {
	state := *p
	return uint8(state)
}

func clearBit(n byte, pos uint8) byte {
	mask := ^(1 << pos)
	n &= byte(mask)
	return n
}

func setBit(n byte, pos uint8) byte {
	n |= (1 << pos)
	return n
}

func hasBit(n byte, pos uint8) bool {
	val := n & (1 << pos)
	return (val > 0)
}
