package utility

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

const (
	defaultResolution            = 20
	unityVectorExtensionsEpsilon = 1e-5
)

// Smooth Path 上的點
type WayPoints []WayPoint
type WayPoint struct {
	Position Vector3
	Roll     float64
}

// 將路徑點轉換為 Vector4
func (w *WayPoint) Vector4() Vector4 {
	return Vector4{
		X: w.Position.X,
		Y: w.Position.Y,
		Z: w.Position.Z,
		W: w.Roll,
	}
}

// 將 []Vector4 轉成 WayPoints
func ToWayPoints(v4s []Vector4) WayPoints {
	var ws WayPoints
	for _, v := range v4s {
		ws = append(ws, WayPoint{
			Position: Vector3{
				X: v.X,
				Y: v.Y,
				Z: v.Z,
			},
			Roll: v.W,
		})
	}
	return ws
}

// 將路徑點轉換為 Vector4
func (ws WayPoints) Vector4Array() []Vector4 {
	vs := make([]Vector4, len(ws))

	for i, w := range ws {
		vs[i] = w.Vector4()
	}

	return vs
}

// 字串轉成 WayPoints，必須為浮點數且數量為 3 的倍數
//
//	"1.0,2.0,3.0,4.0,5.0,6.0" -> WayPoints{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
//	"1.0,2.0" -> nil (錯誤)
//	"a,2.0,3.0" -> nil (錯誤)
func ParseWayPoints(input string) (WayPoints, error) {
	var wayPoints WayPoints

	components := strings.Split(input, ",")
	if len(components)%3 != 0 {
		return nil, fmt.Errorf("invalid number of components")
	}

	for i := 0; i < len(components); i += 3 {
		x, err := strconv.ParseFloat(components[i], 64)
		if err != nil {
			return nil, err
		}

		y, err := strconv.ParseFloat(components[i+1], 64)
		if err != nil {
			return nil, err
		}

		z, err := strconv.ParseFloat(components[i+2], 64)
		if err != nil {
			return nil, err
		}

		wayPoints = append(wayPoints, WayPoint{Position: Vector3{X: x, Y: y, Z: z}})
	}

	return wayPoints, nil
}

// 對齊 Unity 的路徑計算，定義了一條世界空間中的路徑，包含一系列路徑點，每個點都有位置和滾動設置。在路徑點之間進行貝塞爾插值，以獲得平滑連續的路徑。
type CinemachineSmoothPath struct {
	CinemachinePathBase
	Looped         bool
	WayPoints      WayPoints
	controlPoints1 WayPoints
	controlPoints2 WayPoints
}

// 建立一條平滑路徑
func NewCinemachineSmoothPath(looped bool, wayPoints WayPoints) *CinemachineSmoothPath {
	return &CinemachineSmoothPath{
		CinemachinePathBase: CinemachinePathBase{
			m_Resolution: defaultResolution, // Resolution 為固定值
		},
		Looped:    looped,
		WayPoints: wayPoints,
	}
}

// 印出平滑路徑的內部資訊 (Debug 用)
func (c *CinemachineSmoothPath) Info() {
	fmt.Println("Resolution:", c.m_Resolution)
	fmt.Println("DistanceToPos:", c.m_DistanceToPos)
	fmt.Println("PosToDistance:", c.m_PosToDistance)
	fmt.Println("CachedSampleSteps:", c.m_CachedSampleSteps)
	fmt.Println("PathLength:", c.m_PathLength)
	fmt.Println("CachedPosStepSize:", c.m_cachedPosStepSize)
	fmt.Println("CachedDistanceStepSize:", c.m_cachedDistanceStepSize)
}

// 返回沿著路徑的某一點的空間位置
func (c *CinemachineSmoothPath) EvaluateLocalPosition(pos float64) Vector3 {
	var result Vector3
	if len(c.WayPoints) > 0 {
		c.updateControlPoints()
		pos, indexA, indexB := c.getBoundingIndices(pos)
		if indexA == indexB {
			result = c.WayPoints[indexA].Position
		} else {
			result = bezier3(
				pos-float64(indexA),
				c.WayPoints[indexA].Position, c.controlPoints1[indexA].Position,
				c.controlPoints2[indexA].Position, c.WayPoints[indexB].Position,
			)
		}
	}
	return result
}

