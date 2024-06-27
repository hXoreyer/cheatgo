package cheatgo

// WordToScreen 将位置转换为屏幕坐标
// 参数：viewMatrix 视图矩阵，position 位置向量
// 返回值：屏幕X坐标，屏幕Y坐标
func WordToScreen(viewMatrix Matrix, position Vector3) (float32, float32) {
	var screenX float32
	var screenY float32
	screenX = viewMatrix[0][0]*position.X + viewMatrix[0][1]*position.Y + viewMatrix[0][2]*position.Z + viewMatrix[0][3]
	screenY = viewMatrix[1][0]*position.X + viewMatrix[1][1]*position.Y + viewMatrix[1][2]*position.Z + viewMatrix[1][3]
	w := viewMatrix[3][0]*position.X + viewMatrix[3][1]*position.Y + viewMatrix[3][2]*position.Z + viewMatrix[3][3]
	if w < 0.01 {
		return -1, -1
	}
	invw := 1.0 / w
	screenX *= invw
	screenY *= invw
	width, _, _ := getSystemMetrics.Call(0)
	height, _, _ := getSystemMetrics.Call(1)
	widthFloat := float32(width)
	heightFloat := float32(height)
	x := widthFloat / 2
	y := heightFloat / 2
	x += 0.5*screenX*widthFloat + 0.5
	y -= 0.5*screenY*heightFloat + 0.5
	return x, y
}

// GetScreenSize 获取屏幕尺寸
// 返回值：屏幕宽度，屏幕高度
func GetScreenSize() (uintptr, uintptr) {
	screenWidth, _, _ := getSystemMetrics.Call(0)
	screenHeight, _, _ := getSystemMetrics.Call(1)
	return screenWidth, screenHeight
}
