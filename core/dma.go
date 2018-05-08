package core

import (
	"fmt"

	"github.com/snes-emu/gose/io"
)

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

	unused uint8
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
			unused:            0xff,
		}
	}

	// Init the dma registers
	cpu.initDmaen()
	cpu.initDmapx()
	cpu.initBbadx()
	cpu.initA1txl()
	cpu.initA1txh()
	cpu.initA1bx()
	cpu.initDasxL()
	cpu.initDasxH()
	cpu.initDasbx()
	cpu.initA2axl()
	cpu.initA2axh()
	cpu.initNtrlx()
	cpu.initUnusedx()
}

func (cpu *CPU) startDma() {
	fmt.Println("dma started")
	for _, channel := range cpu.dmaChannels {
		if !channel.dmaEnabled {
			continue
		}
		transferCount := uint8(0)
		for ok := true; ok; ok = channel.transferSize != 0 {
			cpuBank, cpuOffset := channel.cpuAddress()
			ppuBank, ppuOffset := channel.ppuAddress(transferCount)
			if channel.transferDirection {
				data := cpu.memory.GetByteBank(cpuBank, cpuOffset)
				cpu.memory.SetByteBank(data, ppuBank, ppuOffset)
			} else {
				data := cpu.memory.GetByteBank(ppuBank, ppuOffset)
				cpu.memory.SetByteBank(data, cpuBank, cpuOffset)
			}
			transferCount++
			channel.transferSize--
		}
	}
}

func (dma *dmaChannel) cpuAddress() (uint8, uint16) {
	bank, offset := dma.srcBank, dma.srcAddr
	if !dma.fixedTransfer {
		if dma.addressDecrement {
			dma.srcAddr--
		} else {
			dma.srcAddr++
		}
	}
	return bank, offset
}

func (dma *dmaChannel) ppuAddress(count uint8) (uint8, uint16) {
	bank, offset := uint8(0), uint16(0x2100)
	switch dma.transferMode {
	case 0:
		offset = offset | uint16(dma.destAddr)
	case 1:
		offset = offset | uint16(dma.destAddr+(count&1))
	case 2:
		offset = offset | uint16(dma.destAddr)
	case 3:
		offset = offset | uint16(dma.destAddr+((count>>1)&1))
	case 4:
		offset = offset | uint16(dma.destAddr+(count&3))
	case 5:
		offset = offset | uint16(dma.destAddr+(count&1))
	case 6:
		offset = offset | uint16(dma.destAddr)
	case 7:
		offset = offset | uint16(dma.destAddr+((count>>1)&1))
	}
	return bank, offset
}

func (cpu *CPU) initDmaen() {
	// 0x420B - MDMAEN - Select General Purpose DMA Channel(s) and Start Transfer (W)
	cpu.dmaRegisters[0x20b] = io.NewRegister(
		nil, func(data uint8) {
			for i := uint8(0); i < 8; i++ {
				cpu.dmaChannels[i].dmaEnabled = data&(1<<i) != 0
			}
			cpu.startDma()
		})

	// 0x420C - HDMAEN - Select H-Blank DMA (H-DMA) Channel(s) (W)
	cpu.dmaRegisters[0x20c] = io.NewRegister(
		nil, func(data uint8) {
			for i := uint8(0); i < 8; i++ {
				cpu.dmaChannels[i].hdmaEnabled = data&(1<<i) != 0
			}
		},
	)

}

func (cpu *CPU) initDmapx() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x300+16*i] = io.NewRegister(
			// 0x43x0 - DMAPx - DMA/HDMA Parameters (R/W)
			func() uint8 {
				var res uint8
				if c.transferDirection {
					res |= 0x80
				}
				if c.indirectMode {
					res |= 0x40
				}
				if c.addressDecrement {
					res |= 0x10
				}
				if c.fixedTransfer {
					res |= 0x8
				}
				// Not mandatory but in case transferMode has incorrect bits
				res |= (c.transferMode & 0x7)
				return res
			},
			// 0x43x0 - DMAPx - DMA/HDMA Parameters (R/W)
			func(data uint8) {
				c.transferDirection = data&0x80 != 0
				c.indirectMode = data&0x40 != 0
				c.addressDecrement = data&0x10 != 0
				c.fixedTransfer = data&0x8 != 0
				c.transferMode = data & 0x7
			},
		)
	}
}

func (cpu *CPU) initBbadx() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x301+16*i] = io.NewRegister(
			// 0x43x1 - BBADx - DMA/HDMA I/O-Bus Address (PPU-Bus aka B-Bus) (R/W)
			func() uint8 {
				return c.destAddr
			},
			// 0x43x1 - BBADx - DMA/HDMA I/O-Bus Address (PPU-Bus aka B-Bus) (R/W)
			func(data uint8) {
				c.destAddr = data
			},
		)
	}
}

