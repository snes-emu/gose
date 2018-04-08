package core

import "fmt"

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

}

func (cpu *CPU) SetDma(addr uint16, data uint8) {
	c := cpu.dmaChannels[addr>>4&0x1]
	switch addr & 0xf0f {
	// 0x420B - MDMAEN - Select General Purpose DMA Channel(s) and Start Transfer (W)
	case 0x20b:
		for i := uint8(0); i < 8; i++ {
			cpu.dmaChannels[i].dmaEnabled = data&(1<<i) != 0
		}

	// 0x420C - HDMAEN - Select H-Blank DMA (H-DMA) Channel(s) (W)
	case 0x20c:
		for i := uint8(0); i < 8; i++ {
			cpu.dmaChannels[i].hdmaEnabled = data&(1<<i) != 0
		}

	// 0x43x0 - DMAPx - DMA/HDMA Parameters (R/W)
	case 0x300:
		c.transferDirection = data&0x80 != 0
		c.indirectMode = data&0x40 != 0
		c.addressDecrement = data&0x10 != 0
		c.fixedTransfer = data&0x8 != 0
		c.transferMode = data & 0x7

	// 0x43x1 - BBADx - DMA/HDMA I/O-Bus Address (PPU-Bus aka B-Bus) (R/W)
	case 0x301:
		c.destAddr = data

	// 0x43x2 - A1TxL - HDMA Table Start Address (low) / DMA Current Addr (low) (R/W)
	case 0x302:
		c.srcAddr = (c.srcAddr & 0xff00) | uint16(data)

	// 0x43x3 - A1TxH - HDMA Table Start Address (hi) / DMA Current Addr (hi) (R/W)
	case 0x303:
		c.srcAddr = (c.srcAddr & 0x00ff) | (uint16(data) << 8)

	// 0x43x4 - A1Bx - HDMA Table Start Address (bank) / DMA Current Addr (bank) (R/W)
	case 0x304:
		c.srcBank = data

	// 0x43x5 - DASxL - Indirect HDMA Address (low) / DMA Byte-Counter (low) (R/W)
	case 0x305:
		c.transferSize = (c.transferSize & 0xff00) | uint16(data)

	// 0x43x6 - DASxH - Indirect HDMA Address (hi) / DMA Byte-Counter (hi) (R/W)
	case 0x306:
		c.transferSize = (c.transferSize & 0x00ff) | (uint16(data) << 8)

	// 0x43x7 - DASBx - Indirect HDMA Address (bank) (R/W)
	case 0x307:
		c.indirectAddrBank = data

	// 0x43x8 - A2AxL - HDMA Table Current Address (low) (R/W)
	case 0x308:
		c.hdmaAddr = (c.hdmaAddr & 0xff00) | uint16(data)

	// 0x43x9 - A2AxH - HDMA Table Current Address (high) (R/W)
	case 0x309:
		c.hdmaAddr = (c.hdmaAddr & 0x00ff) | (uint16(data) << 8)

	// 0x43xA - NTRLx - HDMA Line-Counter (from current Table entry) (R/W)
	case 0x30a:
		c.hdmaLineCounter = data

	default:
		// If this triggers try implementing the unused and unknown registers
		panic(fmt.Sprintf("Unknown register: %v", addr&0xf0f))
	}
}
