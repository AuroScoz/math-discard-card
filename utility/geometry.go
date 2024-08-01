package utility

type CenterRect struct {
	Center        Vector2
	Width, Height float64
}

// 將矩形切割成 N x N 的網格，並回傳各網格的中心點座標
func (rect CenterRect) GetGridPointsInRect(divide int) []Vector2 {
	var points []Vector2

	// 計算每個格子的寬高
	cellWidth := rect.Width / float64(divide)
	cellHeight := rect.Height / float64(divide)

	// 計算起始點(左下角)
	startX := rect.Center.X - rect.Width/2 + cellWidth/2
	startY := rect.Center.Y - rect.Height/2 + cellHeight/2

	// 遍歷生成格子的中心點
	for i := 0; i < divide; i++ {
		for j := 0; j < divide; j++ {
			x := startX + float64(j)*cellWidth
			y := startY + float64(i)*cellHeight
			points = append(points, Vector2{X: x, Y: y})
		}
	}

	return points
}