func (cpu *CPU) initA1txl() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x302+16*i] = io.NewRegister(
			// 0x43x2 - A1TxL - HDMA Table Start Address (low) / DMA Current Addr (low) (R/W)
			func() uint8 {
				return lowerBits(c.srcAddr)
			},
			// 0x43x2 - A1TxL - HDMA Table Start Address (low) / DMA Current Addr (low) (R/W)
			func(data uint8) {
				c.srcAddr = (c.srcAddr & 0xff00) | uint16(data)
			},
		)
	}
}

func (cpu *CPU) initA1txh() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x303+16*i] = io.NewRegister(
			// 0x43x3 - A1TxH - HDMA Table Start Address (hi) / DMA Current Addr (hi) (R/W)
			func() uint8 {
				return upperBits(c.srcAddr)
			},
			// 0x43x3 - A1TxH - HDMA Table Start Address (hi) / DMA Current Addr (hi) (R/W)
			func(data uint8) {
				c.srcAddr = (c.srcAddr & 0x00ff) | (uint16(data) << 8)
			},
		)
	}
}

func (cpu *CPU) initA1bx() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x304+16*i] = io.NewRegister(
			// 0x43x4 - A1Bx - HDMA Table Start Address (bank) / DMA Current Addr (bank) (R/W)
			func() uint8 {
				return c.srcBank
			},
			// 0x43x4 - A1Bx - HDMA Table Start Address (bank) / DMA Current Addr (bank) (R/W)
			func(data uint8) {
				c.srcBank = data
			},
		)
	}
}

func (cpu *CPU) initDasxL() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x305+16*i] = io.NewRegister(
			// 0x43x5 - DASxL - Indirect HDMA Address (low) / DMA Byte-Counter (low) (R/W)
			func() uint8 {
				return lowerBits(c.transferSize)
			},
			// 0x43x5 - DASxL - Indirect HDMA Address (low) / DMA Byte-Counter (low) (R/W)
			func(data uint8) {
				c.transferSize = (c.transferSize & 0xff00) | uint16(data)
			},
		)
	}
}

func (cpu *CPU) initDasxH() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x306+16*i] = io.NewRegister(
			// 0x43x6 - DASxH - Indirect HDMA Address (hi) / DMA Byte-Counter (hi) (R/W)
			func() uint8 {
				return upperBits(c.transferSize)
			},
			// 0x43x6 - DASxH - Indirect HDMA Address (hi) / DMA Byte-Counter (hi) (R/W)
			func(data uint8) {
				c.transferSize = (c.transferSize & 0x00ff) | (uint16(data) << 8)
			},
		)
	}
}

func (cpu *CPU) initDasbx() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x307+16*i] = io.NewRegister(
			// 0x43x7 - DASBx - Indirect HDMA Address (bank) (R/W)
			func() uint8 {
				return c.indirectAddrBank
			},
			// 0x43x7 - DASBx - Indirect HDMA Address (bank) (R/W)
			func(data uint8) {
				c.indirectAddrBank = data
			},
		)
	}
}

func (cpu *CPU) initA2axl() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x308+16*i] = io.NewRegister(
			// 0x43x8 - A2AxL - HDMA Table Current Address (low) (R/W)
			func() uint8 {
				return lowerBits(c.hdmaAddr)
			},
			// 0x43x8 - A2AxL - HDMA Table Current Address (low) (R/W)
			func(data uint8) {
				c.hdmaAddr = (c.hdmaAddr & 0xff00) | uint16(data)
			},
		)
	}
}

func (cpu *CPU) initA2axh() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x309+16*i] = io.NewRegister(
			// 0x43x9 - A2AxH - HDMA Table Current Address (high) (R/W)
			func() uint8 {
				return upperBits(c.hdmaAddr)
			},
			// 0x43x9 - A2AxH - HDMA Table Current Address (high) (R/W)
			func(data uint8) {
				c.hdmaAddr = (c.hdmaAddr & 0x00ff) | (uint16(data) << 8)
			},
		)
	}
}

func (cpu *CPU) initNtrlx() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x30a+16*i] = io.NewRegister(
			// 0x43xA - NTRLx - HDMA Line-Counter (from current Table entry) (R/W)
			func() uint8 {
				return c.hdmaLineCounter
			},
			// 0x43xA - NTRLx - HDMA Line-Counter (from current Table entry) (R/W)
			func(data uint8) {
				c.hdmaLineCounter = data
			},
		)
	}
}

func (cpu *CPU) initUnusedx() {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.dmaRegisters[0x30b+16*i] = io.NewRegister(
			// 0x43xB - UNUSEDx - Unused Byte (R/W)
			func() uint8 {
				return c.unused
			},
			func(data uint8) {
				c.unused = data
			},
		)
	}

	for i := 0; i < 8; i++ {
		// 0x43xF - MIRRx - Read/Write-able mirror of 43xBh (R/W)
		cpu.dmaRegisters[0x30f+16*i] = cpu.dmaRegisters[0x30b+16*i]
	}
}