// 返回路徑上某一點的曲線切線
func (c *CinemachineSmoothPath) EvaluateLocalTangent(pos float64) Vector3 {
	var result Vector3
	if len(c.WayPoints) > 1 {
		c.updateControlPoints()
		pos, indexA, indexB := c.getBoundingIndices(pos)
		if !c.Looped && indexA == len(c.WayPoints)-1 {
			indexA--
		}
		result = bezierTangent3(
			pos-float64(indexA),
			c.WayPoints[indexA].Position, c.controlPoints1[indexA].Position,
			c.controlPoints2[indexA].Position, c.WayPoints[indexB].Position,
		)
	}
	return result
}

// 返回路徑上某一點的方向
func (c *CinemachineSmoothPath) EvaluateLocalOrientation(pos float64) Quaternion {
	var result Quaternion
	if len(c.WayPoints) > 0 {
		var roll float64
		pos, indexA, indexB := c.getBoundingIndices(pos)
		if indexA == indexB {
			roll = c.WayPoints[indexA].Roll
		} else {
			c.updateControlPoints()
			roll = bezier1(
				pos-float64(indexA),
				c.WayPoints[indexA].Roll, c.controlPoints1[indexA].Roll,
				c.controlPoints2[indexA].Roll, c.WayPoints[indexB].Roll,
			)
		}

		fwd := c.EvaluateLocalTangent(pos)
		if !fwd.AlmostZero() {
			result = QuaternionLookRotation(fwd).RollAroundForward(roll)
		}
	}
	return result
}

// Native PathUnits 轉換成 Pos
func (c *CinemachineSmoothPath) FromNativePathUnits(pos float64, units PositionUnits) float64 {
	if units == PATH_UNITS {
		return pos
	}
	length := c.PathLength()
	if c.m_Resolution < 1 || length < unityVectorExtensionsEpsilon {
		return 0
	}
	pos = c.StandardizePos(pos)
	d := pos / c.m_cachedPosStepSize
	i := int(math.Floor(d))
	if i >= len(c.m_PosToDistance)-1 {
		pos = c.PathLength()
	} else {
		t := d - float64(i)
		pos = lerp(c.m_PosToDistance[i], c.m_PosToDistance[i+1], t)
	}
	if units == NORMALIZE {
		pos /= length
	}
	return pos
}

// Pos 轉換成 Native PathUnits
func (c *CinemachineSmoothPath) ToNativePathUnits(pos float64, units PositionUnits) float64 {
	if units == PATH_UNITS {
		return pos
	}
	if c.m_Resolution < 1 || c.PathLength() < unityVectorExtensionsEpsilon {
		return c.MinPos()
	}
	if units == NORMALIZE {
		pos *= c.PathLength()
	}
	pos = c.StandardizePathDistance(pos)
	d := pos / c.m_cachedDistanceStepSize
	i := int(math.Floor(d))
	if i >= len(c.m_DistanceToPos)-1 {
		return c.MaxPos()
	}
	t := d - float64(i)
	return c.MinPos() + lerp(c.m_DistanceToPos[i], c.m_DistanceToPos[i+1], t)
}

// 返回路徑位置的最大值
func (c *CinemachineSmoothPath) MaxPos() float64 {
	count := len(c.WayPoints) - 1
	if count < 1 {
		return 0
	}
	if c.Looped {
		return float64(count + 1)
	}
	return float64(count)
}

// 返回路徑位置的最小值
func (c *CinemachineSmoothPath) MinPos() float64 {
	return 0
}

// 返回給定的 PositionUnits 類型的最小值
func (c *CinemachineSmoothPath) MinUnit(units PositionUnits) float64 {
	if units == NORMALIZE {
		return 0
	}
	if units == DISTANCE {
		return 0
	}
	return c.MinPos()
}

// 返回給定的 PositionUnits 類型的最大值
func (c *CinemachineSmoothPath) MaxUnit(units PositionUnits) float64 {
	if units == NORMALIZE {
		return 1
	}
	if units == DISTANCE {
		return c.PathLength()
	}
	return c.MaxPos()
}

// 返回路徑長
func (c *CinemachineSmoothPath) PathLength() float64 {
	if c.m_Resolution < 1 {
		return 0
	}
	if !c.distanceCacheIsValid() {
		c.resamplePath(c.m_Resolution)
	}
	return c.m_PathLength
}

// 返回標準化的 pos，介於 MinUnit ~ MaxUnit 之間
func (c *CinemachineSmoothPath) StandardizeUnit(pos float64, units PositionUnits) float64 {
	if units == PATH_UNITS {
		return c.StandardizePos(pos)
	}
	if units == DISTANCE {
		return c.StandardizePathDistance(pos)
	}
	len := c.m_PathLength
	if len < unityVectorExtensionsEpsilon {
		return 0
	}
	return c.StandardizePathDistance(pos*len) / len
}

