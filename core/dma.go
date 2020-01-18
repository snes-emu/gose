package core

import (
	"fmt"

	"github.com/snes-emu/gose/log"

	"github.com/snes-emu/gose/bit"
	"github.com/snes-emu/gose/io"
)

type dmaChannel struct {
	dmaEnabled bool

	transferDirection bool
	addressDecrement  bool
	fixedTransfer     bool
	transferMode      uint8

	srcAddr uint16
	srcBank uint8

	destAddr uint8

	transferSize     uint16
	indirectAddrBank uint8

	unused uint8
}

type hdmaChannel struct {
	enabled bool

	transferDirection bool
	transferMode      uint8

	destAddr uint8

	tableStartAddress   uint16
	tableStartBank      uint8
	tableCurrentAddress uint16

	indirectMode    bool
	indirectAddress uint16
	indirectBank    uint8

	lineCounter uint8
	doTransfer  bool
	completed   bool
}

func (cpu *CPU) initDma(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		cpu.dmaChannels[i] = &dmaChannel{
			dmaEnabled:        false,
			transferDirection: true,
			addressDecrement:  true,
			fixedTransfer:     true,
			transferMode:      7,
			srcAddr:           0xffff,
			srcBank:           0xff,
			destAddr:          0xff,
			transferSize:      0xffff,
			indirectAddrBank:  0xff,
			unused:            0xff,
		}

		cpu.hdmaChannels[i] = &hdmaChannel{
			enabled:             false,
			transferDirection:   true,
			transferMode:        7,
			destAddr:            0xFF,
			tableStartAddress:   0xFFFF,
			tableStartBank:      0xFF,
			tableCurrentAddress: 0xFFFF,
			indirectMode:        true,
			indirectAddress:     0xFFFF,
			indirectBank:        0xFF,
			lineCounter:         0xFF,
			doTransfer:          false,
		}
	}

	// Init the dma registers
	cpu.initDmaen(rf)
	cpu.initDmapx(rf)
	cpu.initBbadx(rf)
	cpu.initA1txl(rf)
	cpu.initA1txh(rf)
	cpu.initA1bx(rf)
	cpu.initDasxL(rf)
	cpu.initDasxH(rf)
	cpu.initDasbx(rf)
	cpu.initA2axl(rf)
	cpu.initA2axh(rf)
	cpu.initNtrlx(rf)
	cpu.initUnusedx(rf)
}

