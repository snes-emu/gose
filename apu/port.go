package apu

type port struct {
	outputBuffer uint8
	inputBuffer  uint8
}

func newPort() *port {
	return &port{}
}

// APUInput writes to the input buffer of the communication port
func (p *port) APUInput(value uint8) {
	p.inputBuffer = value
}

// APUOutput reads from the output buffer of the communication port
func (p *port) APUOutput() uint8 {
	return p.outputBuffer
}
