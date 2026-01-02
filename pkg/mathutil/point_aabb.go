package mathutil

type AABB2 [2]Point2
type AABB3 [2]Point3

func NewAABB2(x1, y1, x2, y2 int) AABB2 {
	x1, x2 = Min(x1, x2), Max(x1, x2)
	y1, y2 = Min(y1, y2), Max(y1, y2)
	return AABB2{Point2{x1, y1}, Point2{x2, y2}}
}
func NewAABB2V(p1, p2 Point2) AABB2 { return NewAABB2(p1[0], p1[1], p2[0], p2[1]) }

func (r AABB2) P1() Point2     { return r[0] }
func (r AABB2) P2() Point2     { return r[1] }
func (r AABB2) Center() Point2 { return r.P1().OffsetPt(r.P2()).ScaleF(0.5) }

func (r AABB2) Width() int  { return r.P2().X() - r.P1().X() }
func (r AABB2) Height() int { return r.P2().Y() - r.P1().Y() }

func (r AABB2) Contains(point Point2) bool {
	return point.X() >= r.P1().X() && point.X() <= r.P2().X() &&
		point.Y() >= r.P1().Y() && point.Y() <= r.P2().Y()
}

// AABB3

func NewAABB3(x1, y1, z1, x2, y2, z2 int) AABB3 {
	x1, x2 = Min(x1, x2), Max(x1, x2)
	y1, y2 = Min(y1, y2), Max(y1, y2)
	z1, z2 = Min(z1, z2), Max(z1, z2)
	return AABB3{Point3{x1, y1, z1}, Point3{x2, y2, z2}}
}

func NewAABB3P(p1, p2 Point3) AABB3 { return NewAABB3(p1[0], p1[1], p1[2], p2[0], p2[1], p2[2]) }

func NewAABB3_Of(positions ...Point3) AABB3 {
	var min, max Point3
	for _, p := range positions {
		for i := 0; i < len(p); i++ {
			if p[i] < min[i] {
				min[i] = p[i]
			}
			if p[i] > max[i] {
				max[i] = p[i]
			}
		}
	}
	return NewAABB3P(min, max)
}

func (r AABB3) P1() Point3     { return r[0] }
func (r AABB3) P2() Point3     { return r[1] }
func (r AABB3) Center() Point3 { return r.P1().OffsetPt(r.P2()).ScaleF(0.5) }

func (r AABB3) Width() int  { return r.P2().X() - r.P1().X() }
func (r AABB3) Height() int { return r.P2().Y() - r.P1().Y() }
func (r AABB3) Depth() int  { return r.P2().Z() - r.P1().Z() }

func (r AABB3) Contains(point Point3) bool {
	return point.X() >= r.P1().X() && point.X() <= r.P2().X() &&
		point.Y() >= r.P1().Y() && point.Y() <= r.P2().Y() &&
		point.Z() >= r.P1().Z() && point.Z() <= r.P2().Z()
}

func (r AABB3) Points(dx, dy, dz int) []Point3 {
	width, height, depth := r.Height(), r.Width(), r.Depth()
	result := make([]Point3, 0, int(width*height*depth))
	for z := r.P1().Z(); z <= r.P2().Z(); z += dz {
		for y := r.P1().Y(); y <= r.P2().Y(); y += dy {
			for x := r.P1().X(); x <= r.P2().X(); x += dx {
				result = append(result, Point3{x, y, z})
			}
		}
	}
	return result
}