// 返回標準化的 distance，介於 0 ~ PathLength() 之間
func (c *CinemachineSmoothPath) StandardizePathDistance(distance float64) float64 {
	length := c.m_PathLength
	if length < math.SmallestNonzeroFloat32 {
		return 0
	}
	if c.Looped {
		distance = math.Mod(distance, length)
		if distance < 0 {
			distance += length
		}
	}
	return clamp(distance, 0, length)
}

// 返回標準化的 position，介於 MinPos ~ MaxPos 之間
func (c *CinemachineSmoothPath) StandardizePos(pos float64) float64 {
	if c.Looped && c.MaxPos() > 0 {
		pos = math.Mod(pos, c.MaxPos())
		if pos < 0 {
			pos += c.MaxPos()
		}
		return pos
	}
	return clamp(pos, 0, c.MaxPos())
}

// 更新用於貝塞爾插值的控制點
func (c *CinemachineSmoothPath) updateControlPoints() {
	numPoints := len(c.WayPoints)
	if numPoints > 1 && (c.Looped != c.isLoopedCache() ||
		c.controlPoints1 == nil || len(c.controlPoints1) != numPoints ||
		c.controlPoints2 == nil || len(c.controlPoints2) != numPoints) {
		p1, p2 := c.computeSmoothControlPoints(c.WayPoints, c.Looped)
		c.controlPoints1 = make(WayPoints, numPoints)
		c.controlPoints2 = make(WayPoints, numPoints)
		for i := 0; i < numPoints; i++ {
			c.controlPoints1[i] = p1[i]
			c.controlPoints2[i] = p2[i]
		}
	}
}

// 返回標準化位置
func (c *CinemachineSmoothPath) getBoundingIndices(pos float64) (float64, int, int) {
	pos = c.standardizePos(pos)
	numWayPoints := len(c.WayPoints)
	if numWayPoints < 2 {
		return 0, 0, 0
	}

	indexA := int(math.Floor(float64(pos)))
	if indexA >= numWayPoints {
		pos -= c.MaxPos()
		indexA = 0
	}

	indexB := indexA + 1
	if indexB == numWayPoints {
		if c.Looped {
			indexB = 0
		} else {
			indexB--
			indexA--
		}
	}
	return pos, indexA, indexB
}

func (c *CinemachineSmoothPath) distanceCacheIsValid() bool {
	return c.MaxPos() == c.MinPos() ||
		(c.m_DistanceToPos != nil && c.m_PosToDistance != nil &&
			c.m_CachedSampleSteps == c.m_Resolution &&
			c.m_CachedSampleSteps > 0)
}

// 如果路徑的末端連接以形成連續循環，則返回true
func (c *CinemachineSmoothPath) isLoopedCache() bool {
	return c.Looped
}

// 將位置標準化
func (c *CinemachineSmoothPath) standardizePos(pos float64) float64 {
	if c.Looped {
		return pos
	}
	return pos
}

// 計算貝塞爾插值的平滑控制點
func (c *CinemachineSmoothPath) computeSmoothControlPoints(wayPoints WayPoints, looped bool) (WayPoints, WayPoints) {
	numPoints := len(wayPoints)
	p1 := make(WayPoints, numPoints)
	p2 := make(WayPoints, numPoints)
	K := make([]Vector4, numPoints)
	for i := 0; i < numPoints; i++ {
		K[i] = wayPoints[i].Vector4()
	}

	if looped {
		c1, c2 := computeSmoothControlPointsLooped(K, p1.Vector4Array(), p2.Vector4Array())
		p1 = ToWayPoints(c1)
		p2 = ToWayPoints(c2)
	} else {
		c1, c2 := computeSmoothControlPointsUnlooped(K, p1.Vector4Array(), p2.Vector4Array())
		p1 = ToWayPoints(c1)
		p2 = ToWayPoints(c2)
	}
	return p1, p2
}

