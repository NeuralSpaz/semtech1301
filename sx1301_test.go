package sx1301

import (
	"encoding/binary"
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

// func NewSynchronousReadWriter() {
// 	// return SynchronousReadWriter
//
// 	return
// }
//
// func TestReadRegister(t *testing.T) {
// 	return
// }
//
// func TestWriteRegister(t *testing.T) {
//
// }

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
		ChipSelect: new(spitest.Pin),
	}

	err = loopback.WriteRegister(0x01, 0x05)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

func TestReadRegisterWithLoopBack(t *testing.T) {

	device := spitest.Device{}
	deviceConn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}

	loopback := &SX1301Spi{
		Conn:       deviceConn,
		ChipSelect: new(spitest.Pin),
	}
	var data byte = 0x01

	register, err := loopback.ReadRegister(data)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	if register != 0x00 {
		t.Errorf("expected register to be 0x00 but got : %v", register)
	}
}

func TestChipSelectActiveLow(t *testing.T) {

	device := spitest.Device{}
	deviceConn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}

	loopback := &SX1301Spi{
		Conn:         deviceConn,
		ChipSelect:   new(spitest.Pin),
		csActiveHigh: false,
	}
	loopback.Lock()
	loopback.chipEnable()

	if loopback.ChipSelect.State() != 1 {
		t.Errorf("expected chip to be enabled, %v", loopback.ChipSelect.State())
	}

	loopback.Unlock()

	loopback.Lock()
	loopback.chipDisable()

	if loopback.ChipSelect.State() != 0 {
		t.Errorf("expected chip to be disabled, %v", loopback.ChipSelect.State())
	}

	loopback.Unlock()
}

func TestChipSelectActiveHigh(t *testing.T) {

	device := spitest.Device{}
	deviceConn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}

	loopback := &SX1301Spi{
		Conn:         deviceConn,
		ChipSelect:   new(spitest.Pin),
		csActiveHigh: true,
	}
	loopback.Lock()
	loopback.chipEnable()

	if loopback.ChipSelect.State() != 0 {
		t.Errorf("expected chip to be enabled, %v", loopback.ChipSelect.State())
	}

	loopback.Unlock()

	loopback.Lock()
	loopback.chipDisable()
	if loopback.ChipSelect.State() != 1 {
		t.Errorf("expected chip to be disabled, %v", loopback.ChipSelect.State())
	}

	loopback.Unlock()
}

func TestReadRegisterWithLoopBackMM(t *testing.T) {

	memorymap := make(map[byte]byte)
	memorymap[0x01] = 0xAA
	device := spitest.Device{MemoryMap: true, MM: memorymap}

	conn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}

	loopback := &SX1301Spi{
		Conn:       conn,
		ChipSelect: new(spitest.Pin),
	}

	_, err = loopback.ReadRegister(0x01)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

}

func TestReadRegisterByNameWithLoopBackMM(t *testing.T) {

	memorymap := make(map[byte]byte)
	memorymap[0x01] = 0xAA
	memorymap[0x21] = 0xBB
	device := spitest.Device{MemoryMap: true, MM: memorymap}

	conn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}
	// log.Printf("%+#v", conn)

	loopback := &SX1301Spi{
		Conn:       conn,
		ChipSelect: new(spitest.Pin),
	}

	_, err = loopback.ReadRegisterByName("LGW_RX_INVERT_IQ")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	//
	// value, err = loopback.ReadRegisterByName("LGW_RX_DATA_BUF_ADDR")
	// if err != nil {
	// 	t.Errorf("expected no error but got: %v", err)
	// }
	// log.Printf("%02x\n", value)

}

// helper
func initPageMapForsx1301() map[byte]map[byte]byte {
	// paged := make(map[byte]byte)
	page0 := make(map[byte]byte)
	page1 := make(map[byte]byte)
	page2 := make(map[byte]byte)
	pagemap := make(map[byte]map[byte]byte)
	// pagemap[-1] = paged
	pagemap[0] = page0
	pagemap[1] = page1
	pagemap[2] = page2

	//TODO:
	// need to be actual regiser map combining what actual register would
	// return not just the default value ie map[page]map[address]data so
	// that read register and write register and by name are written correctly
	// and tested before actual hardware is used.
	// Capture from logic analiser should match testing buffer.
	for _, v := range Registers {
		data := register2ByteSlice(int64(v.defaultValue), v.length, v.offset, v.signed)
		var page byte
		if v.page == -1 {
			page = 0
		} else {
			page = byte(v.page)
		}
		addr := v.address
		for i := range data {
			pagemap[page][addr+byte(i)] = data[i] ^ pagemap[page][addr+byte(i)]
		}
	}

	return pagemap
}

