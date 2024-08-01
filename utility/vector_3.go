package utility

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3) AlmostZero() bool {
	const epsilon = 1e-6
	return math.Abs(float64(v.X)) < epsilon && math.Abs(float64(v.Y)) < epsilon && math.Abs(float64(v.Z)) < epsilon
}

func (v Vector3) ToArray() [3]float64 {
	return [3]float64{v.X, v.Y, v.Z}
}

func (v Vector3) ToSlice() []float64 {
	return []float64{v.X, v.Y, v.Z}
}

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vector3) Mul(scalar float64) Vector3 {
	return Vector3{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

// 外積
//
//	V = Vector{ y1*z2 - z1*y2, z1*x2 - x1*z2, x1*y2 - y1*x2 }
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

// 內積(點積)
//
//	dot = x1*x2 + y1*y2 + z1*z2
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// ║v║ = Sqrt(x*x + y*y + z*z)
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// 向量正規化 (v -> V)
//
//	V = v/║v║
//	V = Vector3{ x/║v║, y/║v║, z/║v║ }，║v║ != 0
func (v Vector3) Normalize() Vector3 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector3{X: 0, Y: 0, Z: 0}
	}
	return Vector3{
		X: v.X / mag,
		Y: v.Y / mag,
		Z: v.Z / mag,
	}
}

// 取得兩點間的距離
//
//	distance = Sqrt((x2 - x1)^2 + (y2 - y1)^2 + (z2 - z1)^2)
func (v Vector3) Distance(other Vector3) float64 {
	return math.Sqrt(math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2) + math.Pow(v.Z-other.Z, 2))
}
