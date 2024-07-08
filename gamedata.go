package main

import "math/rand"

const (
	worldSize        int = 500
	worldHeight      int = 10
	screenAreaWidth  int = 91
	screenAreaHeight int = 41
)

var GameData struct {
	world      [worldSize][worldSize][worldHeight]Tile //World as a matrix
	screenArea [screenAreaWidth][screenAreaHeight]Tile //Tiles, visible by camera
	mobs       [10]Mob

	wallTypes  [20]WallType
	mobTypes   [20]MobType
	floorTypes [20]FloorType

	wallTypeMap  map[string]WallType
	mobTypeMap   map[string]MobType
	floorTypeMap map[string]FloorType

	worldSeed int     //Random world generation seed
	camera    Vector3 //Position of camera
	gameTime  int     //Ingame timer
}

//World measurment unit
type Tile struct {
	floor Floor
	mob   int
	wall  Wall
}

func (t *Tile) Init(floor Floor, mob int, wall Wall) {
	t.floor = floor
	t.mob = mob
	t.wall = wall
}

type WallType struct {
	name    string
	tile    int
	tileTop int
	index   int
}

type Wall struct {
	wtype int
	pos   Vector3
	onTop bool
}

func (w *Wall) Init(wtype int, pos Vector3) {
	w.wtype = wtype
	w.pos = pos
	if (pos.z+1) < worldHeight && GameData.world[pos.x][pos.y][pos.z+1].wall.wtype == 0 {
		w.onTop = true
	}
	if (pos.z-1) >= 0 && GameData.world[pos.x][pos.y][pos.z-1].wall.wtype != 0 {
		GameData.world[pos.x][pos.y][pos.z-1].wall.onTop = false
	}
	if (pos.z+1) < worldHeight && GameData.world[pos.x][pos.y][pos.z+1].floor.ftype == 0 {
		GameData.world[pos.x][pos.y][pos.z+1].floor.Init(GameData.floorTypes[GameData.wallTypes[w.wtype].tileTop].index, NewV3(pos.x, pos.y, pos.z+1))
	}
}

type FloorType struct {
	name  string
	tile  int
	index int
}

type Floor struct {
	ftype     int
	pos       Vector3
	underWall bool
}

func (f *Floor) Init(ftype int, pos Vector3) {
	f.ftype = ftype
	f.pos = pos
	if GameData.world[pos.x][pos.y][pos.z].wall.wtype != 0 {
		f.underWall = true
	}
}

type MobType struct {
	name   string
	health int
	tile   int
	index  int
}

type Mob struct {
	mtype  int
	health int
	pos    Vector3
	index  int
}

func (m *Mob) Init(mtype int, pos Vector3, index int) {
	m.mtype = mtype
	m.health = GameData.mobTypes[mtype].health
	m.pos = pos
	m.index = index
}

// Check for walls before moving
func (m *Mob) AttemptMove(v Vector3) {
	if !OutOfBounds(NewV3(m.pos.x+v.x, m.pos.y+v.y, m.pos.z+v.z)) && GameData.world[m.pos.x+v.x][m.pos.y+v.y][0].mob == 0 && GameData.world[m.pos.x+v.x][m.pos.y+v.y][0].wall.wtype == 0 {
		GameData.world[m.pos.x][m.pos.y][m.pos.z].mob = 0
		m.pos = NewV3(m.pos.x+v.x, m.pos.y+v.y, m.pos.z+v.z)
		GameData.world[m.pos.x][m.pos.y][m.pos.z].mob = m.index
	}
}

func BuildWall(v Vector3, height int) {
	for i := range height {
		GameData.world[v.x][v.y][v.z+i].wall.Init(GameData.wallTypeMap["concrete_wall"].index, NewV3(v.x, v.y, v.z+i))
	}
}

//Fills world with tiles
func FillWorld() {
	rand.Seed(int64(GameData.worldSeed))
	for y := 0; y < worldSize; y++ {
		for x := 0; x < worldSize; x++ {
			if rand.Intn(5) == 1 {
				GameData.world[x][y][0].floor.Init(GameData.floorTypeMap["path"].index, NewV3(x, y, 0))
			} else {
				GameData.world[x][y][0].floor.Init(GameData.floorTypeMap["grass"].index, NewV3(x, y, 0))
			}
			GameData.world[x][y][0].mob = Graphic.tileMap["noth"]
			if rand.Intn(121) == 1 {
				BuildWall(NewV3(x, y, 0), 3)
			} else {
				GameData.world[x][y][0].wall.Init(0, NewV3(0, 0, 0))
			}
			if x == 400 || y == 220 {
				if rand.Intn(20) == 1 {
					BuildWall(NewV3(x, y, 2), 3)
				} else {
					BuildWall(NewV3(x, y, 0), 5)
				}
			}
		}
	}

	GameData.camera.Init(330, 320, 0)
	GameData.mobs[1].Init(GameData.mobTypeMap["player"].index, NewV3(330, 320, 0), 1)
	GameData.mobs[2].Init(GameData.mobTypeMap["imp"].index, NewV3(330, 330, 0), 2)
}

//Fills screen area with tiles, currently in camera's view
func UpdateScreenArea() {
	for y := 0; y < screenAreaHeight; y++ {
		for x := 0; x < screenAreaWidth; x++ {
			if !OutOfBounds(WorldPosition(NewV2(x, y))) {
				var index int = worldHeight - 1
				for index > 0 && (OutOfBounds(WorldPosition(NewV2(x, y+index))) || TileEmpty(NewV3(WorldPosition(NewV2(x, y+index)).x, WorldPosition(NewV2(x, y+index)).y, index))) {
					index--
				}
				GameData.screenArea[x][y] = GameData.world[WorldPosition(NewV2(x, y+index)).x][WorldPosition(NewV2(x, y+index)).y][index]
			} else {
				GameData.screenArea[x][y].wall.Init(GameData.wallTypeMap["barrier"].index, NewV3(x, y, 0))
			}
		}
	}
}

//Returns screen position in a world
func WorldPosition(v Vector2) Vector3 {
	return NewV3(GameData.camera.x-(screenAreaWidth/2)+v.x, GameData.camera.y-(screenAreaHeight/2)+v.y, 0)
}

//Checks if coordinates are out of worlds bounds
func OutOfBounds(v Vector3) bool {
	if v.x >= 0 && v.x < worldSize && v.y >= 0 && v.y < worldSize && v.z >= 0 && v.z < worldHeight {
		return false
	} else {
		return true
	}
}

func TileEmpty(v Vector3) bool {
	if GameData.world[v.x][v.y][v.z].floor.ftype == 0 && GameData.world[v.x][v.y][v.z].wall.wtype == 0 && GameData.world[v.x][v.y][v.z].mob == 0 {
		return true
	} else {
		return false
	}
}

func ProceedTime() {
	GameData.gameTime++
	for i := range GameData.mobs {
		if GameData.mobs[i].health != 0 && GameData.mobTypes[GameData.mobs[i].mtype].name != "player" {
			GameData.mobs[i].AttemptMove(NewV3(1-rand.Intn(3), 1-rand.Intn(3), 0))
		}
	}
}