func (cpu *CPU) startDma() {
	log.Debug("dma started")
	for _, channel := range cpu.dmaChannels {
		if !channel.dmaEnabled {
			continue
		}
		transferCount := uint8(0)
		for ok := true; ok; ok = channel.transferSize != 0 {
			cpuBank, cpuOffset := channel.cpuAddress()
			ppuBank, ppuOffset := channel.ppuAddress(transferCount)
			if channel.transferDirection {
				data := cpu.memory.GetByteBank(ppuBank, ppuOffset)
				cpu.memory.SetByteBank(data, cpuBank, cpuOffset)
			} else {
				data := cpu.memory.GetByteBank(cpuBank, cpuOffset)
				cpu.memory.SetByteBank(data, ppuBank, ppuOffset)
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

func (cpu *CPU) reloadHDMA() {
	for _, channel := range cpu.hdmaChannels {
		if channel.enabled {
			channel.tableCurrentAddress = channel.tableStartAddress
			channel.lineCounter = cpu.memory.GetByteBank(channel.tableStartBank, channel.getTableCurrentAddress())
			channel.completed = channel.lineCounter == 0
			if channel.indirectMode {
				indirectAddrLo := cpu.memory.GetByteBank(channel.tableStartBank, channel.getTableCurrentAddress())
				indirectAddrHi := cpu.memory.GetByteBank(channel.tableStartBank, channel.getTableCurrentAddress())
				channel.indirectAddress = bit.JoinUint16(indirectAddrLo, indirectAddrHi)
			}
			channel.doTransfer = true

		}
	}
}

func (cpu *CPU) doHDMA() {
	for i, channel := range cpu.hdmaChannels {
		if !channel.enabled || channel.completed {
			continue
		}
		log.Info(fmt.Sprintln("HDMA channel:", i))

		if channel.doTransfer {
			for c := uint8(0); c < bytesPerCycle[channel.transferMode]; c++ {
				srcBank, srcOffset := channel.cpuAddress()
				data := cpu.memory.GetByteBank(srcBank, srcOffset)

				dstBank, dstOffset := channel.ppuAddress(c)
				cpu.memory.SetByteBank(data, dstBank, dstOffset)
			}
		}
		channel.lineCounter--
		channel.doTransfer = channel.lineCounter&0x80 != 0
		lineCounter := channel.lineCounter & 0x7F
		if lineCounter == 0 {
			channel.lineCounter = cpu.memory.GetByteBank(channel.tableStartBank, channel.getTableCurrentAddress())
			channel.completed = channel.lineCounter == 0
			if channel.indirectMode {
				indirectAddrLo := cpu.memory.GetByteBank(channel.tableStartBank, channel.getTableCurrentAddress())
				indirectAddrHi := cpu.memory.GetByteBank(channel.tableStartBank, channel.getTableCurrentAddress())
				channel.indirectAddress = bit.JoinUint16(indirectAddrLo, indirectAddrHi)
			}
		}
	}
}

func (dma *hdmaChannel) getTableCurrentAddress() uint16 {
	addr := dma.tableCurrentAddress
	dma.tableCurrentAddress++

	return addr
}

func (dma *hdmaChannel) getIndirectAddress() uint16 {
	addr := dma.indirectAddress
	dma.indirectAddress++

	return addr
}

func (dma *hdmaChannel) cpuAddress() (uint8, uint16) {
	if dma.indirectMode {
		return dma.indirectBank, dma.getIndirectAddress()
	}
	return dma.tableStartBank, dma.getTableCurrentAddress()
}

func (dma *hdmaChannel) ppuAddress(count uint8) (uint8, uint16) {
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

func (cpu *CPU) initDmaen(rf *io.RegisterFactory) {
	// 0x420B - MDMAEN - Select General Purpose DMA Channel(s) and Start Transfer (W)
	cpu.ioRegisters[0x20b] = rf.NewRegister(
		nil, func(data uint8) {
			for i := uint8(0); i < 8; i++ {
				cpu.dmaChannels[i].dmaEnabled = data&(1<<i) != 0
			}
			cpu.startDma()
		},
		"MDMAEN",
	)

	// 0x420C - HDMAEN - Select H-Blank DMA (H-DMA) Channel(s) (W)
	cpu.ioRegisters[0x20c] = rf.NewRegister(
		nil, func(data uint8) {
			fmt.Printf("hdma enable: %02X\n", data)
			for i := uint8(0); i < 8; i++ {
				cpu.hdmaChannels[i].enabled = data&(1<<i) != 0
			}
		},
		"HMDMAEN",
	)

}

func (cpu *CPU) initDmapx(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x300+16*i] = rf.NewRegister(
			// 0x43x0 - DMAPx - DMA/HDMA Parameters (R/W)
			func() uint8 {
				var res uint8
				if c.transferDirection {
					res |= 0x80
				}
				if hc.indirectMode {
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
				hc.transferDirection = data&0x80 != 0
				hc.indirectMode = data&0x40 != 0
				c.addressDecrement = data&0x10 != 0
				c.fixedTransfer = data&0x8 != 0
				c.transferMode = data & 0x7
				hc.transferMode = data & 0x7
			},
			fmt.Sprintf("DMAP%v", i),
		)
	}
}

func (cpu *CPU) initBbadx(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x301+16*i] = rf.NewRegister(
			// 0x43x1 - BBADx - DMA/HDMA I/O-Bus Address (PPU-Bus aka B-Bus) (R/W)
			func() uint8 {
				return c.destAddr
			},
			// 0x43x1 - BBADx - DMA/HDMA I/O-Bus Address (PPU-Bus aka B-Bus) (R/W)
			func(data uint8) {
				c.destAddr = data
				hc.destAddr = data
			},
			fmt.Sprintf("BBAD%v", i),
		)
	}
}

func (cpu *CPU) initA1txl(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x302+16*i] = rf.NewRegister(
			// 0x43x2 - A1TxL - HDMA Table Start Address (low) / DMA Current Addr (low) (R/W)
			func() uint8 {
				return bit.LowByte(c.srcAddr)
			},
			// 0x43x2 - A1TxL - HDMA Table Start Address (low) / DMA Current Addr (low) (R/W)
			func(data uint8) {
				c.srcAddr = bit.SetLowByte(c.srcAddr, data)
				hc.tableStartAddress = bit.SetLowByte(hc.tableStartAddress, data)
			},
			fmt.Sprintf("A1T%vL", i),
		)
	}
}

func (cpu *CPU) initA1txh(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x303+16*i] = rf.NewRegister(
			// 0x43x3 - A1TxH - HDMA Table Start Address (hi) / DMA Current Addr (hi) (R/W)
			func() uint8 {
				return bit.HighByte(c.srcAddr)
			},
			// 0x43x3 - A1TxH - HDMA Table Start Address (hi) / DMA Current Addr (hi) (R/W)
			func(data uint8) {
				c.srcAddr = bit.SetHighByte(c.srcAddr, data)
				hc.tableStartAddress = bit.SetHighByte(hc.tableStartAddress, data)
			},
			fmt.Sprintf("A1T%vH", i),
		)
	}
}

