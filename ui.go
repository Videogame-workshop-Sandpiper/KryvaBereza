package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Displays UI
func UIDisplayStats() {
	UIDisplayText(fmt.Sprintf("%s%d%s%d", "X: ", GameData.camera.x, " Y: ", GameData.camera.y), 1)
	UIDisplayText(fmt.Sprintf("%s%d", "Seed: ", GameData.worldSeed), 2)
	UIDisplayText(fmt.Sprintf("%s%d", "Time: ", GameData.gameTime), 3)
}

// Displays text
func UIDisplaySymbol(a string, col int, row int) {
	PrerenderTile(a, sdl.Color{R: 255, G: 255, B: 255}).Blit(nil, Graphic.surface, &sdl.Rect{X: int32((screenAreaWidth + col) * Graphic.charSize.x), Y: int32(row * Graphic.charSize.y), W: 0, H: 0})
}

// Diplays rune
func UIDisplayText(a string, row int) {
	for i := range a {
		UIDisplaySymbol(string([]rune(a)[i]), i, row)
	}
}
