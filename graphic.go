package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	fontPath    = "./data/NotoSansMono.ttf" //Path to ttf file
	fontSize    = 16                        //Font size
	tileSetSize = 20                        //Size of tileset
)

var Graphic struct {
	window  *sdl.Window
	font    *ttf.Font
	surface *sdl.Surface

	tiles    [tileSetSize]*sdl.Surface //Set of tiles
	tileMap  map[string]int            //Correspondace names to tiles
	charSize Vector2                   //Space, given for a single tile on screen
}

// Prerender tile
func PrerenderTile(r string, color sdl.Color) *sdl.Surface {
	var text *sdl.Surface
	text, err := Graphic.font.RenderUTF8Blended(string(r), color)
	if err != nil {
		return nil
	}
	return text
}

// Updates users window view
func UpdateGameScreen() {
	var text *sdl.Surface
	Graphic.surface.FillRect(nil, sdl.MapRGB(Graphic.surface.Format, 0, 0, 0))
	for y := 0; y < screenAreaHeight; y++ {
		for x := 0; x < screenAreaWidth; x++ {
			if GameData.screenArea[x][y].wall.wtype != 0 {
				text = Graphic.tiles[GameData.wallTypes[GameData.screenArea[x][y].wall.wtype].tile]
			} else if GameData.screenArea[x][y].mob != 0 {
				text = Graphic.tiles[GameData.mobTypes[GameData.mobs[GameData.screenArea[x][y].mob].mtype].tile]
			} else {
				text = Graphic.tiles[GameData.floorTypes[GameData.screenArea[x][y].floor.ftype].tile]
			}
			text.Blit(nil, Graphic.surface, &sdl.Rect{X: int32(x * Graphic.charSize.x), Y: int32(y * Graphic.charSize.y), W: 0, H: 0})
		}
	}
	UIDisplayStats()
	Graphic.window.UpdateSurface()
}
