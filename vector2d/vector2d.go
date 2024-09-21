package vector2d

import (
	"math"
)

type Vector2D struct {
	X, Y float64
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
func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}

// Subtract returns the vector subtraction of v2 from v.
func (v Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{v.X - v2.X, v.Y - v2.Y}
}

// Multiply returns the vector v scaled by the given scalar.
func (v Vector2D) Multiply(scalar float64) Vector2D {
	return Vector2D{v.X * scalar, v.Y * scalar}
}

// Divide returns the vector v divided by the given scalar.
func (v Vector2D) Divide(scalar float64) Vector2D {
	return Vector2D{v.X / scalar, v.Y / scalar}
}

// Dot returns the dot product of v and v2.
func (v Vector2D) Dot(v2 Vector2D) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y)
}

// Cross returns the cross product of v and v2.
func (v Vector2D) Cross(v2 Vector2D) float64 {
	return (v.X * v2.Y) - (v.Y * v2.X)
}

// Magnitude returns the magnitude (length) of the vector v.
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AngleTo returns the angle in radians between v and v2.
func (v Vector2D) AngleTo(v2 Vector2D) Radian {
	dotProduct := v.Dot(v2)
	v1Magnitude := v.Magnitude()
	v2Magnitude := v2.Magnitude()
	cosTheta := dotProduct / (v1Magnitude * v2Magnitude)
	cosTheta = math.Max(-1, math.Min(1, cosTheta)) // clamp to prevent floating point precision issues
	return Radian(math.Acos(cosTheta))
}

// ProjectOnto returns the projection of v onto v2.
func (v Vector2D) ProjectOnto(v2 Vector2D) Vector2D {
	dotProduct := v.Dot(v2)
	projectionScalar := dotProduct / math.Pow(v2.Magnitude(), 2)
	return v2.Multiply(projectionScalar)
}

// DistanceTo returns the distance between v and v2.
func (v Vector2D) DistanceTo(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2))
}

// Lerp performs linear interpolation between v and v2 by the factor t.
func (v Vector2D) Lerp(v2 Vector2D, t float64) Vector2D {
	newX := (1-t)*v.X + t*v2.X
	newY := (1-t)*v.Y + t*v2.Y
	return Vector2D{newX, newY}
}

// Rotate rotates the vector v by the given angle in radians.
func (v Vector2D) Rotate(angle Radian) Vector2D {
	newX := v.X*math.Cos(float64(angle)) - v.Y*math.Sin(float64(angle))
	newY := v.X*math.Sin(float64(angle)) + v.Y*math.Cos(float64(angle))
	return Vector2D{newX, newY}
}

// ReflectAcross returns the reflection of v across the given normal vector.
func (v Vector2D) ReflectAcross(normal Vector2D) Vector2D {
	normal = normal.Normal() // ensuring the normal is normalized
	dotProduct := v.Dot(normal)
	return Vector2D{v.X - 2*dotProduct*normal.X, v.Y - 2*dotProduct*normal.Y}
}

// Normal returns the normalized (unit) vector of v.
func (v Vector2D) Normal() Vector2D {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector2D{0, 0}
	}
	return Vector2D{v.X / magnitude, v.Y / magnitude}
}
