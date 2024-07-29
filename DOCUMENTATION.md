# Global Variables
|       Name         |        Description         |
| ------------------ | -------------------------- |
| `FPS_CAP`          | FPS cap                    |
| `worldSize`        | Width and length of world  |
| `worldHeight`      | Height of world            |
| `screenAreaWidth`  | Width of game screen       |
| `screenAreaHeight` | Height of game screen      |
| `fontPath`         | Path to ttf file           |
| `fontSize`         | Font size                  |
| `tileSetSize`      | Size of tileset            |

# Structs
### `GameData`
```go 
world      [worldSize][worldSize][worldHeight]Tile  //World as a matrix
screenArea [screenAreaWidth][screenAreaHeight]Tile  //Tiles, visible by camera
mobs       [10]Mob                                  //Array of mobs

wallTypes  [20]WallType
mobTypes   [20]MobType                              //Arrays of types of tile elements
floorTypes [20]FloorType

wallTypeMap  map[string]WallType
mobTypeMap   map[string]MobType                     //Maps, needed for a quicker access to a type
floorTypeMap map[string]FloorType
worldSeed int                                       //Random world generation seed
camera    Vector3                                   //Position of camera
gameTime  int                                       //Ingame timer
cameraDir int                                       //Camera direction. North, east, south, west
```
### `Tile` World measurment unit
```go 
floor Floor
mob   int                                           //Mob's index in an Gamedata.mobTypes
wall  Wall

func Init(floor Floor, mob int, wall Wall) (t *Tile) {
	t.floor = floor
	t.mob = mob
	t.wall = wall
}
```
### `WallType` Wall parameters
```go 
name    string                                      //Material
tile    int                                         //TODO
tileTop int                                         //Floor type, that will appear on top of said wall during wall generation
index   int                                         //Index in GameData.wallTypes 
```
### `Wall`
```go 
wtype WallType                                      //Wall parameters
pos   Vector3                                       //Position
onTop bool                                          //Wheteher or not this wall is on top of another wall. No implementation yet

func Init(wtype WallType, pos Vector3) (w *Wall) {
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
```
### `FloorType` Floor parameters
```go 
name  string                                        //Material
tile  int                                           //TODO
index int                                           //Index in GameData.floorTypes 
```
### `Floor`
```go 
ftype     FloorType                                 //Floor parameters
pos       Vector3                                   //Position
underWall bool                                      //Is this floor located under a wall. No implementation yet

func Init(ftype FloorType, pos Vector3) (f *Floor) {
	f.ftype = ftype
	f.pos = pos
	//Checks if it's under a wall
	if GameData.world[pos.x][pos.y][pos.z].wall.wtype.index != 0 {
		f.underWall = true
	}
}
```
### `MobType` Creature parameters
```go 
name   string                                       //Name of the mob
health int                                          //Selfexplanatory
tile   int                                          //TODO
index  int                                          //Index in GameData.mobTypes 
```
### `Mob`
```go 
mtype  MobType                                      //Creature parameters
health int                                          //AGAIN
pos    Vector3                                      //Position
index  int                                          //AGAIN

func Init(mtype MobType, pos Vector3, index int) (m *Mob)  {
	m.mtype = mtype
	m.health = GameData.mobTypes[mtype.index].health
	m.pos = pos
	m.index = index

	GameData.world[pos.x][pos.y][pos.z].mob = index
}

func AttemptMove(v Vector3) (m *Mob)  {
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
```
### `Graphic`
```go 
window  *sdl.Window                                 //Game window
font    *ttf.Font                                   //Opened font
surface *sdl.Surface                                //Pixel map of the window
tiles    [tileSetSize]*sdl.Surface                  //Set of tiles
tileMap  map[string]int                             //Correspondence of names to tiles
charSize Vector2                                    //Space given for a single tile on screen
```
### `Vector2`
```go 
x int                                               //Selfexplanatory
y int                                               //Selfexplanatory

// Set vector's coordinates
func (v *Vector2) Init(X int, Y int) {
	v.x = X
	v.y = Y
}

// Create vector in one line
func NewV2(X int, Y int) Vector2 {
	var v Vector2
	v.Init(X, Y)                                    //AGAIN
	return v
}
```
### `Vector3`
```go 
x int                                               //Selfexplanatory
y int                                               //Selfexplanatory
z int                                               //Selfexplanatory

// Set vector's coordinates
func (v *Vector3) Init(X int, Y int, Z int) {
	v.x = X
	v.y = Y
	v.z = Z
}

// Create vector in one line
func NewV3(X int, Y int, Z int) Vector3 {
	var v Vector3
	v.Init(X, Y, Z)                                 //AGAIN
	return v
}
```
# Functions
```go
//dataload.go
func LoadDataFiles()															// Wrapper to load all data
func CreateTiles()																// Renderes all tiles from data/tiles.txt for future use
func CreateWallTypes()															// Creates wall types from data/walls.txt
func CreateFloorTypes()															// Creates floor types from data/floors.txt
func CreateMobTypes()															// Creates mob types from data/mobs.txt

//gamedata.go
func Init(floor Floor, mob int, wall Wall) (t *Tile)							// Constructor for Tyle struct
func Init(wtype WallType, pos Vector3) (w *Wall)								// Constructor for Wall struct
func Init(ftype FloorType, pos Vector3) (f *Floor)								// Constructor for Floor struct
func Init(mtype MobType, pos Vector3, index int) (m *Mob) 						// Constructor for Mob struct
func AttemptMove(v Vector3) (m *Mob) 											// Moves the mob if possible
func UpdateScreenArea()															// Fills screen area with tiles, currently in camera's view
func WorldPosition(v Vector2) Vector3											// Returns screen position in a world depending on a camera rotation
func OutOfBounds(v Vector3) bool												// Checks if the position is out of bounds
func TileEmpty(v Vector3) bool													// Checks if given tile is empty (basically air)
func ProceedTime()																// Increment ingame timer and do all time-related actions

//game.go
func run() (err error)															// Init all services, make pre-game calculations and init the main cycle
func main()																		// Program entry point

//graphic.go
func PrerenderTile(r string, color sdl.Color, bgcolor sdl.Color) *sdl.Surface	// Renderes SDL surface from given parameters
func PrerenderTileBlended(r string, color sdl.Color) *sdl.Surface				// Renderes colored SDL surface from given parameters
func UpdateGameScreen()															// Draws the next frame

//math.go
func StrToInt8(i string) uint8													// Converter
func StrToInt(i string) int														// Converter

//player.go
func MovePlayer(v Vector3)														// Move player and proceed time
func AttemptMove(v Vector3)														// Move player and camera if possible
func RotateFromCamera(v Vector3) Vector3										// Rotates player movement, according to the camera direction

//ui.go
func UIDisplayStats()															// Display stats in the right top corner
func UIDisplaySymbol(a string, col int, row int)								// Display a single symbol in the screen
func UIDisplayText(a string, row int)											// Display series of symbols on screen
func UIDrawCompass(row int)														// Draw a compass in the right bottom corner

//worldgen.go
func BuildWall(v Vector3, height int, material string)							// Buiild a wall with given height
func BuildTree(v Vector3, height int)											// Build a tree
func GenerateWorld()															// Generate the world
func ClearWorld()																// Wipe the world of the face of the earth
func FillWorld()																// Fills world with tiles
func SpawnMobs()																// Spawn mobs
```