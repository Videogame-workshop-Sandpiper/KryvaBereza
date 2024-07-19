package main

import "math/rand"

const (
	worldSize        int = 500 //Width and length of world
	worldHeight      int = 12  //Height of world
	screenAreaWidth  int = 100 //Width of game screen
	screenAreaHeight int = 30  //Height of game screen
)

var GameData struct {
	world      [worldSize][worldSize][worldHeight]Tile //World as a matrix
	screenArea [screenAreaWidth][screenAreaHeight]Tile //Tiles, visible by camera
	mobs       [10]Mob                                 //Array of mobs
	//Arrays of types of tile elements
	wallTypes  [20]WallType
	mobTypes   [20]MobType
	floorTypes [20]FloorType
	//Maps, needed for a quicker access to a type
	wallTypeMap  map[string]WallType
	mobTypeMap   map[string]MobType
	floorTypeMap map[string]FloorType

	worldSeed int     //Random world generation seed
	camera    Vector3 //Position of camera
	gameTime  int     //Ingame timer
	cameraDir int     //Camera direction. North, east, south, west
}

//World measurment unit
type Tile struct {
	floor Floor
	mob   int //Mob's index in an array
	wall  Wall
}

//I have made separate Init functions in case i want to have some script playing when object is being initialized
//Tile initialization function
func (t *Tile) Init(floor Floor, mob int, wall Wall) {
	t.floor = floor
	t.mob = mob
	t.wall = wall
}

//Wall parameters
type WallType struct {
	name    string
	tile    int
	tileTop int //floor type, that will appear on top of said wall due=ring wall generation
	index   int //index in type array
}

//Wall data
type Wall struct {
	wtype WallType
	pos   Vector3
	onTop bool //Wheteher or not this wall is on top of another wall. No implementation yet
}

//Wall init function
func (w *Wall) Init(wtype WallType, pos Vector3) {
	w.wtype = wtype
	w.pos = pos
	//Set if this wall is on top of another wall
	if (pos.z+1) < worldHeight && GameData.world[pos.x][pos.y][pos.z+1].wall.wtype.index == 0 {
		w.onTop = true
	}
	//Set if the wall below this is on top
	if (pos.z-1) >= 0 && GameData.world[pos.x][pos.y][pos.z-1].wall.wtype.index != 0 {
		GameData.world[pos.x][pos.y][pos.z-1].wall.onTop = false
	}
	//There always should be floor over wall
	if (pos.z+1) < worldHeight && GameData.world[pos.x][pos.y][pos.z+1].floor.ftype.index == 0 {
		GameData.world[pos.x][pos.y][pos.z+1].floor.Init(GameData.floorTypes[wtype.tileTop], NewV3(pos.x, pos.y, pos.z+1))
	}
}

//Floor parameters
type FloorType struct {
	name  string
	tile  int
	index int //index in type array
}

//Floor data
type Floor struct {
	ftype     FloorType
	pos       Vector3
	underWall bool //Is this floor located under a wall. No implementation yet
}

func (f *Floor) Init(ftype FloorType, pos Vector3) {
	f.ftype = ftype
	f.pos = pos
	//Checks if it's under a wall
	if GameData.world[pos.x][pos.y][pos.z].wall.wtype.index != 0 {
		f.underWall = true
	}
}

//Creature parameters
type MobType struct {
	name   string
	health int
	tile   int
	index  int //index in type array
}

//Creature data
type Mob struct {
	mtype  MobType
	health int
	pos    Vector3
	index  int //index in array
}

//Mob init function
func (m *Mob) Init(mtype MobType, pos Vector3, index int) {
	m.mtype = mtype
	m.health = GameData.mobTypes[mtype.index].health
	m.pos = pos
	m.index = index

	GameData.world[pos.x][pos.y][pos.z].mob = index
}

