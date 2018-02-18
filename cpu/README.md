# gose: package CPU

This package is responsible for the 65C816 CPU emulation.
The naming convention is the same as the one used [there](http://6502.org/tutorials/65c816opcodes.html)

## File structure

### CPU structure

The cpu structure is defined in the _cpu.go_ file.
Getters and setters for the cpu registers are located in the _register.go_ file

### Operations

All operations are defined in the files with the _op_ prefix. In a specific file there can be multiple operations and all the opcodes linked to theses operations.

- _op\_bit.go_ contains _BIT_ operation
- _op\_branch.go_ contains _BCC_, _BCS_, _BEQ_, _BMI_, _BNE_, _BPL_, _BRA_, _BVC_, _BVS_, _BRL_, _BCC_ operations
- _op\_carry.go_ contains _ADC_, _SBC_ operations
- _op\_compare.go_ contains _CMP_, _CPX_, _CPY_ operations
- _op\_decinc.go_ contains _DEC_, _INC_ operations
- _op\_exchange.go_ contains _XBA_, _XCE_ operations
- _op\_flags.go_ contains _CLC_, _CLD_, _CLI_, _CLV_, _SEC_, _SED_, _SEI_, _REP_, _SEP_ operations
- _op\_jump.go_ contains _JMP_, _JSL_, _JSR_ operations
- _op\_ldst.go_ contains _LDA_, _LDX_, _LDY_, _STA_, _STX_, _STY_, _STZ_ operations
- _op\_move.go_ contains _MVN_, _MVP_ operations
- _op\_nop.go_ contains _NOP_, _WDM_ operations
- _op\_return.go_ contains _RTI_, _RTL_, _RTS_ operations
- _op_shift.go_ contains _ASL_, _LSR_, _ROL_, _ROR_ operations
- _op\_stack.go_ contains _PHA_, _PHX_, _PHY_, _PLA_, _PLX_, _PLY_ operations
- _op\_tbit.go_ contains _TRB_, _TSB_ operations
- _op\_transfer.go_ contains _TCD_, _TCS_, _TDC_, _TSC_ operations
- _hardware\_interrupt.go_ contains _ABORT_, _NMI_, _RESET_, _IRQ_ operations
- _software\_interrupt.go_ contains _BRK_, _COP_ operations

### Addressing modes

All addressing modes are in the files beginning with the prefix _adm_.
