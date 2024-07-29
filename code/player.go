package main

//Player movement
func MovePlayer(v Vector3) {
	ProceedTime()
	AttemptMove(v)
}

//Move player and camera
func AttemptMove(v Vector3) {
	GameData.mobs[1].AttemptMove(RotateFromCamera(v))
	GameData.camera = GameData.mobs[1].pos
}

//Rotates player movement, according to the camera direction
func RotateFromCamera(v Vector3) Vector3 {
	switch GameData.cameraDir {
	case 1:
		return NewV3(v.y, -v.x, v.z)
	case 2:
		return NewV3(-v.x, -v.y, v.z)
	case 3:
		return NewV3(-v.y, v.x, v.z)
	default:
		return v
	}
}
