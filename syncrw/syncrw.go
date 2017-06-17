package syncrw

type SynchronousReadWriter interface {
	ReadRegister(addr byte) (byte, error)
	WriteRegister(addr byte, data byte) error
	// MultiRead(addr byte, n uint) ([]byte, error)
	// MultiWrite(addr byte, data []byte) error
	// ChipEnable()
	// ChipDisable()
}

type DuplexReadWriter interface {
	ReadFromReg(addr byte, n uint) (rx []byte, err error)
	WriteToReg(addr byte, tx []byte) (err error)
}
