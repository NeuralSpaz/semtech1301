package sx1301

import (
	"fmt"
	"time"
)

var loraGateWayStarted bool

// big function does lots of stuff will be hard to test as it is.
func (s *SX1301Spi) start() {

	if loraGateWayStarted {
		fmt.Println("gateway has alread been started, but restarting it now")
	}

	// err := connect()

	// resets registers and shutsdown radios
	// err := softReset()

	// get clocks
	s.WriteRegisterByName("LGW_GLOBAL_EN", 0x00)
	s.WriteRegisterByName("LGW_CLK32M_EN", 0x00)
	// switch on and rest the radios also starts the 32Mhz XTAL
	s.WriteRegisterByName("LGW_RADIO_A_EN", 0x01)
	s.WriteRegisterByName("LGW_RADIO_B_EN", 0x01)
	<-time.After(time.Millisecond * 500)
	s.WriteRegisterByName("LGW_RADIO_RST", 0x01)
	<-time.After(time.Millisecond * 5)
	s.WriteRegisterByName("LGW_RADIO_RST", 0x00)

	//Setup Radios
	// setup_sx125x(0, rf_clkout, rf_enable[0], rf_radio_type[0], rf_rx_freq[0]);
	// setup_sx125x(1, rf_clkout, rf_enable[1], rf_radio_type[1], rf_rx_freq[1]);

	// gives AGC control of GPIOs to enable Tx external digital filter
	s.WriteRegisterByName("LGW_GPIO_MODE", 31) // Set all GPIOs as output
	s.WriteRegisterByName("LGW_GPIO_SELECT_OUTPUT", 2)

	// TODO: @ hal.c line 664 pick it up another day

}
