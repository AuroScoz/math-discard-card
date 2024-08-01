package utility

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) ToArray() [2]float64 {
	return [2]float64{v.X, v.Y}
}
func (v Vector2) ToSlice() []float64 {
	return []float64{v.X, v.Y}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector2) Mul(scalar float64) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

// 外積
//
//	cross = x1*y2 - y1*x2
func (v Vector2) Cross(other Vector2) float64 {
	return v.X*other.Y - v.Y*other.X
}

// 內積(點積)
//
//	dot = x1*x2 + y1*y2
func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// ║v║ = Sqrt(x*x + y*y)
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 向量正規化 (v -> V)
//
//	V = v/║v║
//	V = Vector2{ x/║v║, y/║v║ }，║v║ != 0
func (v Vector2) Normalize() Vector2 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2{X: 0, Y: 0}
	}
	return Vector2{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

// 取得兩點間的距離
//
//	distance = Sqrt((x2 - x1)^2 + (y2 - y1)^2)
func (v Vector2) Distance(other Vector2) float64 {
	return math.Sqrt(math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2))
}

// 計算與目標點之間的角度
//
//	atan2(θ) = (y2-y1) / (x2-x1)
//	angleFromTargetPoint = (atan2(θ) * 180 / π + 360) % 360
func (v Vector2) GetAngleFromTargetPoint(other Vector2) float64 {
	deltaX := other.X - v.X
	deltaY := other.Y - v.Y
	return math.Mod(math.Atan2(deltaY, deltaX)*180/math.Pi+360, 360)
}

// 將角度轉向量 (θ -> v)
//
//	R (rad) = θ(°) * π / 180
//	v = Vector2{ cos(R), sin(R) }
func AngleToVector(angleDeg float64) Vector2 {
	angleRad := angleDeg * math.Pi / 180 // 角度轉弧度
	return Vector2{
		X: math.Cos(angleRad), // 計算X向量
		Y: math.Sin(angleRad), // 計算Y向量
	}
}

// 點到線的最短(垂直)距離
//
//	0: Point ; 1,2: Line 的兩端點
//	distance = |(x2-x1)(y1-y0) - (x1-x0)(y2-y1)| / sqrt((x2-x1)^2 + (y2-y1)^2)
func GetDistanceFromPointToLine(point, linePoint, lineDir Vector2) float64 {
	// Point
	x0, y0 := point.X, point.Y
	// Line
	x1, y1 := linePoint.X, linePoint.Y
	x2, y2 := linePoint.X+lineDir.X, linePoint.Y+lineDir.Y

	numerator := math.Abs((x2-x1)*(y1-y0) - (x1-x0)*(y2-y1))
	denominator := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))

	distance := numerator / denominator
	return distance
}

// 求兩點間的向量
func Direction(from, to Vector2) Vector2 {
	return Vector2{X: to.X - from.X, Y: to.Y - from.Y}
}

// 計算向量線性插植
func Lerp(start, end Vector2, t float64) Vector2 {
	return Vector2{
		X: start.X + (end.X-start.X)*t,
		Y: start.Y + (end.Y-start.Y)*t,
	}
}

// 取二維向量字串轉成二維向量
//
//	"3,2" -> Vector2{3,2}
func NewVector2(splitedStr string) (Vector2, error) {
	vSlice, err := SplitFloat(splitedStr, ",")
	if err != nil {
		return Vector2{}, fmt.Errorf("在NewVector2時SplitFloat錯誤: %v", err)
	}
	if len(vSlice) != 2 {
		return Vector2{}, fmt.Errorf("在NewVector2時SplitFloat, 結果長度不為2")
	}
	return Vector2{X: vSlice[0], Y: vSlice[1]}, nil
}

// 取三維向量字串的 X, Y 數值轉成二維向量
//
//	"3,1,3" -> Vector2{3,3}
func NewVector2XZ(splitedStr string) (Vector2, error) {
	vSlice, err := SplitFloat(splitedStr, ",")
	if err != nil {
		return Vector2{}, fmt.Errorf("在NewVector2XZ時SplitFloat錯誤: %v", err)
	}
	if len(vSlice) != 3 {
		return Vector2{}, fmt.Errorf("在NewVector2XZ時SplitFloat, 結果長度不為3")
	}
	return Vector2{X: vSlice[0], Y: vSlice[2]}, nil
}
