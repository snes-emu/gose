package core

type dmaChannel struct {
	dmaEnabled  bool
	hdmaEnabled bool

	transferDirection bool
	indirectMode      bool
	addressDecrement  bool
	fixedTransfer     bool
	transferMode      uint8

	srcAddr uint16
	srcBank uint8

	destAddr uint8

	transferSize     uint16
	indirectAddrBank uint8

	hdmaAddr        uint16
	hdmaLineCounter uint8
}
