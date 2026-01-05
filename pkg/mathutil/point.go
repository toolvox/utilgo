package mathutil

type Point2 [2]int
type Point3 [3]int
type Point4 [4]int

func (v Point2) X() int         { return v[0] }
func (v Point2) Y() int         { return v[1] }
func (v Point2) XY() (int, int) { return v[0], v[1] }
func (v Point2) Point2() Point2 { return v }
func (v Point2) Point3() Point3 { return Point3{v[0], v[1]} }
func (v Point2) Point4() Point4 { return Point4{v[0], v[1]} }
func (v Point2) OffsetPt(o Point2) Point2 {
	v[0] += o[0]
	v[1] += o[1]
	return v
}
func (v Point2) Offset(x, y int) Point2 {
	v[0] += x
	v[1] += y
	return v
}

func (v Point3) X() int         { return v[0] }
func (v Point3) Y() int         { return v[1] }
func (v Point3) Z() int         { return v[2] }
func (v Point3) Point2() Point2 { return Point2{v[0], v[1]} }
func (v Point3) Point3() Point3 { return v }
func (v Point3) Point4() Point4 { return Point4{v[0], v[1], v[2]} }
func (v Point3) OffsetPt(o Point3) Point3 {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
	return v
}
func (v Point3) Offset(x, y, z int) Point3 {
	v[0] += x
	v[1] += y
	v[2] += z
	return v
}

func (v Point4) X() int         { return v[0] }
func (v Point4) Y() int         { return v[1] }
func (v Point4) Z() int         { return v[2] }
func (v Point4) W() int         { return v[3] }
func (v Point4) Point2() Point2 { return Point2{v[0], v[1]} }
func (v Point4) Point3() Point3 { return Point3{v[0], v[1], v[2]} }
func (v Point4) Point4() Point4 { return v }
func (v Point4) OffsetPt(o Point4) Point4 {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
	v[3] += o[3]
	return v
}
func (v Point4) Offset(x, y, z, w int) Point4 {
	v[0] += x
	v[1] += y
	v[2] += z
	v[3] += w
	return v
}

var (
	N = Point2{0, 1}
	S = Point2{0, -1}
	E = Point2{1, 0}
	W = Point2{-1, 0}
)