// helper
func newPageMappedSX1301() *SX1301Spi {
	pagemap := initPageMapForsx1301()

	device := spitest.Device{MemoryMap: true, PM: pagemap}

	conn, err := device.Open()
	if err != nil {
		log.Println("unable to open device")
	}

	loopback := &SX1301Spi{
		Conn:       conn,
		ChipSelect: new(spitest.Pin),
	}
	return loopback
}

func TestBuildPageMap(t *testing.T) {

	loopback := newPageMappedSX1301()

	value, err := loopback.ReadRegisterByName("LGW_ADJUST_MODEM_START_OFFSET_SF12_RDX4")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	registerDefault := Registers["LGW_ADJUST_MODEM_START_OFFSET_SF12_RDX4"].defaultValue
	want := make([]byte, 4)
	binary.BigEndian.PutUint32(want, uint32(registerDefault))
	for k := range value {
		if value[k] != want[k+(4-len(value))] {
			t.Errorf("expected the same value as default %v but got %v", want, value)
		}
	}
}

func TestChangePagedRegisterMapPageRange(t *testing.T) {

	loopback := newPageMappedSX1301()

	tests := []struct {
		page int8
		want error
	}{
		{-1, PageOutOfRangeError},
		{0, nil},
		{1, nil},
		{2, nil},
		{3, PageOutOfRangeError},
	}

	for i, test := range tests {

		got := loopback.changeRegisterPage(test.page)
		if got != test.want {
			t.Errorf("case %d: for page %d expected %v but got %v ", i, test.page, test.want, got)
		}
	}
}

func TestPageMappedRegisterReadWithInvalidName(t *testing.T) {

	loopback := newPageMappedSX1301()

	_, err := loopback.ReadRegisterByName("DUMMY")
	if err != UknownRegisterNameError {
		t.Errorf("expected unknown register but got: %v", err)
	}
}

func TestReadRegisterByNamePMWithPageChange(t *testing.T) {

	loopback := newPageMappedSX1301()

	value, err := loopback.ReadRegisterByName("LGW_TX_STATUS")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	registerDefault := Registers["LGW_TX_STATUS"].defaultValue
	want := make([]byte, 4)
	binary.BigEndian.PutUint32(want, uint32(registerDefault))
	for k := range value {
		if value[k] != want[k+(4-len(value))] {
			t.Errorf("expected the same value as default %v but got %v", want, value)
		}
	}
}

func register2ByteSlice(value int64, length uint8, offset uint8, signed bool) []byte {
	buf := make([]byte, 4) // maximum length of value
	binary.BigEndian.PutUint32(buf, uint32(value))

	var byteToReturn uint8

	switch {
	case length > 24:
		byteToReturn = 4
		break
	case length > 16:
		byteToReturn = 3
		break
	case length > 8:
		byteToReturn = 2
		break
	default:
		byteToReturn = 1
	}

	if signed {
		if length%8 != 0 {
			bitsToClear := 8 - length%8
			for i := 0; i < int(bitsToClear); i++ {
				buf[4-byteToReturn] = clearBit(buf[4-byteToReturn], 7-uint8(i))
			}
		}
	}
	// only register with single values < 1 byte are offset
	if byteToReturn == 1 {
		buf[3] = buf[3] << offset
	}

	return buf[4-byteToReturn:]
}

