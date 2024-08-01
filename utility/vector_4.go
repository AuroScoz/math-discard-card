package utility

type Vector4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v Vector4) ToArray() [4]float64 {
	return [4]float64{v.X, v.Y, v.Z, v.W}
}
func (v Vector4) ToSlice() []float64 {
	return []float64{v.X, v.Y, v.Z, v.W}
}

func (v Vector4) Add(other Vector4) Vector4 {
	return Vector4{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
		W: v.W + other.W,
	}
}

func (v Vector4) Sub(other Vector4) Vector4 {
	return Vector4{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v Vector4) Mul(scalar float64) Vector4 {
	return Vector4{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
		W: v.W * scalar,
	}
}

func (v Vector4) Lerp(other Vector4, t float64) Vector4 {
	return Vector4{
		X: v.X + (other.X-v.X)*t,
		Y: v.Y + (other.Y-v.Y)*t,
		Z: v.Z + (other.Z-v.Z)*t,
		W: v.W + (other.W-v.W)*t,
	}
}

// 取得 Axis 方位的數值
func (v Vector4) Get(index int) float64 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	case 3:
		return v.W
	default:
		return 0
	}
}

// 設定 Axis 方位的數值
func (v *Vector4) Set(index int, value float64) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	case 3:
		v.W = value
	}
}
