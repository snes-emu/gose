package apu

type port struct {
	init         bool
	outputBuffer uint8
	inputBuffer  uint8
}

func newPort() *port {
	return &port{}
}

// APUInput writes to the input buffer of the communication port
func (p *port) CPUInput(value uint8) {
	p.inputBuffer = value
	if p.init {
		p.outputBuffer = p.inputBuffer
	}
}

// APUOutput reads from the output buffer of the communication port
func (p *port) CPUOutput() uint8 {
	p.init = true
	return p.outputBuffer
}

// APUInput writes to the input buffer of the communication port
func (p *port) apuInput(value uint8) {
	p.outputBuffer = value
}

// APUOutput reads from the output buffer of the communication port
func (p *port) apuOutput() uint8 {
	return p.inputBuffer
}