// ResamplePath resamples the path with given steps per segment
func (c *CinemachineSmoothPath) resamplePath(stepsPerSegment int) {
	c.invalidateDistanceCache()

	minPos := c.MinPos()
	maxPos := c.MaxPos()
	stepSize := 1.0 / math.Max(1, float64(stepsPerSegment))

	// Sample the positions
	numKeys := int(math.Round((maxPos-minPos)/stepSize)) + 1
	c.m_PosToDistance = make([]float64, numKeys)
	c.m_CachedSampleSteps = stepsPerSegment
	c.m_cachedPosStepSize = stepSize

	p0 := c.EvaluateLocalPosition(0)
	c.m_PosToDistance[0] = 0
	pos := minPos
	for i := 1; i < numKeys; i++ {
		pos += stepSize
		p := c.EvaluateLocalPosition(pos)
		d := p0.Distance(p)
		c.m_PathLength += d
		p0 = p
		c.m_PosToDistance[i] = c.m_PathLength
	}

	// Resample the distances
	c.m_DistanceToPos = make([]float64, numKeys)
	c.m_DistanceToPos[0] = 0
	if numKeys > 1 {
		stepSize = c.m_PathLength / float64(numKeys-1)
		c.m_cachedDistanceStepSize = stepSize
		distance := float64(0)
		posIndex := 1
		for i := 1; i < numKeys; i++ {
			distance += stepSize
			d := c.m_PosToDistance[posIndex]
			for d < distance && posIndex < numKeys-1 {
				posIndex++
				d = c.m_PosToDistance[posIndex]
			}
			d0 := c.m_PosToDistance[posIndex-1]
			delta := d - d0
			if math.Abs(delta) > unityVectorExtensionsEpsilon {
				t := (distance - d0) / delta
				c.m_DistanceToPos[i] = c.m_cachedPosStepSize * (t + float64(posIndex-1))
			} else {
				c.m_DistanceToPos[i] = c.m_cachedPosStepSize * float64(posIndex-1)
			}
		}
	}
}

// 計算循環樣條曲線的平滑切線值，結果切線保證曲線的二階平滑度。
func computeSmoothControlPointsUnlooped(knot, ctrl1, ctrl2 []Vector4) (c1, c2 []Vector4) {
	numPoints := len(knot)
	if numPoints <= 2 {
		if numPoints == 2 {
			ctrl1[0] = knot[0].Lerp(knot[1], 0.33333)
			ctrl2[0] = knot[0].Lerp(knot[1], 0.66666)
		} else if numPoints == 1 {
			ctrl1[0] = knot[0]
			ctrl2[0] = knot[0]
		}
		return ctrl1, ctrl2
	}

	a := make([]float64, numPoints)
	b := make([]float64, numPoints)
	c := make([]float64, numPoints)
	r := make([]float64, numPoints)
	for axis := 0; axis < 4; axis++ {
		n := numPoints - 1

		// Linear into the first segment
		a[0] = 0
		b[0] = 2
		c[0] = 1
		r[0] = knot[0].Get(axis) + 2*knot[1].Get(axis)

		// Internal segments
		for i := 1; i < n-1; i++ {
			a[i] = 1
			b[i] = 4
			c[i] = 1
			r[i] = 4*knot[i].Get(axis) + 2*knot[i+1].Get(axis)
		}

		// Linear out of the last segment
		a[n-1] = 2
		b[n-1] = 7
		c[n-1] = 0
		r[n-1] = 8*knot[n-1].Get(axis) + knot[n].Get(axis)

		// Solve with Thomas algorithm
		for i := 1; i < n; i++ {
			m := a[i] / b[i-1]
			b[i] -= m * c[i-1]
			r[i] -= m * r[i-1]
		}

		// Compute ctrl1
		ctrl1[n-1].Set(axis, r[n-1]/b[n-1])
		for i := n - 2; i >= 0; i-- {
			ctrl1[i].Set(axis, (r[i]-c[i]*ctrl1[i+1].Get(axis))/b[i])
		}

		// Compute ctrl2 from ctrl1
		for i := 0; i < n; i++ {
			ctrl2[i].Set(axis, 2*knot[i+1].Get(axis)-ctrl1[i+1].Get(axis))
		}
		ctrl2[n-1].Set(axis, 0.5*(knot[n].Get(axis)+ctrl1[n-1].Get(axis)))
	}

	return ctrl1, ctrl2
}

