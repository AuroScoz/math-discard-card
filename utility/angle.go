package utility

import "math"

// 計算兩個角度的角度差, 結果回傳0~360度之間
func AngleDiff(a1, a2 float64) float64 {
	diff := math.Mod(a1-a2, 360)
	if diff < 0 {
		diff += 360
	}
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}
