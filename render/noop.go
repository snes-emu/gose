package render

var _ Renderer = &NoOpRenderer{}

type NoOpRenderer struct{}

func (n NoOpRenderer) SetRomTitle(string) {}

func (n NoOpRenderer) Render(*Screen) {}

func (n NoOpRenderer) Stop() {}
