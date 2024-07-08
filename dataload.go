package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

func LoadDataFiles() {
	CreateTiles()
	CreateFloorTypes()
	CreateWallTypes()
	CreateMobTypes()
}

// Creates a tileset
func CreateTiles() {
	Graphic.tileMap = make(map[string]int)

	Graphic.tiles[0] = PrerenderTile(" ", sdl.Color{R: 0, G: 0, B: 0})
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
		Graphic.tiles[i] = PrerenderTile(line[1], sdl.Color{R: StrToInt8(line[2]), G: StrToInt8(line[3]), B: StrToInt8(line[4])})
		Graphic.tileMap[line[0]] = i
		i++
	}
}

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
		GameData.wallTypes[i].tileTop = GameData.floorTypeMap[line[3]].index
		GameData.wallTypes[i].index = i
		GameData.wallTypeMap[line[1]] = GameData.wallTypes[i]
		i++
	}
}

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
		GameData.floorTypes[i].index = i
		GameData.floorTypeMap[line[1]] = GameData.floorTypes[i]
		i++
	}
}

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
		GameData.mobTypes[i].index = i
		GameData.mobTypeMap[line[1]] = GameData.mobTypes[i]
		i++
	}
}
