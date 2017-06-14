package syncrw

type SynchronousReadWriter interface {
	ReadRegister(addr byte) (byte, error)
	WriteRegister(addr byte, data byte) error
	MultiRead(addr byte, n uint) ([]byte, error)
	MultiWrite(addr byte, data []byte) error
}
