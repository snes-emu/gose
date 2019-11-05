package render

// Renderer defines the interface required to render pixels
type Renderer interface {
	Render(*Screen)
	Stop()
	SetRomTitle(string)
	Run()
}

type rendererFactory func(int, int) (Renderer, error)

//NewRenderer returns one of the available renderers
func NewRenderer(width, height int) (Renderer, error) {
	if rf, ok := renderers["ebiten"]; ok {
		return rf(width, height)
	}

	return renderers["noop"](width, height)
}

var renderers = make(map[string]rendererFactory)

func register(name string, rf rendererFactory) {
	renderers[name] = rf
}
