package utility

import (
	"math"
)

// 二維貝茲曲線
type BezierCurve2 struct {
	ControlPoints []Vector2
}

// Interpolate 方法計算貝茲曲線上參數值 t 的插值點。
//
//	P(t) = Σ(C(n,i) * t^i * (1-t)^(n-i) * P[i])
func (b *BezierCurve2) Interpolate(t float64) Vector2 {
	// 貝茲曲線的次數 n 為控制點數量減去 1。
	// 該方法遍歷貝茲曲線的所有控制點，根據 Bernstein 多項式計算各個控制點的權重並進行加權。
	n := len(b.ControlPoints) - 1
	var p Vector2
	for i, cp := range b.ControlPoints {
		blend := bernstein(n, i, t)
		p.X += cp.X * blend
		p.Y += cp.Y * blend
	}
	return p
}

// 三維貝茲曲線
type BezierCurve3 struct {
	ControlPoints []Vector3
}

// Interpolate 方法計算貝茲曲線上參數值 t 的插值點。
//
//	P(t) = Σ(C(n,i) * t^i * (1-t)^(n-i) * P[i])
func (b *BezierCurve3) Interpolate(t float64) Vector3 {
	// 貝茲曲線的次數 n 為控制點數量減去 1。
	// 該方法遍歷貝茲曲線的所有控制點，根據 Bernstein 多項式計算各個控制點的權重並進行加權。
	n := len(b.ControlPoints) - 1
	var p Vector3
	for i, cp := range b.ControlPoints {
		blend := bernstein(n, i, t)
		p.X += cp.X * blend
		p.Y += cp.Y * blend
		p.Z += cp.Z * blend
	}
	return p
}

// Bernstein 多項式 B(n, i, t) 代表在參數值 t 下，貝茲曲線中第 i 個控制點的混合因子。
// 在貝茲曲線中，n 為曲線的次數（控制點數量減去 1）。
//
//	B(n, i, t) = C(n, i) * t^i * (1-t)^(n-i)
func bernstein(n, i int, t float64) float64 {
	return float64(binomialCoefficient(n, i)) * math.Pow(t, float64(i)) * math.Pow(1-t, float64(n-i))
}

// BinomialCoefficient 計算二項式係數 (n choose k)。
//
//	C(n, k) = n! / (k! * (n-k)!)
func binomialCoefficient(n, k int) int {
	if k > n-k {
		k = n - k
	}
	var c = 1
	for i := 0; i < k; i++ {
		c *= (n - i)
		c /= (i + 1)
	}
	return c
}
