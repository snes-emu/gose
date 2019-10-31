package render

// Renderer defines the interface required to render pixels
type Renderer interface {
	Render(*Screen)
	Stop()
	SetRomTitle(string)
}
