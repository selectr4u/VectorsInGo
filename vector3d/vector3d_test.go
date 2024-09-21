package vector3d

import (
	"math"
	"testing"
)

// nearEnough checks if two floating-point numbers are close enough to be considered equal.
func nearEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDegreesToRadian(t *testing.T) {
	degrees := Degrees(180)
	expected := Radian(math.Pi)
	result := DegreesToRadian(degrees)
	if !nearEnough(float64(result), float64(expected), 1e-9) {
		t.Errorf("DegreesToRadian(%v) = %v; want %v", degrees, result, expected)
	}
}

func TestRadianToDegrees(t *testing.T) {
	radian := Radian(math.Pi)
	expected := Degrees(180)
	result := RadianToDegrees(radian)
	if !nearEnough(float64(result), float64(expected), 1e-9) {
		t.Errorf("RadianToDegrees(%v) = %v; want %v", radian, result, expected)
	}
}

func TestAdd(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	expected := Vector3D{5, 7, 9}
	result := v1.Add(v2)
	if result != expected {
		t.Errorf("Add(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestSubtract(t *testing.T) {
	v1 := Vector3D{4, 5, 6}
	v2 := Vector3D{1, 2, 3}
	expected := Vector3D{3, 3, 3}
	result := v1.Subtract(v2)
	if result != expected {
		t.Errorf("Subtract(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestMultiply(t *testing.T) {
	v := Vector3D{1, 2, 3}
	scalar := 2.0
	expected := Vector3D{2, 4, 6}
	result := v.Multiply(scalar)
	if result != expected {
		t.Errorf("Multiply(%v, %v) = %v; want %v", v, scalar, result, expected)
	}
}

func TestDivide(t *testing.T) {
	v := Vector3D{2, 4, 6}
	scalar := 2.0
	expected := Vector3D{1, 2, 3}
	result := v.Divide(scalar)
	if result != expected {
		t.Errorf("Divide(%v, %v) = %v; want %v", v, scalar, result, expected)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	expected := 32.0
	result := v1.Dot(v2)
	if !nearEnough(result, expected, 1e-9) {
		t.Errorf("Dot(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestCross(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	expected := Vector3D{-3, 6, -3}
	result := v1.Cross(v2)
	if !nearEnough(result.X, expected.X, 1e-9) || !nearEnough(result.Y, expected.Y, 1e-9) || !nearEnough(result.Z, expected.Z, 1e-9) {
		t.Errorf("Cross(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector3D{1, 2, 2}
	expected := 3.0
	result := v.Magnitude()
	if !nearEnough(result, expected, 1e-9) {
		t.Errorf("Magnitude(%v) = %v; want %v", v, result, expected)
	}
}

func TestAngleTo(t *testing.T) {
	v1 := Vector3D{1, 0, 0}
	v2 := Vector3D{0, 1, 0}
	expected := Radian(math.Pi / 2)
	result := v1.AngleTo(v2)
	if !nearEnough(float64(result), float64(expected), 1e-9) {
		t.Errorf("AngleTo(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestLerp(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	tValue := 0.5
	expected := Vector3D{2.5, 3.5, 4.5}
	result := v1.Lerp(v2, tValue)
	if !nearEnough(result.X, expected.X, 1e-9) || !nearEnough(result.Y, expected.Y, 1e-9) || !nearEnough(result.Z, expected.Z, 1e-9) {
		t.Errorf("Lerp(%v, %v, %v) = %v; want %v", v1, v2, tValue, result, expected)
	}
}

func TestProjectOnto(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	expected := Vector3D{1.662337662337662, 2.077922077922078, 2.4935064935064934}
	result := v1.ProjectOnto(v2)
	if !nearEnough(result.X, expected.X, 1e-9) || !nearEnough(result.Y, expected.Y, 1e-9) || !nearEnough(result.Z, expected.Z, 1e-9) {
		t.Errorf("ProjectOnto(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestDistanceTo(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	expected := math.Sqrt(27)
	result := v1.DistanceTo(v2)
	if !nearEnough(result, expected, 1e-9) {
		t.Errorf("DistanceTo(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestReflectAcross(t *testing.T) {
	v := Vector3D{1, -1, 0}
	normal := Vector3D{0, 1, 0}
	expected := Vector3D{1, 1, 0}
	result := v.ReflectAcross(normal)
	if !nearEnough(result.X, expected.X, 1e-9) || !nearEnough(result.Y, expected.Y, 1e-9) || !nearEnough(result.Z, expected.Z, 1e-9) {
		t.Errorf("ReflectAcross(%v, %v) = %v; want %v", v, normal, result, expected)
	}
}

func TestNormal(t *testing.T) {
	v := Vector3D{1, 2, 2}
	expected := Vector3D{1 / 3.0, 2 / 3.0, 2 / 3.0}
	result := v.Normal()
	if !nearEnough(result.X, expected.X, 1e-9) || !nearEnough(result.Y, expected.Y, 1e-9) || !nearEnough(result.Z, expected.Z, 1e-9) {
		t.Errorf("Normal(%v) = %v; want %v", v, result, expected)
	}
}
