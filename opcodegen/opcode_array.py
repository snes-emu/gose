ARRAY_NAME = "cpu.opcodes"
ARRAY_SIZE = 256


def opcode_array_filler():
    return '\n'.join(
        ['{0}[0x{1:X}]=cpu.op{1:X}'.format(ARRAY_NAME, op) for op in range(ARRAY_SIZE)
         ]
    )


if __name__ == '__main__':
    print(opcode_array_filler())
