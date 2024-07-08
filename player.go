package main

//Player movement
func MovePlayer(v Vector3) {
	ProceedTime()
	AttemptMove(v)
}

//Check for walls before moving
func AttemptMove(v Vector3) {
	if !OutOfBounds(NewV3(GameData.camera.x+v.x, GameData.camera.y+v.y, GameData.camera.z+v.z)) && GameData.world[GameData.camera.x+v.x][GameData.camera.y+v.y][0].mob == 0 && GameData.world[GameData.camera.x+v.x][GameData.camera.y+v.y][0].wall.wtype == 0 {
		GameData.camera = NewV3(GameData.camera.x+v.x, GameData.camera.y+v.y, GameData.camera.z+v.z)
		GameData.mobs[1].AttemptMove(v)
	}
}
