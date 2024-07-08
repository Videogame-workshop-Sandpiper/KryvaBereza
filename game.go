package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Main cycle
func run() (err error) {

	if err = ttf.Init(); err != nil {
		return
	}
	defer ttf.Quit()

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		return
	}
	defer sdl.Quit()

	if Graphic.window, err = sdl.CreateWindow("Loading...", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1280, 720, sdl.WINDOW_SHOWN); err != nil {
		return
	}
	defer Graphic.window.Destroy()

	if Graphic.surface, err = Graphic.window.GetSurface(); err != nil {
		return
	}

	if Graphic.font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
		return
	}
	defer Graphic.font.Close()
	running := true
	var tick = 0
	//windowWidth, windowHeight := Graphic.window.GetSize()
	Graphic.charSize.x, Graphic.charSize.y, _ = Graphic.font.SizeUTF8("a")
	Graphic.charSize.y -= 5
	Graphic.charSize.y -= 1
	//rows = int(windowHeight) / Graphic.charSize.y
	//cols = int(windowWidth) / Graphic.charSize.x

	LoadDataFiles()
	GameData.worldSeed = rand.Int()
	FillWorld()

	for running {
		if tick == 8 {
			tick = 0
		} else {
			tick++
		}
		UpdateScreenArea()
		var keyst = sdl.GetKeyboardState()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		start := time.Now()
		if tick == 0 {
			if keyst[sdl.SCANCODE_RIGHT] != 0 {
				MovePlayer(NewV3(1, 0, 0))
			}
			if keyst[sdl.SCANCODE_LEFT] != 0 {
				MovePlayer(NewV3(-1, 0, 0))
			}
			if keyst[sdl.SCANCODE_UP] != 0 {
				MovePlayer(NewV3(0, -1, 0))
			}
			if keyst[sdl.SCANCODE_DOWN] != 0 {
				MovePlayer(NewV3(0, 1, 0))
			}
			if keyst[sdl.SCANCODE_SPACE] != 0 {
				ProceedTime()
			}
		}
		UpdateGameScreen()

		fps := 1000000 / time.Since(start).Microseconds()
		Graphic.window.SetTitle("Kryva Bereza vPre-anything FPS:" + strconv.FormatInt(int64(int(fps)), 10))

	}

	return
}

// Main function
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
