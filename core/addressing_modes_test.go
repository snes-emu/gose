package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/snes-emu/gose/io"
)

func newTestCPU() *CPU {
	mem := newTestMemory()
	return newCPU(mem, io.NewRegisterFactory())
}

func TestAbsolute(t *testing.T) {
	cpu := newTestCPU()
	cpu.DBR = 0x12
	cpu.K = 0x12

	//opcode arguments
	cpu.memory.SetByteBank(0xFF, 0x12, 1)
	cpu.memory.SetByteBank(0xFF, 0x12, 2)

	assert.Equal(t, cpu.admAbsoluteJ(), uint16(0xFFFF))
	lo, hi := cpu.admAbsoluteP()
	assert.Equal(t, uint32(0x12FFFF), lo)
	assert.Equal(t, uint32(0x130000), hi)
}

func TestAbsoluteX(t *testing.T) {
	cpu := newTestCPU()
	cpu.DBR = 0x12
	cpu.X = 0x000A

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)
	cpu.memory.SetByteBank(0xFF, 0, 2)

	lo, hi := cpu.admAbsoluteXP()
	assert.Equal(t, uint32(0x130008), lo)
	assert.Equal(t, uint32(0x130009), hi)
}

func TestPAbsolute(t *testing.T) {
	cpu := newTestCPU()
	cpu.memory.SetByteBank(0x34, 0, 0x1FFF)
	cpu.memory.SetByteBank(0x56, 0, 0x1FFE)

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)
	cpu.memory.SetByteBank(0x1F, 0, 2)

	assert.Equal(t, uint16(0x3456), cpu.admPAbsoluteJ())
}

func TestPAbsoluteX(t *testing.T) {
	cpu := newTestCPU()
	cpu.K = 0x12
	cpu.X = 0x000A
	cpu.memory.SetByteBank(0x56, 0x12, 8)
	cpu.memory.SetByteBank(0x34, 0x12, 9)

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0x12, 1)
	cpu.memory.SetByteBank(0xFF, 0x12, 2)

	assert.Equal(t, uint16(0x3456), cpu.admPAbsoluteXJ())
}

func TestDirect(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0xFF00

	//opcode arguments
	cpu.memory.SetByteBank(0xFF, 0, 1)

	lo, hi := cpu.admDirectP()
	assert.Equal(t, uint32(0x00FFFF), lo)
	assert.Equal(t, uint32(0x000000), hi)

	cpu.eFlag = true
	cpu.mFlag = true
	lo, _ = cpu.admDirectP()
	assert.Equal(t, uint32(0x00FFFF), lo)
}

func TestDirectX(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0xFF00
	cpu.X = 0x000A

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)

	lo, hi := cpu.admDirectXP()
	assert.Equal(t, uint32(0x000008), lo)
	assert.Equal(t, uint32(0x000009), hi)

	cpu.eFlag = true
	cpu.mFlag = true
	lo, _ = cpu.admDirectXP()
	assert.Equal(t, uint32(0x00FF08), lo)
}

func TestPDirect(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0x1E00
	cpu.DBR = 0x12
	cpu.memory.SetByteBank(0xFF, 0, 0x1EFF)
	cpu.memory.SetByteBank(0xFF, 0, 0x1E00)
	cpu.memory.SetByteBank(0xFF, 0, 0x1F00)

	//opcode arguments
	cpu.memory.SetByteBank(0xFF, 0, 1)

	lo, hi := cpu.admPDirectP()
	assert.Equal(t, uint32(0x12FFFF), lo)
	assert.Equal(t, uint32(0x130000), hi)

	cpu.eFlag = true
	cpu.mFlag = true
	lo, _ = cpu.admPDirectP()
	assert.Equal(t, uint32(0x12FFFF), lo)
}

func TestBDirect(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0x1E00
	cpu.memory.SetByteBank(0xFF, 0, 0x1EFE)
	cpu.memory.SetByteBank(0xFF, 0, 0x1EFF)
	cpu.memory.SetByteBank(0x12, 0, 0x1F00)

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)

	lo, hi := cpu.admBDirectP()
	assert.Equal(t, uint32(0x12FFFF), lo)
	assert.Equal(t, uint32(0x130000), hi)
}

func TestPDirectX(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0x1E00
	cpu.X = 0x000A
	cpu.DBR = 0x12
	cpu.memory.SetByteBank(0xFF, 0, 0x1F08)
	cpu.memory.SetByteBank(0xFF, 0, 0x1F09)
	cpu.memory.SetByteBank(0xFF, 0, 0x1E08)
	cpu.memory.SetByteBank(0xFF, 0, 0x1E09)

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)

	lo, hi := cpu.admPDirectXP()
	assert.Equal(t, uint32(0x12FFFF), lo)
	assert.Equal(t, uint32(0x130000), hi)

	cpu.eFlag = true
	cpu.mFlag = true
	lo, _ = cpu.admPDirectXP()
	assert.Equal(t, uint32(0x12FFFF), lo)
}

