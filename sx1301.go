package sx1301

import (
	"errors"
	"fmt"
	"sync"

	"golang.org/x/exp/io/spi/driver"
)

// func New(name string) device {
//
// }

// func Open() (syncrw.SynchronousReadWriter, error) {
// 	// raspberry pi "/dev/spidev0.0"
// 	// kerlink "/dev/spidev32766.0"
// 	// linklabs 	"/dev/spidev0.0"
// 	// Lorank "/dev/spidev1.0"
// 	// multitech :S // is it serial, yes it is -eric@ttn
//
// 	device := spi.Devfs{
// 		Dev:      "/dev/spidev0.1",
// 		Mode:     spi.Mode0,
// 		MaxSpeed: 10000000,
// 	}
// 	deviceConn, err := device.Open()
// 	if err != nil {
// 		log.Println("unable to open device")
// 	}
// 	//TODO: variadic conf options
// 	// dev.SetBitOrder(spi.MSBFirst)
// 	// dev.SetBitsPerWord(8)
// 	// dev.SetCSChange(false)
// 	// cs.Output()
// 	// cs.PullUp()
// 	//
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	//TODO: make testable using a supplied buffer struct
// 	return &SX1301Spi{
// 		Conn: deviceConn,
// 		// ChipSelect: new(Pin),
// 	}, nil
// }

type SX1301Spi struct {
	page int8
	sync.Mutex
	driver.Conn
	ChipSelect   GPIOPin
	csActiveHigh bool
}

func (s *SX1301Spi) chipEnable() {
	if s.csActiveHigh {
		s.ChipSelect.Low()
		return
	}
	s.ChipSelect.High()

}
func (s *SX1301Spi) chipDisable() {
	if s.csActiveHigh {
		s.ChipSelect.High()
		return
	}
	s.ChipSelect.Low()
}

type GPIOPin interface {
	Low()
	High()
	State() uint8 // not really needed TODO: remove from definition
}

func (s *SX1301Spi) ReadRegister(address byte) (byte, error) {
	rx := make([]byte, 2)
	tx := make([]byte, 2)

	tx[0] = clearBit(address, 7)
	tx[1] = 0x00 // send empty byte for response

	s.Lock()
	s.chipEnable()
	err := s.Tx(tx, rx)
	s.chipDisable()
	s.Unlock()
	if err != nil {
		return 0, err
	}
	// return second byte.
	return rx[1], nil
}

func (s *SX1301Spi) WriteRegister(address byte, data byte) error {
	rx := make([]byte, 2)
	tx := make([]byte, 2)

	tx[0] = setBit(address, 7) // set write bit
	tx[1] = data

	s.Lock()
	s.chipEnable()
	err := s.Tx(tx, rx)
	s.chipDisable()
	s.Unlock()

	return err
}

func (s *SX1301Spi) ReadRegisterByName(sx1301Register string) ([]byte, error) {

	reg, ok := Registers[sx1301Register]
	fmt.Printf("reg %+v\n", reg)
	if !ok {
		return nil, errors.New("error unknown register")
	}
	if reg.address > 127 {
		return nil, errors.New("invalid address range, address greater than 127")
	}

	s.Lock()
	defer s.Unlock()
	if reg.address > 32 && reg.page != s.page {
		fmt.Println("Changing Page")
		s.changeRegisterPage(reg.page)
		// TODO: Confirm Page has been changed
	}

	buffersize := (reg.length / 8) + 1 //needs extra byte
	align := reg.length % 8
	if align > 0 {
		buffersize++
	}
	fmt.Println("buffersize ", buffersize)

	rx := make([]byte, buffersize)
	tx := make([]byte, buffersize)

	tx[0] = clearBit(reg.address, 7)
	tx[1] = 0x00 // send empty byte for response

	// s.chipSelect.Low()
	err := s.Tx(tx, rx)
	// s.chipSelect.High()

	return rx[1:], err
}

// only call if you have taken the lock and are within a transaction
func (s *SX1301Spi) changeRegisterPage(pg int8) error {
	fmt.Println("staring page change")
	if pg < 1 || pg > 2 {
		return errors.New("invalid address page")
	}
	rx := make([]byte, 2)
	tx := make([]byte, 2)
	tx[0] = setBit(0x00, 7)
	tx[1] = byte(pg)
	return s.Tx(tx, rx)
}

func (s *SX1301Spi) WriteRegisterByName(sx1301Register string, data ...byte) error { return nil } //TODO: write to the registers by name}

// func (s *SX1301Spi) MultiRead(address byte, n uint) ([]byte, error) {
// 	rx := make([]byte, n)
// 	tx := make([]byte, n)
//
// 	tx[0] = clearBit(address, 7)
// 	s.Lock()
// 	// s.chipSelect.Low()
// 	err := s.Tx(tx, rx)
// 	// s.chipSelect.High()
// 	s.Unlock()
//
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	// return whole buffer, may need to trim first byte.
// 	return rx, nil
// }
//
// func (s *SX1301Spi) MultiWrite(address byte, data []byte) error {
// 	rx := make([]byte, len(data)+1)
// 	tx := make([]byte, 1)
//
// 	tx[0] = clearBit(address, 7)
// 	tx = append(tx, data...)
//
// 	s.Lock()
// 	// s.chipSelect.Low()
// 	err := s.Tx(tx, rx)
// 	// s.chipSelect.High()
// 	s.Unlock()
//
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

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
