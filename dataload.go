package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// Loads all the necessary data
func LoadDataFiles() {
	CreateTiles()
	CreateFloorTypes()
	CreateWallTypes()
	CreateMobTypes()
}

// Loads tileset
func CreateTiles() {
	Graphic.tileMap = make(map[string]int)

	Graphic.tiles[0] = PrerenderTile(" ", sdl.Color{R: 0, G: 0, B: 0}, sdl.Color{R: 0, G: 0, B: 0})
	Graphic.tileMap["noth"] = 0
	file, err := os.Open("data/tiles.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i int = 1
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		//Checks if tile has background color
		if len(line) == 5 {
			Graphic.tiles[i] = PrerenderTile(line[1], sdl.Color{R: StrToInt8(line[2]), G: StrToInt8(line[3]), B: StrToInt8(line[4])}, sdl.Color{R: 0, G: 0, B: 0})
		} else {
			Graphic.tiles[i] = PrerenderTile(line[1], sdl.Color{R: StrToInt8(line[2]), G: StrToInt8(line[3]), B: StrToInt8(line[4])}, sdl.Color{R: StrToInt8(line[5]), G: StrToInt8(line[6]), B: StrToInt8(line[7])})
		}
		Graphic.tileMap[line[0]] = i
		i++
	}
}

// Loads wall data
func CreateWallTypes() {
	GameData.wallTypeMap = make(map[string]WallType)

	file, err := os.Open("data/walls.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i int = 1
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		GameData.wallTypes[i].name = line[1]
		GameData.wallTypes[i].tile = Graphic.tileMap[line[2]]
		GameData.wallTypes[i].tileTop = GameData.floorTypeMap[line[3]]
		GameData.wallTypeMap[line[1]] = GameData.wallTypes[i]
		i++
	}
}

// Loads floor data
func CreateFloorTypes() {
	GameData.floorTypeMap = make(map[string]FloorType)

	file, err := os.Open("data/floors.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i int = 1
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		GameData.floorTypes[i].name = line[1]
		GameData.floorTypes[i].tile = Graphic.tileMap[line[2]]
		GameData.floorTypeMap[line[1]] = GameData.floorTypes[i]
		i++
	}
}

// Loads creature data
func CreateMobTypes() {
	GameData.mobTypeMap = make(map[string]MobType)

	file, err := os.Open("data/mobs.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i int = 1
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		GameData.mobTypes[i].name = line[1]
		GameData.mobTypes[i].tile = Graphic.tileMap[line[3]]
		GameData.mobTypes[i].health = StrToInt(line[2])
		GameData.mobTypeMap[line[1]] = GameData.mobTypes[i]
		i++
	}
}
