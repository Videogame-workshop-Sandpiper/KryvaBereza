package main

import "math/rand"

//Build a 1x1 collumn of wall
func BuildWall(v Vector3, height int, material string) {
	for i := range height {
		GameData.world[v.x][v.y][v.z+i].wall.Init(GameData.wallTypeMap[material], NewV3(v.x, v.y, v.z+i))
	}
}

//Build a tree
func BuildTree(v Vector3, height int) {
	//Grow stump
	BuildWall(v, height, "tree_stump")
	//Surround stump by foliage
	for i := -3; i <= 3; i++ {
		for j := -3; j <= 3; j++ {
			for k := -1; k <= 2; k++ {
				if !OutOfBounds(NewV3(v.x+i, v.y+j, v.z-k+height)) && TileEmpty(NewV3(v.x+i, v.y+j, v.z-k+height)) && rand.Intn(5) != 1 {
					GameData.world[v.x+i][v.y+j][v.z-k+height].wall.Init(GameData.wallTypeMap["leaves"], NewV3(v.x+i, v.y+j, v.z-k+height))
				}
			}
		}
	}
}

//Regenerate world
func GenerateWorld() {
	ClearWorld()
	GameData.worldSeed = rand.Int()
	FillWorld()
}

//Clear all world data
func ClearWorld() {
	for i := 0; i < worldSize; i++ {
		for j := 0; j < worldSize; j++ {
			for k := 0; k < worldHeight; k++ {
				GameData.world[i][j][k] = Tile{}
			}
		}
	}
	GameData.mobs = [10]Mob{}
}

//Fills world with tiles
func FillWorld() {
	//Set seed
	rand.Seed(int64(GameData.worldSeed))
	for y := 0; y < worldSize; y++ {
		for x := 0; x < worldSize; x++ {
			//Place floor
			if rand.Intn(5) == 1 {
				GameData.world[x][y][0].floor.Init(GameData.floorTypeMap["path"], NewV3(x, y, 0))
			} else {
				GameData.world[x][y][0].floor.Init(GameData.floorTypeMap["grass"], NewV3(x, y, 0))
			}
			//Build trees
			if rand.Intn(321) == 1 {
				BuildTree(NewV3(x, y, 0), 2+rand.Intn(6))
			}
			//Build The Wall
			if x == 250 || y == 250 {
				if rand.Intn(20) == 1 {
					BuildWall(NewV3(x, y, 2), 3, "concrete_wall")
				} else {
					BuildWall(NewV3(x, y, 0), 5, "concrete_wall")
				}
			}
		}
	}
	SpawnMobs()
}

func SpawnMobs() {
	//Spawn player and imp
	GameData.mobs[1].Init(GameData.mobTypeMap["player"], NewV3(250, 250, 5), 1)
	GameData.mobs[2].Init(GameData.mobTypeMap["imp"], NewV3(330, 330, 0), 2)
	GameData.camera = GameData.mobs[1].pos
}
