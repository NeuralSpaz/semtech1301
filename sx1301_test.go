package sx1301

import "testing"

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
