ARRAY_NAME = "cpu.opcodes"
ARRAY_SIZE = 256

def opcode_array_filler():
    for op in range(ARRAY_SIZE):
        print('{0}[0x{1:X}]=cpu.op{1:X}'.format(ARRAY_NAME, op))


if __name__ == '__main__':
    opcode_array_filler()