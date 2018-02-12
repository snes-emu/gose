import sys

# Opcode template
template = '''func (cpu *CPU) op{}(){{
    dataHi, dataLo := cpu.adm{}()
    cpu.{}(dataHi, dataLo)
    cpu.cycles += {}
}}

'''

# Cycles variable correspondance
cycles_corres = {
    "w": "utils.BoolToUint16[cpu.getDLRegister() == 0]",
    "p": "utils.BoolToUint16[cpu.pFlag]",
    "m": "utils.BoolToUint16[cpu.mFlag]",
    "x": "utils.BoolToUint16[cpu.xFlag]",
    "e": "utils.BoolToUint16[cpu.eFlag]",
}

# Possible addressing modes
adm_modes = {
    '(dir,X)': 'PDirectX',
    'stk,S': 'StackS',
    'dir': 'Direct',
    '[dir]': 'BDirect',
    'imm': 'Immediate',
    'abs': 'Absolute',
    'long': 'Long',
    '(dir),Y': 'PDirectY',
    '(dir)': 'PDirect',
    '(stk,S),Y': 'PStackSY',
    'dir,X': 'DirectX',
    '[dir],Y': 'BDirectY',
    'abs,Y': 'AbsoluteY',
    'abs,X': 'AbsoluteX',
    'long,X': 'LongX',
    'acc': 'Accumulator'
}


def generate_code(rows):
    """
        generate_code takes the documentation rows from http://6502.org/tutorials/65c816opcodes.html#6 and returns the corresponding go code
        params:
            - rows a list of string rows like '61 2   7-m+w       (dir,X)   mm....mm . ADC ($10,X)'
    
        returns:
            the go code for theses instructions
    """

    datas = [[x for x in row.split(' ') if x != '' ] for row in rows]

    code = ""

    for data in datas:
        opcode, cycles, adm_mode, operation = data[0], data[2], data[3], data[6]

        if not(adm_modes.get(adm_mode, False)):
            raise ValueError("Adm mode not implemented :" + adm_mode)

        if 't' in cycles:
            raise ValueError("t flag found in cycles expression and not implemented :" + cycles)

        cycles = ''.join(cycles_corres[char] if char in cycles_corres.keys() else char for char in cycles)

        code += template.format(opcode, adm_modes[adm_mode], operation.lower(), cycles)

    return code



if __name__ == '__main__':
    with open(sys.argv[1], 'r') as file:
        print(generate_code(file.readlines()))