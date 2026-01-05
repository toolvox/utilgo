package mathutil

import "math"

type Vector2 [2]float64
type Vector3 [3]float64
type Vector4 [4]float64

func (v Vector2) X() float64 { return v[0] }
func (v Vector2) Y() float64 { return v[1] }

func (v Vector2) XY() (float64, float64) { return v[0], v[1] }

func (v Vector2) Vector2() Vector2 { return v }
func (v Vector2) Vector3() Vector3 { return Vector3{v[0], v[1]} }
func (v Vector2) Vector4() Vector4 { return Vector4{v[0], v[1]} }
func (v Vector2) OffsetV(o Vector2) Vector2 {
	v[0] += o[0]
	v[1] += o[1]
	return v
}
func (v Vector2) Scale(f float64) Vector2 {
	v[0] *= f
	v[1] *= f
	return v
}
func (v Vector2) Offset(x, y float64) Vector2 {
	v[0] += x
	v[1] += y
	return v
}
func (v Vector2) Norm() Vector2 {
	mag := float64(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
	if mag == 0 {
		return v
	}
	return Vector2{v[0] / mag, v[1] / mag}
}

func (v Vector3) X() float64 { return v[0] }
func (v Vector3) Y() float64 { return v[1] }
func (v Vector3) Z() float64 { return v[2] }

func (v Vector3) XYZ() (float64, float64, float64) { return v[0], v[1], v[2] }

func (v Vector3) Vector2() Vector2 { return Vector2{v[0], v[1]} }
func (v Vector3) Vector3() Vector3 { return v }
func (v Vector3) Vector4() Vector4 { return Vector4{v[0], v[1], v[2]} }
func (v Vector3) OffsetV(o Vector3) Vector3 {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
	return v
}
func (v Vector3) Scale(f float64) Vector3 {
	v[0] *= f
	v[1] *= f
	v[2] *= f
	return v
}
func (v Vector3) Offset(x, y, z float64) Vector3 {
	v[0] += x
	v[1] += y
	v[2] += z
	return v
}

func (v Vector3) Norm() Vector3 {
	mag := float64(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
	if mag == 0 {
		return v
	}
	return Vector3{v[0] / mag, v[1] / mag, v[2] / mag}
}

func (v Vector4) X() float64 { return v[0] }
func (v Vector4) Y() float64 { return v[1] }
func (v Vector4) Z() float64 { return v[2] }
func (v Vector4) W() float64 { return v[3] }

func (v Vector4) XYZW() (float64, float64, float64, float64) { return v[0], v[1], v[2], v[3] }

func (v Vector4) Vector2() Vector2 { return Vector2{v[0], v[1]} }
func (v Vector4) Vector3() Vector3 { return Vector3{v[0], v[1], v[2]} }
func (v Vector4) Vector4() Vector4 { return v }
func (v Vector4) OffsetV(o Vector4) Vector4 {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
	v[3] += o[3]
	return v
}
func (v Vector4) Scale(f float64) Vector4 {
	v[0] *= f
	v[1] *= f
	v[2] *= f
	v[3] *= f
	return v
}
func (v Vector4) Offset(x, y, z, w float64) Vector4 {
	v[0] += x
	v[1] += y
	v[2] += z
	v[3] += w
	return v
}

func (v Vector4) Norm() Vector4 {
	mag := float64(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
	if mag == 0 {
		return v
	}
	return Vector4{v[0] / mag, v[1] / mag, v[2] / mag, v[3] / mag}
}
