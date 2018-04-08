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

func (cpu *CPU) initDma() {
	for i := 0; i < 8; i++ {
		cpu.dmaChannels[i] = &dmaChannel{
			dmaEnabled:        false,
			hdmaEnabled:       false,
			transferDirection: true,
			indirectMode:      true,
			addressDecrement:  true,
			fixedTransfer:     true,
			transferMode:      7,
			srcAddr:           0xffff,
			srcBank:           0xff,
			destAddr:          0xff,
			transferSize:      0xffff,
			indirectAddrBank:  0xff,
			hdmaAddr:          0xffff,
			hdmaLineCounter:   0xff,
		}
	}

	cpu.Registers[0x20b] = cpu.mdmaen
	cpu.Registers[0x20c] = cpu.hdmaen
}

// 0x420B - MDMAEN - Select General Purpose DMA Channel(s) and Start Transfer (W)
func (cpu *CPU) mdmaen(x uint8) uint8 {
	for i := uint8(0); i < 8; i++ {
		cpu.dmaChannels[i].dmaEnabled = x&(1<<i) != 0
	}
	return 0
}

// 0x420C - HDMAEN - Select H-Blank DMA (H-DMA) Channel(s) (W)
func (cpu *CPU) hdmaen(x uint8) uint8 {
	for i := uint8(0); i < 8; i++ {
		cpu.dmaChannels[i].hdmaEnabled = x&(1<<i) != 0
	}
	return 0
}