func TestPDirectY(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0x1E00
	cpu.Y = 0x000A
	cpu.DBR = 0x12
	cpu.memory.SetByteBank(0xFF, 0, 0x1F00)
	cpu.memory.SetByteBank(0xFF, 0, 0x1E00)
	cpu.memory.SetByteBank(0xFE, 0, 0x1EFF)

	//opcode arguments
	cpu.memory.SetByteBank(0xFF, 0, 1)

	lo, hi := cpu.admPDirectYP()
	assert.Equal(t, uint32(0x130008), lo)
	assert.Equal(t, uint32(0x130009), hi)

	cpu.eFlag = true
	cpu.mFlag = true
	lo, _ = cpu.admPDirectYP()
	assert.Equal(t, uint32(0x130008), lo)
}

func TestBDirectY(t *testing.T) {
	cpu := newTestCPU()
	cpu.D = 0x1E00
	cpu.Y = 0x000A
	cpu.memory.SetByteBank(0x12, 0, 0x1F00)
	cpu.memory.SetByteBank(0xFC, 0, 0x1EFE)
	cpu.memory.SetByteBank(0xFF, 0, 0x1EFF)

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)

	lo, hi := cpu.admBDirectYP()
	assert.Equal(t, uint32(0x130006), lo)
	assert.Equal(t, uint32(0x130007), hi)
}

func TestImmediate(t *testing.T) {
	cpu := newTestCPU()

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)
	cpu.memory.SetByteBank(0xFF, 0, 2)

	lo, hi := cpu.admImmediateM()
	assert.Equal(t, uint8(0xFE), lo)
	assert.Equal(t, uint8(0xFF), hi)

}

func TestLong(t *testing.T) {
	cpu := newTestCPU()

	//opcode arguments
	cpu.memory.SetByteBank(0xFF, 0, 1)
	cpu.memory.SetByteBank(0xFF, 0, 2)
	cpu.memory.SetByteBank(0x12, 0, 3)

	PC, K := cpu.admLongJ()
	assert.Equal(t, uint8(0x12), K)
	assert.Equal(t, uint16(0xFFFF), PC)

	lo, hi := cpu.admLongP()
	assert.Equal(t, uint32(0x12FFFF), lo)
	assert.Equal(t, uint32(0x130000), hi)
}

func TestLongX(t *testing.T) {
	cpu := newTestCPU()
	cpu.X = 0x000A

	//opcode arguments
	cpu.memory.SetByteBank(0xFE, 0, 1)
	cpu.memory.SetByteBank(0xFF, 0, 2)
	cpu.memory.SetByteBank(0x12, 0, 3)

	lo, hi := cpu.admLongXP()
	assert.Equal(t, uint32(0x130008), lo)
	assert.Equal(t, uint32(0x130009), hi)
}

func TestSourceDestination(t *testing.T) {
	cpu := newTestCPU()
	cpu.C = 0x0002
	cpu.X = 0xFFFE
	cpu.Y = 0xFFFF

	//opcode arguments
	cpu.memory.SetByteBank(0x12, 0, 1)
	cpu.memory.SetByteBank(0x34, 0, 2)

	sBank, sOffset, dBank, dOffset := cpu.admSourceDestination()
	assert.Equal(t, uint8(0x12), sBank)
	assert.Equal(t, uint16(0xFFFE), sOffset)
	assert.Equal(t, uint8(0x34), dBank)
	assert.Equal(t, uint16(0xFFFF), dOffset)
}

func TestStackS(t *testing.T) {
	cpu := newTestCPU()
	cpu.S = 0xFF10

	//opcode arguments
	cpu.memory.SetByteBank(0xFA, 0, 1)

	lo, hi := cpu.admStackSP()
	assert.Equal(t, uint32(0x00000A), lo)
	assert.Equal(t, uint32(0x00000B), hi)
}

func TestStackSY(t *testing.T) {
	cpu := newTestCPU()
	cpu.S = 0xFF10
	cpu.DBR = 0x12
	cpu.Y = 0x0050
	cpu.memory.SetByteBank(0xF0, 0, 0x000A)
	cpu.memory.SetByteBank(0xFF, 0, 0x000B)

	//opcode arguments
	cpu.memory.SetByteBank(0xFA, 0, 1)

	lo, hi := cpu.admPStackSYP()
	assert.Equal(t, uint32(0x130040), lo)
	assert.Equal(t, uint32(0x130041), hi)
}
