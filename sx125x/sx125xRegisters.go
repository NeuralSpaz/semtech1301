package sx125x

// SX1257 frequency setting :
// F_register(24bit) = F_rf (Hz) / F_step(Hz)
//                   = F_rf (Hz) * 2^19 / F_xtal(Hz)
//                   = F_rf (Hz) * 2^19 / 32e6
//                   = F_rf (Hz) * 256/15625
//
// SX1255 frequency setting :
// F_register(24bit) = F_rf (Hz) / F_step(Hz)
//                   = F_rf (Hz) * 2^20 / F_xtal(Hz)
//                   = F_rf (Hz) * 2^20 / 32e6
//                   = F_rf (Hz) * 512/15625

const (
	SX125x_TX_DAC_CLK_SEL  byte = 1  // 0:int, 1:ext
	SX125x_TX_DAC_GAIN     byte = 2  //3:0, 2:-3, 1:-6, 0:-9 dBFS (default 2)
	SX125x_TX_MIX_GAIN     byte = 14 //-38 + 2*TxMixGain dB (default 14)
	SX125x_TX_PLL_BW       byte = 1  //0:75, 1:150, 2:225, 3:300 kHz (default 3)
	SX125x_TX_ANA_BW       byte = 0  //17.5 / 2*(41-TxAnaBw) MHz (default 0)
	SX125x_TX_DAC_BW       byte = 5  //24 + 8*TxDacBw Nb FIR taps (default 2)
	SX125x_RX_LNA_GAIN     byte = 1  //1 to 6, 1 highest gain
	SX125x_RX_BB_GAIN      byte = 12 //0 to 15 , 15 highest gain
	SX125x_LNA_ZIN         byte = 1  //0:50, 1:200 Ohms (default 1)
	SX125x_RX_ADC_BW       byte = 7  //0 to 7, 2:100<BW<200, 5:200<BW<400,7:400<BW kHz SSB (default 7)
	SX125x_RX_ADC_TRIM     byte = 6  //0 to 7, 6 for 32MHz ref, 5 for 36MHz ref
	SX125x_RX_BB_BW        byte = 0  //0:750, 1:500, 2:375; 3:250 kHz SSB (default 1, max 3)
	SX125x_RX_PLL_BW       byte = 0  //0:75, 1:150, 2:225, 3:300 kHz (default 3, max 3)
	SX125x_ADC_TEMP        byte = 0  //ADC temperature measurement mode (default 0)
	SX125x_XOSC_GM_STARTUP byte = 13 //(default 13)
	SX125x_XOSC_DISABLE    byte = 2  //Disable of Xtal Oscillator blocks bit0:regulator, bit1:core(gm), bit2:amplifier
)
