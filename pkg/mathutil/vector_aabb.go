package mathutil

type AABB2F [2]Vector2
type AABB3F [2]Vector3

func NewAABB2F(x1, y1, x2, y2 float64) AABB2F {
	x1, x2 = Min(x1, x2), Max(x1, x2)
	y1, y2 = Min(y1, y2), Max(y1, y2)
	return AABB2F{Vector2{x1, y1}, Vector2{x2, y2}}
}
func NewAABB2VF(p1, p2 Vector2) AABB2F { return NewAABB2F(p1[0], p1[1], p2[0], p2[1]) }

func (r AABB2F) P1() Vector2     { return r[0] }
func (r AABB2F) P2() Vector2     { return r[1] }
func (r AABB2F) Center() Vector2 { return r.P1().OffsetV(r.P2()).Scale(0.5) }

func (r AABB2F) Width() float64  { return r.P2().X() - r.P1().X() }
func (r AABB2F) Height() float64 { return r.P2().Y() - r.P1().Y() }

func (r AABB2F) Contains(point Vector2) bool {
	return point.X() >= r.P1().X() && point.X() <= r.P2().X() &&
		point.Y() >= r.P1().Y() && point.Y() <= r.P2().Y()
}

// AABB3F

func NewAABB3F(x1, y1, z1, x2, y2, z2 float64) AABB3F {
	x1, x2 = Min(x1, x2), Max(x1, x2)
	y1, y2 = Min(y1, y2), Max(y1, y2)
	z1, z2 = Min(z1, z2), Max(z1, z2)
	return AABB3F{Vector3{x1, y1, z1}, Vector3{x2, y2, z2}}
}

func NewAABB3VF(p1, p2 Vector3) AABB3F { return NewAABB3F(p1[0], p1[1], p1[2], p2[0], p2[1], p2[2]) }

func NewAABB3F_Of(positions ...Vector3) AABB3F {
	var min, max Vector3
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
	return NewAABB3VF(min, max)
}

func (r AABB3F) P1() Vector3     { return r[0] }
func (r AABB3F) P2() Vector3     { return r[1] }
func (r AABB3F) Center() Vector3 { return r.P1().OffsetV(r.P2()).Scale(0.5) }

func (r AABB3F) Width() float64  { return r.P2().X() - r.P1().X() }
func (r AABB3F) Height() float64 { return r.P2().Y() - r.P1().Y() }
func (r AABB3F) Depth() float64  { return r.P2().Z() - r.P1().Z() }

func (r AABB3F) Contains(point Vector3) bool {
	return point.X() >= r.P1().X() && point.X() <= r.P2().X() &&
		point.Y() >= r.P1().Y() && point.Y() <= r.P2().Y() &&
		point.Z() >= r.P1().Z() && point.Z() <= r.P2().Z()
}

func (r AABB3F) Points(dx, dy, dz float64) []Vector3 {
	width, height, depth := r.Height(), r.Width(), r.Depth()
	result := make([]Vector3, 0, int(width*height*depth))
	for z := r.P1().Z(); z <= r.P2().Z(); z += dz {
		for y := r.P1().Y(); y <= r.P2().Y(); y += dy {
			for x := r.P1().X(); x <= r.P2().X(); x += dx {
				result = append(result, Vector3{x, y, z})
			}
		}
	}
	return result
}