func (cpu *CPU) initA1bx(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x304+16*i] = rf.NewRegister(
			// 0x43x4 - A1Bx - HDMA Table Start Address (bank) / DMA Current Addr (bank) (R/W)
			func() uint8 {
				return c.srcBank
			},
			// 0x43x4 - A1Bx - HDMA Table Start Address (bank) / DMA Current Addr (bank) (R/W)
			func(data uint8) {
				c.srcBank = data
				hc.tableStartBank = data
			},
			fmt.Sprintf("A1B%v", i),
		)
	}
}

func (cpu *CPU) initDasxL(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x305+16*i] = rf.NewRegister(
			// 0x43x5 - DASxL - Indirect HDMA Address (low) / DMA Byte-Counter (low) (R/W)
			func() uint8 {
				return bit.LowByte(c.transferSize)
			},
			// 0x43x5 - DASxL - Indirect HDMA Address (low) / DMA Byte-Counter (low) (R/W)
			func(data uint8) {
				c.transferSize = bit.SetLowByte(c.transferSize, data)
				hc.indirectAddress = bit.SetLowByte(hc.indirectAddress, data)
			},
			fmt.Sprintf("DAS%vL", i),
		)
	}
}

func (cpu *CPU) initDasxH(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x306+16*i] = rf.NewRegister(
			// 0x43x6 - DASxH - Indirect HDMA Address (hi) / DMA Byte-Counter (hi) (R/W)
			func() uint8 {
				return bit.HighByte(c.transferSize)
			},
			// 0x43x6 - DASxH - Indirect HDMA Address (hi) / DMA Byte-Counter (hi) (R/W)
			func(data uint8) {
				c.transferSize = bit.SetHighByte(c.transferSize, data)
				hc.indirectAddress = bit.SetHighByte(hc.indirectAddress, data)
			},
			fmt.Sprintf("DAS%vH", i),
		)
	}
}

func (cpu *CPU) initDasbx(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x307+16*i] = rf.NewRegister(
			// 0x43x7 - DASBx - Indirect HDMA Address (bank) (R/W)
			func() uint8 {
				return hc.indirectBank
			},
			// 0x43x7 - DASBx - Indirect HDMA Address (bank) (R/W)
			func(data uint8) {
				hc.indirectBank = data
			},
			fmt.Sprintf("DASB%v", i),
		)
	}
}

func (cpu *CPU) initA2axl(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x308+16*i] = rf.NewRegister(
			// 0x43x8 - A2AxL - HDMA Table Current Address (low) (R/W)
			func() uint8 {
				return bit.LowByte(hc.tableCurrentAddress)
			},
			// 0x43x8 - A2AxL - HDMA Table Current Address (low) (R/W)
			func(data uint8) {
				hc.tableCurrentAddress = bit.SetLowByte(hc.tableCurrentAddress, data)
			},
			fmt.Sprintf("A2A%vL", i),
		)
	}
}

func (cpu *CPU) initA2axh(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x309+16*i] = rf.NewRegister(
			// 0x43x9 - A2AxH - HDMA Table Current Address (high) (R/W)
			func() uint8 {
				return bit.HighByte(hc.tableCurrentAddress)
			},
			// 0x43x9 - A2AxH - HDMA Table Current Address (high) (R/W)
			func(data uint8) {
				hc.tableCurrentAddress = bit.SetHighByte(hc.tableCurrentAddress, data)
			},
			fmt.Sprintf("A2A%vH", i),
		)
	}
}

func (cpu *CPU) initNtrlx(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		hc := cpu.hdmaChannels[i]
		cpu.ioRegisters[0x30a+16*i] = rf.NewRegister(
			// 0x43xA - NTRLx - HDMA Line-Counter (from current Table entry) (R/W)
			func() uint8 {
				return hc.lineCounter
			},
			// 0x43xA - NTRLx - HDMA Line-Counter (from current Table entry) (R/W)
			func(data uint8) {
				hc.lineCounter = data
			},
			fmt.Sprintf("NTRL%v", i),
		)
	}
}

func (cpu *CPU) initUnusedx(rf *io.RegisterFactory) {
	for i := 0; i < 8; i++ {
		c := cpu.dmaChannels[i]
		cpu.ioRegisters[0x30b+16*i] = rf.NewRegister(
			// 0x43xB - UNUSEDx - Unused Byte (R/W)
			func() uint8 {
				return c.unused
			},
			func(data uint8) {
				c.unused = data
			},
			fmt.Sprintf("UNUSED%v", i),
		)
	}

	for i := 0; i < 8; i++ {
		// 0x43xF - MIRRx - Read/Write-able mirror of 43xBh (R/W)
		cpu.ioRegisters[0x30f+16*i] = cpu.ioRegisters[0x30b+16*i]
	}
}

var bytesPerCycle = [8]uint8{1, 2, 2, 4, 4, 4, 2, 4}
