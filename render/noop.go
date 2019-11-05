package render

var _ Renderer = &NoOpRenderer{}

type NoOpRenderer struct{}

func (n NoOpRenderer) SetRomTitle(string) {}

func (n NoOpRenderer) Render(*Screen) {}

func (n NoOpRenderer) Stop() {}

func (n NoOpRenderer) Run() {}

func newNoOpRenderer(width, height int) (Renderer, error) {
	return &NoOpRenderer{}, nil
}

func init() {
	register("noop", newNoOpRenderer)
}