// Attempt moving
func (m *Mob) AttemptMove(v Vector3) {
	//Applies gravity if no floor underneath
	if (!OutOfBounds(NewV3(m.pos.x+v.x, m.pos.y+v.y, m.pos.z+v.z)) && GameData.world[m.pos.x+v.x][m.pos.y+v.y][m.pos.z].floor.ftype.index == 0) || GameData.world[m.pos.x][m.pos.y][m.pos.z].floor.ftype.index == 0 {
		v.z = -1
	}
	//Nullifies vectors if there are obstacles
	if OutOfBounds(NewV3(m.pos.x+v.x, m.pos.y+v.y, m.pos.z)) || GameData.world[m.pos.x+v.x][m.pos.y+v.y][m.pos.z].mob != 0 || GameData.world[m.pos.x+v.x][m.pos.y+v.y][m.pos.z].wall.wtype.index != 0 {
		v = NewV3(0, 0, v.z)
	}
	//Moves mob if neither of vectors are zero
	if !(v.x == 0 && v.y == 0 && v.z == 0) {
		GameData.world[m.pos.x][m.pos.y][m.pos.z].mob = 0
		m.pos = NewV3(m.pos.x+v.x, m.pos.y+v.y, m.pos.z+v.z)
		GameData.world[m.pos.x][m.pos.y][m.pos.z].mob = m.index
	}
}

//Fills screen area with tiles, currently in camera's view
func UpdateScreenArea() {
	for y := 0; y < screenAreaHeight; y++ {
		for x := 0; x < screenAreaWidth; x++ {
			//Finds highest tile to be drawn in that part of the screen
			var Y int = y - GameData.mobs[1].pos.z
			var index int = worldHeight - 1
			for index > 0 && (OutOfBounds(WorldPosition(NewV2(x, Y+index))) || TileEmpty(NewV3(WorldPosition(NewV2(x, Y+index)).x, WorldPosition(NewV2(x, Y+index)).y, index))) {
				index--
			}
			//Draws barrier if tile appears to be out of screen
			if !OutOfBounds(WorldPosition(NewV2(x, Y+index))) {
				GameData.screenArea[x][y] = GameData.world[WorldPosition(NewV2(x, Y+index)).x][WorldPosition(NewV2(x, Y+index)).y][index]
			} else {
				GameData.screenArea[x][y].wall.Init(GameData.wallTypeMap["barrier"], NewV3(x, y, 0))
			}
		}
	}
}

//Returns screen position in a world depending on a camera rotation
func WorldPosition(v Vector2) Vector3 {
	var offsetx int = ((screenAreaWidth - screenAreaHeight) / 2)
	var offsety int = (screenAreaHeight - ((screenAreaWidth - screenAreaHeight) / 2))
	switch GameData.cameraDir {
	case 1:
		return NewV3(GameData.camera.x-(screenAreaWidth/2)+v.y+offsetx, GameData.camera.y+(screenAreaHeight/2)+(screenAreaHeight-v.x)-offsety, 0)
	case 2:
		return NewV3(GameData.camera.x-(screenAreaWidth/2)+(screenAreaWidth-v.x), GameData.camera.y-(screenAreaHeight/2)+(screenAreaHeight-v.y), 0)
	case 3:
		return NewV3(GameData.camera.x+(screenAreaWidth/2)-v.y-offsetx, GameData.camera.y-(screenAreaHeight/2)+v.x-screenAreaHeight+offsety, 0)
	default:
		return NewV3(GameData.camera.x-(screenAreaWidth/2)+v.x, GameData.camera.y-(screenAreaHeight/2)+v.y, 0)
	}
}

//Checks if coordinates are out of worlds bounds
func OutOfBounds(v Vector3) bool {
	if v.x >= 0 && v.x < worldSize && v.y >= 0 && v.y < worldSize && v.z >= 0 && v.z < worldHeight {
		return false
	} else {
		return true
	}
}

//Check if tile is empty
func TileEmpty(v Vector3) bool {
	if GameData.world[v.x][v.y][v.z].floor.ftype.index == 0 && GameData.world[v.x][v.y][v.z].wall.wtype.index == 0 && GameData.world[v.x][v.y][v.z].mob == 0 {
		return true
	} else {
		return false
	}
}

//Increment ingame timer and do all time-related actions
func ProceedTime() {
	GameData.gameTime++
	//Creature behaviour
	for i := range GameData.mobs {
		if GameData.mobs[i].health != 0 && GameData.mobs[i].mtype.name != "player" {
			GameData.mobs[i].AttemptMove(NewV3(1-rand.Intn(3), 1-rand.Intn(3), 0))
		}
	}
}
