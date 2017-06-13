package sx1301

import (
	"sync"

	rpio "github.com/stianeikeland/go-rpio"

	"golang.org/x/exp/io/spi"
)

type SynchronousReadWriter interface {
	ReadRegister(addr byte) (byte, error)
	WriteRegister(addr byte, data byte) error
	MultiRead(addr byte, n uint) ([]byte, error)
	MultiWrite(addr byte, data []byte) error
}

func Open(path string, cs rpio.Pin) (SynchronousReadWriter, error) {
	dev, err := spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev0.1",
		Mode:     spi.Mode0,
		MaxSpeed: 10000000,
	})
	//TODO: variadic conf options
	dev.SetBitOrder(spi.MSBFirst)
	dev.SetBitsPerWord(8)
	dev.SetCSChange(false)
	cs.Output()
	cs.PullUp()

	if err != nil {
		return nil, err
	}
	//TODO: make testable using a supplied buffer struct
	return &sx1301DirectSpi{
		device:     dev,
		chipSelect: cs,
	}, nil
}

type sx1301DirectSpi struct {
	sync.Mutex
	device     *spi.Device
	chipSelect rpio.Pin
}

func (s *sx1301DirectSpi) ReadRegister(address byte) (byte, error) {
	rx := make([]byte, 2)
	tx := make([]byte, 2)

	tx[0] = clearBit(address, 7)
	tx[1] = 0x00 // send empty byte for response

	s.Lock()
	s.chipSelect.Low()
	err := s.device.Tx(tx, rx)
	s.chipSelect.High()
	s.Unlock()
	if err != nil {
		return 0, err
	}
	// return second byte.
	return rx[1], nil
}

func (s *sx1301DirectSpi) WriteRegister(address byte, data byte) error {
	rx := make([]byte, 2)
	tx := make([]byte, 2)

	tx[0] = setBit(address, 7) // set write bit
	tx[1] = data

	s.Lock()
	s.chipSelect.Low()
	err := s.device.Tx(tx, rx)
	s.chipSelect.High()
	s.Unlock()

	if err != nil {
		return err
	}

	return err
}

func (s *sx1301DirectSpi) MultiRead(address byte, n uint) ([]byte, error) {
	rx := make([]byte, n)
	tx := make([]byte, n)

	tx[0] = clearBit(address, 7)
	s.Lock()
	s.chipSelect.Low()
	err := s.device.Tx(tx, rx)
	s.chipSelect.High()
	s.Unlock()

	if err != nil {
		return []byte{}, err
	}
	// return whole buffer, may need to trim first byte.
	return rx, nil
}

func (s *sx1301DirectSpi) MultiWrite(address byte, data []byte) error {
	rx := make([]byte, len(data)+1)
	tx := make([]byte, 1)

	tx[0] = clearBit(address, 7)
	tx = append(tx, data...)

	s.Lock()
	s.chipSelect.Low()
	err := s.device.Tx(tx, rx)
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
