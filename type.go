package cheatgo

import "math"

// Matrix 是一个4x4的浮点数矩阵
type Matrix [4][4]float32

// Vector3 是一个三维向量
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// Add 方法用于向量加法
func (v *Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

// Sub 方法用于向量减法
func (v *Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

// Mul 方法用于向量乘法
func (v *Vector3) Mul(v2 Vector3) Vector3 {
	return Vector3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

// Div 方法用于向量除法
func (v *Vector3) Div(v2 Vector3) Vector3 {
	return Vector3{v.X / v2.X, v.Y / v2.Y, v.Z / v2.Z}
}

// Dot 方法用于计算点积
func (v *Vector3) Dot(v2 Vector3) float32 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Dist 方法用于计算两点之间的距离
func (v *Vector3) Dist(v2 Vector3) float32 {
	return float32(math.Abs(float64(v.X-v2.X)) + math.Abs(float64(v.Y-v2.Y)) + math.Abs(float64(v.Z-v2.Z)))
}

// Vector2 二维向量
type Vector2 struct {
	X float32
	Y float32
}

// Add 方法用于向量加法
func (v *Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{v.X + v2.X, v.Y + v2.Y}
}

// Sub 方法用于向量减法
func (v *Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{v.X - v2.X, v.Y - v2.Y}
}

// Mul 方法用于向量乘法
func (v *Vector2) Mul(v2 Vector2) Vector2 {
	return Vector2{v.X * v2.X, v.Y * v2.Y}
}

// Div 方法用于向量除法
func (v *Vector2) Div(v2 Vector2) Vector2 {
	return Vector2{v.X / v2.X, v.Y / v2.Y}
}

// Dot 方法用于计算点积
func (v *Vector2) Dot(v2 Vector2) float32 {
	return v.X*v2.X + v.Y*v2.Y
}

// Dist 方法用于计算两点之间的距离
func (v *Vector2) Dist(v2 Vector2) float32 {
	return float32(math.Abs(float64(v.X-v2.X)) + math.Abs(float64(v.Y-v2.Y)))
}
