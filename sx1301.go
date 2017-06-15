package sx1301

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"golang.org/x/exp/io/spi/driver"

	"github.com/NeuralSpaz/semtech1301/syncrw"
	rpio "github.com/stianeikeland/go-rpio"
	"golang.org/x/exp/io/spi"
)

func Open() (syncrw.SynchronousReadWriter, error) {
	// raspberry pi "/dev/spidev0.0"
	// kerlink "/dev/spidev32766.0"
	// linklabs 	"/dev/spidev0.0"
	// Lorank "/dev/spidev1.0"
	// multitech :S // is it serial

	device := spi.Devfs{
		Dev:      "/dev/spidev0.1",
		Mode:     spi.Mode0,
		MaxSpeed: 10000000,
	}
	deviceConn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}
	//TODO: variadic conf options
	// dev.SetBitOrder(spi.MSBFirst)
	// dev.SetBitsPerWord(8)
	// dev.SetCSChange(false)
	// cs.Output()
	// cs.PullUp()
	//
	// if err != nil {
	// 	return nil, err
	// }
	//TODO: make testable using a supplied buffer struct
	return &SX1301Spi{
		Conn:       deviceConn,
		chipSelect: 0,
	}, nil
}

type SX1301Spi struct {
	page int8
	sync.Mutex
	driver.Conn
	chipSelect rpio.Pin
}

func (s *SX1301Spi) ReadRegister(address byte) (byte, error) {
	rx := make([]byte, 2)
	tx := make([]byte, 2)

	tx[0] = clearBit(address, 7)
	tx[1] = 0x00 // send empty byte for response

	s.Lock()
	// s.chipSelect.Low()
	err := s.Tx(tx, rx)
	// s.chipSelect.High()
	s.Unlock()
	if err != nil {
		return 0, err
	}
	// return second byte.
	return rx[1], nil
}

func (s *SX1301Spi) ReadRegisterByName(sx1301Register string) ([]byte, error) {

	reg, ok := Registers[sx1301Register]
	if !ok {
		return nil, errors.New("error unknown register")
	}
	s.Lock()
	defer s.Unlock()
	if reg.page != s.page {
		fmt.Println("Changing Page")
		// TODO:Change Page
		// Confirm Page has been changed
	}

	address := reg.address
	length := reg.length
	buffersize := (length / 8)

	if address > 127 {
		return nil, errors.New("invalid address range, address greater than 127")
	}

	rx := make([]byte, buffersize+2)
	tx := make([]byte, buffersize+2)

	tx[0] = clearBit(address, 7)
	tx[1] = 0x00 // send empty byte for response

	// s.chipSelect.Low()
	err := s.Tx(tx, rx)
	// s.chipSelect.High()

	return rx[1:], err
}

func (s *SX1301Spi) WriteRegister(address byte, data byte) error {
	rx := make([]byte, 2)
	tx := make([]byte, 2)

	tx[0] = setBit(address, 7) // set write bit
	tx[1] = data

	s.Lock()
	// s.chipSelect.Low()
	err := s.Tx(tx, rx)
	// s.chipSelect.High()
	s.Unlock()

	return err
}

func (s *SX1301Spi) MultiRead(address byte, n uint) ([]byte, error) {
	rx := make([]byte, n)
	tx := make([]byte, n)

	tx[0] = clearBit(address, 7)
	s.Lock()
	s.chipSelect.Low()
	err := s.Tx(tx, rx)
	s.chipSelect.High()
	s.Unlock()

	if err != nil {
		return []byte{}, err
	}
	// return whole buffer, may need to trim first byte.
	return rx, nil
}

func (s *SX1301Spi) MultiWrite(address byte, data []byte) error {
	rx := make([]byte, len(data)+1)
	tx := make([]byte, 1)

	tx[0] = clearBit(address, 7)
	tx = append(tx, data...)

	s.Lock()
	s.chipSelect.Low()
	err := s.Tx(tx, rx)
	s.chipSelect.High()
	s.Unlock()

	if err != nil {
		return err
	}
	return nil
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
