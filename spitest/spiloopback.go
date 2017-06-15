package spitest

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"sync"

	"golang.org/x/exp/io/spi/driver"
)

type Device struct {
	MemoryMap bool
	MM        map[byte]byte           // address mapped
	PM        map[int8]map[byte]int32 // page address mapped
}

type DeviceConn struct {
	Txbuf []byte
	Rxbuf []byte
	sync.Mutex
	cs pin
}

type MMDeviceConn struct {
	MM    map[byte]byte
	Txbuf []byte
	Rxbuf []byte
	sync.Mutex
	cs pin
}

type PMDeviceConn struct {
	PM    map[int8]map[byte]int32
	Txbuf []byte
	Rxbuf []byte
	sync.Mutex
	cs pin
}

func (s *DeviceConn) Configure(k, v int) error { return nil }
func (s *DeviceConn) Tx(w, r []byte) error {
	fmt.Printf("TX: ")
	for k := range w {
		fmt.Printf("%02x ", w[k])
	}
	copy(s.Txbuf, w)
	copy(s.Rxbuf, r)
	// s.Txbuf := make([]byte, len(w))
	// s.Rxbuf := make([]byte, len(r))
	copy(r, w)
	fmt.Printf("\nRX: ")
	for k := range w {
		fmt.Printf("%02x ", r[k])
	}
	fmt.Printf("\n")
	return nil
}

func (s *DeviceConn) Close() error { return nil }

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
	address := clearBit(w[0], 7)
	value, ok := s.MM[address]
	if !ok {
		return errors.New("no value mapped to register")
	}
	if hasBit(w[0], 7) {
		s.MM[address] = w[1]
	} else {
		r[1] = value
	}
	return nil
}

func (s *MMDeviceConn) Close() error { return nil }

func (s *PMDeviceConn) Configure(k, v int) error { return nil }
func (s *PMDeviceConn) Tx(w, r []byte) error {
	log.Println("Length of tx", len(w))
	// get current page
	pg := int8(s.PM[-1][0x00] & 0x03)

	fmt.Printf("current page: %d\n", pg)
	address := clearBit(w[0], 7)
	if address < 33 {
		pg = -1 //all page -1 registers are all 32 or less
	}
	if hasBit(w[0], 7) { // WriteRegister
		fmt.Println("Write request")
		if address == 0x00 { // change page
			fmt.Println("page change request")
			s.PM[-1][0x00] = int32(w[1] & 0x03) // force int32
		}
	} else { // ReadRegister
		value, ok := s.PM[pg][address]
		if !ok {
			return errors.New("no value mapped to register")
		}
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(value))
		fmt.Printf("%08x\n", buf)
		r[0] = 0x00 // always garbage
		copy(r[1:], buf[len(buf)-(len(r)-1):])

	}

	return nil
}

func (s *PMDeviceConn) Close() error { return nil }

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
