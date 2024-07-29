package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Displays UI
func UIDisplayStats() {
	UIDisplayText(fmt.Sprintf("%s%d%s%d%s%d", "X: ", GameData.camera.x, " Y: ", GameData.camera.y, " Z: ", GameData.camera.z), 1)
	UIDisplayText(fmt.Sprintf("%s%d", "Seed: ", GameData.worldSeed), 2)
	UIDisplayText(fmt.Sprintf("%s%d", "Time: ", GameData.gameTime), 3)

	UIDrawCompass(screenAreaHeight - 5)
}

// Displays text
func UIDisplaySymbol(a string, col int, row int) {
<<<<<<< HEAD:ui.go
	var prerender = PrerenderTileBlended(a, sdl.Color{R: 255, G: 255, B: 255})
	prerender.Blit(nil, Graphic.surface, &sdl.Rect{X: int32((screenAreaWidth + col) * Graphic.charSize.x), Y: int32(row * Graphic.charSize.y), W: 0, H: 0})
	prerender.Free()
=======
	var prerenderedTile = PrerenderTileBlended(a, sdl.Color{R: 255, G: 255, B: 255})
	prerenderedTile.Blit(nil, Graphic.surface, &sdl.Rect{X: int32((screenAreaWidth + col) * Graphic.charSize.x), Y: int32(row * Graphic.charSize.y), W: 0, H: 0})
	prerenderedTile.Free();
>>>>>>> 24dd5dcfea0bb2bb3f417252706c958a4398656f:code/ui.go
}

// Diplays rune
func UIDisplayText(a string, row int) {
	var I int = 0
	for i, c := range a {
		UIDisplaySymbol(string(c), I, row)
		if i == -1 {
			i = -1
		}
		I++
	}
}

// Display compass, that shows camera's dierction
func UIDrawCompass(row int) {
	UIDisplayText("  N ", row)
	UIDisplayText(" ┌─┐", row+1)
	switch GameData.cameraDir {
	case 0:
		UIDisplayText("W│▲│E", row+2)
	case 1:
		UIDisplayText("W│►│E", row+2)
	case 2:
		UIDisplayText("W│▼│E", row+2)
	case 3:
		UIDisplayText("W│◄│E", row+2)
	}
	UIDisplayText(" └─┘", row+3)
	UIDisplayText("  S ", row+4)
}
