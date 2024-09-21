package vector3d

import "math"

type Vector3D struct {
	X, Y, Z float64
}

type Degrees float64
type Radian float64

// DegreesToRadian converts an angle from degrees to radians.
func DegreesToRadian(angle Degrees) Radian {
	return Radian(angle * (math.Pi / 180))
}

// RadianToDegrees converts an angle from radians to degrees.
func RadianToDegrees(radian Radian) Degrees {
	return Degrees(radian * (180 / math.Pi))
}

// Add returns the vector addition of v and v2.
func (v Vector3D) Add(v2 Vector3D) Vector3D {
	return Vector3D{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

// Subtract returns the vector subtraction of v2 from v.
func (v Vector3D) Subtract(v2 Vector3D) Vector3D {
	return Vector3D{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

// Multiply returns the vector v scaled by the given scalar.
func (v Vector3D) Multiply(scalar float64) Vector3D {
	return Vector3D{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

// Divide returns the vector v divided by the given scalar.
func (v Vector3D) Divide(scalar float64) Vector3D {
	return Vector3D{v.X / scalar, v.Y / scalar, v.Z / scalar}
}

// Dot returns the dot product of v and v2.
func (v Vector3D) Dot(v2 Vector3D) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y) + (v.Z * v2.Z)
}

// Cross returns the cross product of v and v2.
func (v Vector3D) Cross(v2 Vector3D) Vector3D {
	newX := (v.Y * v2.Z) - (v.Z * v2.Y)
	newY := (v.Z * v2.X) - (v.X * v2.Z)
	newZ := (v.X * v2.Y) - (v.Y * v2.X)
	return Vector3D{newX, newY, newZ}
}

// Magnitude returns the magnitude (length) of the vector v.
func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// AngleTo returns the angle in radians between v and v2.
func (v Vector3D) AngleTo(v2 Vector3D) Radian {
	dotProduct := v.Dot(v2)
	v1Magnitude := v.Magnitude()
	v2Magnitude := v2.Magnitude()
	cosTheta := dotProduct / (v1Magnitude * v2Magnitude)
	cosTheta = math.Max(-1, math.Min(1, cosTheta)) // clamp to prevent floating point precision issues
	return Radian(math.Acos(cosTheta))
}

// Lerp performs linear interpolation between v and v2 by the factor t.
func (v Vector3D) Lerp(v2 Vector3D, t float64) Vector3D {
	newX := (1-t)*v.X + t*v2.X
	newY := (1-t)*v.Y + t*v2.Y
	newZ := (1-t)*v.Z + t*v2.Z
	return Vector3D{newX, newY, newZ}
}

// ProjectOnto returns the projection of v onto v2.
func (v Vector3D) ProjectOnto(v2 Vector3D) Vector3D {
	dotProduct := v.Dot(v2)
	projectionScalar := dotProduct / math.Pow(v2.Magnitude(), 2)
	return v2.Multiply(projectionScalar)
}

// DistanceTo returns the distance between v and v2.
func (v Vector3D) DistanceTo(v2 Vector3D) float64 {
	return math.Sqrt(math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2) + math.Pow(v2.Z-v.Z, 2))
}

// ReflectAcross returns the reflection of v across the given normal vector.
func (v Vector3D) ReflectAcross(normal Vector3D) Vector3D {
	normal = normal.Normal() // ensure the normal is normalized
	dotProduct := v.Dot(normal)
	return Vector3D{v.X - 2*dotProduct*normal.X, v.Y - 2*dotProduct*normal.Y, v.Z - 2*dotProduct*normal.Z}
}

// Normal returns the normalized (unit) vector of v.
func (v Vector3D) Normal() Vector3D {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector3D{0, 0, 0}
	}
	return Vector3D{v.X / magnitude, v.Y / magnitude, v.Z / magnitude}
}