// 計算循環樣條曲線的平滑切線值，結果切線保證曲線的二階平滑度。
func computeSmoothControlPointsLooped(knot, ctrl1, ctrl2 []Vector4) (c1, c2 []Vector4) {
	numPoints := len(knot)
	if numPoints < 2 {
		if numPoints == 1 {
			ctrl1[0] = knot[0]
			ctrl2[0] = knot[0]
		}
		return ctrl1, ctrl2
	}

	margin := int(math.Min(4, float64(numPoints-1)))
	knotLooped := make([]Vector4, numPoints+2*margin)
	ctrl1Looped := make([]Vector4, numPoints+2*margin)
	ctrl2Looped := make([]Vector4, numPoints+2*margin)
	for i := 0; i < margin; i++ {
		knotLooped[i] = knot[numPoints-(margin-i)]
		knotLooped[numPoints+margin+i] = knot[i]
	}
	for i := 0; i < numPoints; i++ {
		knotLooped[i+margin] = knot[i]
	}
	ctrl1, ctrl2 = computeSmoothControlPointsUnlooped(knotLooped, ctrl1Looped, ctrl2Looped)
	for i := 0; i < numPoints; i++ {
		ctrl1[i] = ctrl1Looped[i+margin]
		ctrl2[i] = ctrl2Looped[i+margin]
	}

	return ctrl1, ctrl2
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

// 將值限制在 0 ~ 1 之間
func clamp01(value float64) float64 {
	return math.Max(0.0, math.Min(1.0, value))
}

// 將值限制在 min ~ max 之間
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// 計算貝塞爾插值
func bezier3(t float64, p0, p1, p2, p3 Vector3) Vector3 {
	t = clamp01(t)
	d := 1 - t

	return p0.Mul(d * d * d).
		Add(p1.Mul(3 * d * d * t)).
		Add(p2.Mul(3 * d * t * t)).
		Add(p3.Mul(t * t * t))
}

// 計算貝塞爾曲線的切線
func bezierTangent3(t float64, p0, p1, p2, p3 Vector3) Vector3 {
	t = clamp01(t)
	d := 1 - t

	return p0.Mul(-3 * d * d).
		Add(p1.Mul(3*d*d - 6*d*t)).
		Add(p2.Mul(-3*t*t + 6*d*t)).
		Add(p3.Mul(3 * t * t))
}

// 計算單個值的貝塞爾插值
func bezier1(t, p0, p1, p2, p3 float64) float64 {
	t = clamp01(t)
	d := 1 - t

	return d*d*d*p0 +
		3*d*d*t*p1 +
		3*d*t*t*p2 +
		t*t*t*p3
}

// 計算單個四點一維貝茲曲線的切線 (Unused)
func _() { _ = bezierTangent1 }
func bezierTangent1(t float64, p0, p1, p2, p3 float64) float64 {
	t = clamp01(t)
	return (-3*p0+9*p1-9*p2+3*p3)*t*t + (6*p0-12*p1+6*p2)*t - 3*p0 + 3*p1
}

// 四元數
type Quaternion struct {
	X, Y, Z, W float64
}

// 使用指定的前向方向創建一個旋轉
func QuaternionLookRotation(forward Vector3) Quaternion {
	forward.Normalize()
	up := Vector3{0, 1, 0}
	right := forward.Cross(up)
	up = right.Cross(forward)
	return QuaternionLookRotationInternal(forward, up)
}

// 使用指定的前向和向上方向創建一個旋轉
func QuaternionLookRotationInternal(forward, up Vector3) Quaternion {
	identity := Quaternion{0, 0, 0, 1}
	if reflect.DeepEqual(forward, Vector3{0, 0, -1}) && reflect.DeepEqual(up, Vector3{0, 1, 0}) {
		return identity
	}

	rotation := identity
	if !reflect.DeepEqual(forward, Vector3{0, 0, 1}) {
		forward.Normalize()
		axis := forward.Cross(Vector3{0, 0, 1})
		angle := float64(math.Acos(float64(forward.Dot(Vector3{0, 0, 1}))))
		rotation = QuaternionAngleAxis(angle, axis)
	}

	return rotation
}

// 創建一個繞著軸旋轉一定角度的旋轉。
func QuaternionAngleAxis(angle float64, axis Vector3) Quaternion {
	axis.Normalize()
	sinHalfAngle := float64(math.Sin(float64(angle) * 0.5))
	cosHalfAngle := float64(math.Cos(float64(angle) * 0.5))

	return Quaternion{
		axis.X * sinHalfAngle,
		axis.Y * sinHalfAngle,
		axis.Z * sinHalfAngle,
		cosHalfAngle,
	}
}

// 繞前向量旋轉四元數一定角度
func (q Quaternion) RollAroundForward(angle float64) Quaternion {
	halfAngle := angle * 0.5
	sinHalfAngle := float64(math.Sin(float64(halfAngle)))
	cosHalfAngle := float64(math.Cos(float64(halfAngle)))

	forward := Vector3{0, 0, 1}
	axis := Vector3{q.X, q.Y, q.Z}
	right := forward.Cross(axis)
	right.Normalize()

	return Quaternion{
		right.X * sinHalfAngle,
		right.Y * sinHalfAngle,
		right.Z * sinHalfAngle,
		cosHalfAngle,
	}
}
