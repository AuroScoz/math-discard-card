package utility

// CinemachinePathBase 的類型
type PositionUnits int

const (
	PATH_UNITS PositionUnits = iota // 路徑單位 (CinemachinePathBase 的類型)
	DISTANCE                        // 距離單位 (CinemachinePathBase 的類型)
	NORMALIZE                       // 標準化單位 (CinemachinePathBase 的類型)
)

// 對齊 Unity 的路徑計算，定義了一條世界空間中的路徑，包含一系列路徑點，每個點都有位置和滾動設置。在路徑點之間進行貝塞爾插值，以獲得平滑連續的路徑。
type CinemachinePathBase struct {
	m_Resolution             int // Override DistanceCacheSampleStepsPerSegment
	m_DistanceToPos          []float64
	m_PosToDistance          []float64
	m_CachedSampleSteps      int
	m_PathLength             float64
	m_cachedPosStepSize      float64
	m_cachedDistanceStepSize float64
}

func (c *CinemachinePathBase) invalidateDistanceCache() {
	c.m_DistanceToPos = nil
	c.m_PosToDistance = nil
	c.m_CachedSampleSteps = 0
	c.m_PathLength = 0
}
