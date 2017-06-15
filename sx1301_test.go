package sx1301

import (
	"log"
	"testing"

	"github.com/NeuralSpaz/semtech1301/spitest"
)

// type SynchronousReadWriter interface {
// 	ReadRegister(addr byte) (byte, error)
// 	WriteRegister(addr byte, data byte) error
// 	MultiRead(addr byte, n uint) ([]byte, error)
// 	MultiWrite(addr byte, data []byte) error
// }

func NewSynchronousReadWriter() {
	// return SynchronousReadWriter

	return
}

func TestReadRegister(t *testing.T) {
	return
}

func TestWriteRegister(t *testing.T) {

}

func TestClearbit(t *testing.T) {
	var bitTests = []struct {
		n    byte
		pos  uint8
		want byte
	}{
		{0xFF, 7, 0x7F},
		{0xFF, 6, 0xBF},
		{0xFF, 5, 0xDF},
		{0xFF, 4, 0xEF},
		{0xFF, 3, 0xF7},
		{0xFF, 2, 0xFB},
		{0xFF, 1, 0xFD},
		{0xFF, 0, 0xFE},
	}
	for _, test := range bitTests {
		got := clearBit(test.n, test.pos)
		if got != test.want {
			t.Errorf("want %02x but got %02x for %02x", test.want, got, test.n)
		}
	}
}

func TestSetbit(t *testing.T) {
	var bitTests = []struct {
		n    byte
		pos  uint8
		want byte
	}{
		{0x7F, 7, 0xFF},
		{0xBF, 6, 0xFF},
		{0xDF, 5, 0xFF},
		{0xEF, 4, 0xFF},
		{0xF7, 3, 0xFF},
		{0xFB, 2, 0xFF},
		{0xFD, 1, 0xFF},
		{0xFE, 0, 0xFF},
	}
	for _, test := range bitTests {
		got := setBit(test.n, test.pos)
		if got != test.want {
			t.Errorf("want %02x but got %02x for %02x", test.want, got, test.n)
		}
	}
}

func TestHasbit(t *testing.T) {
	var bitTests = []struct {
		n    byte
		pos  uint8
		want bool
	}{
		{0x7F, 7, false},
		{0xBF, 6, false},
		{0xDF, 5, false},
		{0xEF, 4, false},
		{0xF7, 3, false},
		{0xFB, 2, false},
		{0xFD, 1, false},
		{0xFE, 0, false},
		{0x80, 7, true},
		{0x40, 6, true},
		{0x20, 5, true},
		{0x10, 4, true},
		{0x08, 3, true},
		{0x04, 2, true},
		{0x02, 1, true},
		{0x01, 0, true},
	}
	for _, test := range bitTests {
		got := hasBit(test.n, test.pos)
		if got != test.want {
			t.Errorf("want %t but got %t for %02x", test.want, got, test.n)
		}
	}
}

func TestWriteRegisterWithLoopBack(t *testing.T) {
	device := spitest.Device{}
	deviceConn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}

	loopback := &SX1301Spi{
		Conn:       deviceConn,
		chipSelect: 0,
	}

	_, err = loopback.ReadRegister(0x01)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	err = loopback.WriteRegister(0x01, 0x05)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

func TestWriteRegisterWithLoopBackMM(t *testing.T) {

	memorymap := make(map[byte]byte)
	memorymap[0x01] = 0xAA
	device := spitest.Device{MemoryMap: true, MM: memorymap}

	conn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}
	log.Printf("%+#v", conn)

	loopback := &SX1301Spi{
		Conn:       conn,
		chipSelect: 0,
	}

	log.Printf("%+#v\n", loopback)

	value, err := loopback.ReadRegister(0x01)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("%02x\n", value)

}

func TestWriteRegisterByNameWithLoopBackMM(t *testing.T) {

	memorymap := make(map[byte]byte)
	memorymap[0x01] = 0xAA
	memorymap[0x21] = 0xBB
	device := spitest.Device{MemoryMap: true, MM: memorymap}

	conn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}
	log.Printf("%+#v", conn)

	loopback := &SX1301Spi{
		Conn:       conn,
		chipSelect: 0,
	}

	log.Printf("%+#v\n", loopback)

	value, err := loopback.ReadRegisterByName("LGW_RX_INVERT_IQ")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("%02x\n", value)
	//
	// value, err = loopback.ReadRegisterByName("LGW_RX_DATA_BUF_ADDR")
	// if err != nil {
	// 	t.Errorf("expected no error but got: %v", err)
	// }
	// log.Printf("%02x\n", value)

}

func TestBuildPageMap(t *testing.T) {
	paged := make(map[byte]int32)
	page0 := make(map[byte]int32)
	page1 := make(map[byte]int32)
	page2 := make(map[byte]int32)
	page3 := make(map[byte]int32)
	pagemap := make(map[int8]map[byte]int32)
	pagemap[-1] = paged
	pagemap[0] = page0
	pagemap[1] = page1
	pagemap[2] = page2
	pagemap[3] = page3
	log.Printf("%+#v\n", pagemap)
	for _, v := range Registers {
		pagemap[v.page][v.address] = v.defaultValue
	}

	device := spitest.Device{MemoryMap: true, PM: pagemap}

	conn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}
	log.Printf("%+#v", conn)

	loopback := &SX1301Spi{
		Conn:       conn,
		chipSelect: 0,
	}

	log.Printf("%+#v\n", loopback)

	value, err := loopback.ReadRegisterByName("LGW_ADJUST_MODEM_START_OFFSET_SF12_RDX4")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("LGW_ADJUST_MODEM_START_OFFSET_SF12_RDX4: %d\n", value)

	value, err = loopback.ReadRegisterByName("LGW_IF_FREQ_0")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("LGW_IF_FREQ_0: %d\n", value)

	value, err = loopback.ReadRegisterByName("LGW_TX_STATUS")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("LGW_TX_STATUS: %d\n", value)

	value, err = loopback.ReadRegisterByName("LGW_TIMESTAMP")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("LGW_TIMESTAMP: %d\n", value)

	value, err = loopback.ReadRegisterByName("LGW_VERSION")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	log.Printf("LGW_VERSION: %d\n", value)
}
