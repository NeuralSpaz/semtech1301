package spitest

import "testing"

func TestEmptyTests(t *testing.T) {}

func TestDummyPin(t *testing.T) {
	var p Pin
	if p != 0 {
		t.Errorf("expected pin default pin value to be false but got %d", p)
	}
	p.Low()
	if p != 0 {
		t.Errorf("expected pin to be HIGH to be but got %d", p)
	}
	p.High()
	if p != 1 {
		t.Errorf("expected pin to be LOW to be but got %d", p)
	}
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

func TestOpenWithNil(t *testing.T) {
	loopback := &Device{}
	_, err := loopback.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
}

func TestOpenWithMemoryMap(t *testing.T) {
	mm := make(map[byte]byte)
	memMaped := &Device{MM: mm}
	_, err := memMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
}

func TestOpenWithPagedMap(t *testing.T) {
	mm := make(map[int8]map[byte]int32)
	pageMaped := &Device{PM: mm}
	_, err := pageMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
}

func TestSimpleLoopback(t *testing.T) {
	// var cs Pin
	looper := &Device{}
	l, err := looper.Open()
	if err != nil {
		t.Errorf("failed to open looper")
	}
	w := []byte{0x01, 0x02}
	r := []byte{0x01, 0x03}
	err = l.Tx(w, r)
	if err != nil {
		t.Errorf("expected No error but got %v ", err)
	}
	for k := range w {
		if w[k] != r[k] {
			t.Errorf("expected slices to be the same")
		}
	}
}

func TestSimpleLoopbackUnqualBuffers(t *testing.T) {
	looper := &Device{}
	l, err := looper.Open()
	if err != nil {
		t.Errorf("failed to open looper")
	}
	w := []byte{0x01, 0x02}
	r := []byte{0x01, 0x03, 0x04}
	err = l.Tx(w, r)
	if err != UnequalBufferError {
		t.Errorf("expected UnequalBufferError but got %v ", err)
	}
}

func TestMemoryMapUnqualBuffers(t *testing.T) {
	mm := make(map[byte]byte)
	memMaped := &Device{MM: mm}
	mmloop, err := memMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	w := make([]byte, 2) // buffers unequalSize
	r := make([]byte, 3)
	err = mmloop.Tx(w, r)
	if err != UnequalBufferError {
		t.Errorf("expected UnequalBufferError but got %v ", err)
	}
}

func TestPageMapUnqualBuffers(t *testing.T) {
	mm := make(map[int8]map[byte]int32)
	pageMaped := &Device{PM: mm}
	pmloop, err := pageMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	w := make([]byte, 2) // buffers unequalSize
	r := make([]byte, 3)
	err = pmloop.Tx(w, r)
	if err != UnequalBufferError {
		t.Errorf("expected UnequalBufferError but got %v ", err)
	}
}

func TestMemoryMapMemoryMapValueError(t *testing.T) {
	mm := make(map[byte]byte)
	memMaped := &Device{MM: mm}
	mmloop, err := memMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	// nothing in memorymap
	w := make([]byte, 2)
	r := make([]byte, 2)
	err = mmloop.Tx(w, r)
	if err != MemoryMapValueError {
		t.Errorf("expected UnequalBufferError but got %v ", err)
	}
}

func TestMemoryMapRead(t *testing.T) {
	mm := make(map[byte]byte)
	memMaped := &Device{MM: mm}
	mmloop, err := memMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	// simple 2 register MemoryMap
	mm[0x01] = 0xAA
	// simulate read
	w := []byte{0x01, 0x00}
	// empty response buffer
	r := make([]byte, 2)
	err = mmloop.Tx(w, r)
	if err != nil {
		t.Errorf("expected No error but got %v ", err)
	}
	// second byte should be value of register
	if r[1] != 0xAA {
		t.Errorf("expected 0xAA but got %v: ", r[1])
	}
}

func TestMemoryMapWrite(t *testing.T) {
	mm := make(map[byte]byte)
	memMaped := &Device{MM: mm}
	mmloop, err := memMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	// simple 1 register MemoryMap
	mm[0x01] = 0xAA
	// simulate write to memory location 0x01 with byte 0xBB
	w := []byte{0x81, 0xBB}
	//
	r := make([]byte, 2)
	err = mmloop.Tx(w, r)
	if err != nil {
		t.Errorf("expected No error but got %v ", err)
	}

	if mm[0x01] != 0xBB {
		t.Errorf("ewpected 0xBB but got %02x: ", mm[0x01])
	}
}

func TestPageMapWithEmptyMap(t *testing.T) {
	pm := make(map[int8]map[byte]int32)
	pageMaped := &Device{PM: pm}
	pmloop, err := pageMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	w := make([]byte, 2)
	r := make([]byte, 2)
	err = pmloop.Tx(w, r)
	if err != PagePointerError {
		t.Errorf("expected invalid page referance but got %v ", err)
	}
}

func TestPageMapReadSimpleMap(t *testing.T) {
	pm := make(map[int8]map[byte]int32)
	pageMaped := &Device{PM: pm}
	pmloop, err := pageMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	defaultPage := make(map[byte]int32)
	pm[-1] = defaultPage
	pm[-1][0x00] = 0
	w := make([]byte, 2)
	r := make([]byte, 2)
	err = pmloop.Tx(w, r)
	if err != nil {
		t.Errorf("expected no error but got %v ", err)
	}
}

func TestPageMapWriteSimpleMap(t *testing.T) {
	pm := make(map[int8]map[byte]int32)
	pageMaped := &Device{PM: pm}
	pmloop, err := pageMaped.Open()
	if err != nil {
		t.Errorf("expect nil error but got %v: ", err)
	}
	defaultPage := make(map[byte]int32)
	pm[-1] = defaultPage
	pm[-1][0x00] = 0
	pm[-1][0x0B] = 0xBB

	w := []byte{0x8B, 0xAA}
	r := make([]byte, 2)

	err = pmloop.Tx(w, r)

	if byte(pm[-1][0x0B]) != 0xAA {
		t.Errorf("expected 0xAA but got %v ", byte(pm[-1][0x0B]))
	}
}