func TestRegister2ByteSlice(t *testing.T) {
	tests := []struct {
		value  int64
		length uint8
		offset uint8
		signed bool
		want   []byte
	}{
		{1, 8, 0, false, []byte{0x01}},            //single byte values
		{15, 8, 0, false, []byte{0x0F}},           //single byte values
		{127, 8, 0, false, []byte{0x7F}},          //single byte values
		{-120, 8, 0, true, []byte{0x88}},          //single byte values
		{255, 8, 0, false, []byte{0xFF}},          //single byte values
		{-7, 5, 0, true, []byte{0x19}},            //signed 5 bit value
		{-7, 5, 1, true, []byte{0x32}},            //signed 5 bit value offset by 1
		{-7, 5, 2, true, []byte{0x64}},            //signed 5 bit value offset by 2
		{-7, 5, 3, true, []byte{0xC8}},            //signed 5 bit value offset by 2
		{1, 0, 0, false, []byte{0x01}},            // single bit values
		{1, 1, 1, false, []byte{0x02}},            // single bit values
		{1, 1, 2, false, []byte{0x04}},            // single bit values
		{1, 1, 3, false, []byte{0x08}},            // single bit values
		{1, 1, 4, false, []byte{0x10}},            // single bit values
		{1, 1, 5, false, []byte{0x20}},            // single bit values
		{1, 1, 6, false, []byte{0x40}},            // single bit values
		{1, 1, 7, false, []byte{0x80}},            // single bit values
		{1, 0, 0, false, []byte{0x01}},            // four bit values
		{1, 4, 1, false, []byte{0x02}},            // four bit values
		{1, 4, 2, false, []byte{0x04}},            // four bit values
		{1, 4, 3, false, []byte{0x08}},            // four bit values
		{1, 4, 4, false, []byte{0x10}},            // four bit values
		{15, 0, 0, false, []byte{0x0F}},           // four bit values
		{15, 4, 1, false, []byte{0x1E}},           // four bit values
		{15, 4, 2, false, []byte{0x3C}},           // four bit values
		{15, 4, 3, false, []byte{0x78}},           // four bit values
		{15, 4, 4, false, []byte{0xF0}},           // four bit values
		{0, 7, 0, false, []byte{0x00}},            // seven bit values
		{1, 7, 0, false, []byte{0x01}},            // seven bit values
		{1, 7, 1, false, []byte{0x02}},            // seven bit values
		{127, 7, 0, false, []byte{0x7F}},          // seven bit values
		{127, 7, 1, false, []byte{0xFE}},          // seven bit values
		{1, 16, 0, false, []byte{0x00, 0x01}},     //two byte values
		{65280, 16, 0, false, []byte{0xFF, 0x00}}, //two byte values
		{8191, 16, 0, false, []byte{0x1F, 0xFF}},  //two byte values
		{-123, 16, 0, true, []byte{0xFF, 0x85}},   //two byte values
		// 13 bit signed int -384 = 0b 1 1110 1000 0000
		// want other bits in two byte response to be zero
		// -384 {0x1e,0x80} eg for LGW_IF_FREQ
		{-384, 13, 0, true, []byte{0x1e, 0x80}},                    //two byte values
		{4092, 12, 0, false, []byte{0x0F, 0xFC}},                   //two byte values
		{11184810, 24, 0, false, []byte{0xAA, 0xAA, 0xAA}},         //three byte values
		{1789569706, 32, 0, false, []byte{0x6A, 0xAA, 0xAA, 0xAA}}, //four byte values
		{2863311530, 32, 0, false, []byte{0xAA, 0xAA, 0xAA, 0xAA}}, //four byte values
		{-1, 32, 0, true, []byte{0xFF, 0xFF, 0xFF, 0xFF}},          //four byte values signed

	}
	for index, test := range tests {
		got := register2ByteSlice(test.value, test.length, test.offset, test.signed)
		if len(got) != len(test.want) {
			t.Errorf("test %d: expected length to be %d but got length %d", index, len(test.want), len(got))
		}
		for i := range test.want {
			if test.want[i] != got[i] {
				t.Errorf("test %d:%d: expected %02x but got %02x", index, i, test.want[i], got[i])
			}
		}
	}
}

// func TestAllPageMappedRegistersReadbyName(t *testing.T){
// 	loopback := newPageMappedSX1301()
//
// 	value, err := loopback.ReadRegisterByName("LGW_TX_STATUS")
// 	if err != nil {
// 		t.Errorf("expected no error but got: %v", err)
// 	}
// 	registerDefault := Registers["LGW_TX_STATUS"].defaultValue
// 	want := make([]byte, 4)
// 	binary.BigEndian.PutUint32(want, uint32(registerDefault))
// 	for k := range value {
// 		if value[k] != want[k+(4-len(value))] {
// 			t.Errorf("expected the same value as default %v but got %v", want, value)
// 		}
// 	}
// }
