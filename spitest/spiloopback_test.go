package spitest

import "testing"

func TestEmptyTests(t *testing.T) {}

func TestDummyPin(t *testing.T) {
	var p pin
	if p != chipdisable {
		t.Errorf("expected pin default pin value to be false but got %t", p)
	}
	p.chipEnable()
	if p != chipenable {
		t.Errorf("expected pin to be HIGH to be but got %t", p)
	}
	p.chipDisable()
	if p != chipdisable {
		t.Errorf("expected pin to be LOW to be but got %t", p)
	}
}

func TestLoopback(t *testing.T) {
	looper := &Device{}
	l, err := looper.Open()
	if err != nil {
		t.Errorf("failed to open looper")
	}
	w := []byte{0x01, 0x02}
	r := []byte{0x01, 0x03}
	l.Tx(w, r)
	for k := range w {
		if w[k] != r[k] {
			t.Errorf("expected slices to be the same")
		}
	}
}
