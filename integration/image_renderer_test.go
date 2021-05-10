package integration

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/snes-emu/gose/core"
)

func TestImageRenderer(t *testing.T) {
	renderer := NewImageRenderer(core.WIDTH, core.HEIGHT)
	emu := core.New(renderer, true)
	emu.ReadROM("./data/inc/rom.sfc")

	emu.Start()
	defer emu.Stop()
	emu.StepAndWait(10_000_000)

	expected, err := ioutil.ReadFile("./data/inc/expected.png")
	require.NoError(t, err)

	img, err := png.Decode(bytes.NewReader(expected))
	require.NoError(t, err)

	if img != renderer.Image {
		assert.NoError(t, renderer.SaveToFile("/tmp/test.png"))
		assert.Fail(t, "Expected rendering do not match, see /tmp/test.png")
	}
}
