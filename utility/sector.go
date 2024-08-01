package utility

import ()

// 扇形
type Sector struct {
	Center     Vector2 // 中心點
	Radius     float64 // 半徑
	Angle      float64 // 扇形中軸線
	AngleRange float64 // 扇形張開角度
}

// 判斷目標點是否在扇形面積內
func (sector Sector) IsPointInSector(targetPoint Vector2) bool {
	// 計算與目標點之間的角度
	toTargetAngle := sector.Center.GetAngleFromTargetPoint(targetPoint)

	// 判斷是否在半徑內
	if sector.Center.Distance(targetPoint) <= sector.Radius {
		// 判斷 與目標點之間角度 跟 扇形中軸線角度 差，在扇形張開角度內與否
		if AngleDiff(toTargetAngle, sector.Angle) <= sector.AngleRange/2 {
			return true
		}
	}

	return false
}
