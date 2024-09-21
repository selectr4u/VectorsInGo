package vector2d

import (
	"math"
	"testing"
)

func TestDegreesToRadian(t *testing.T) {
	degrees := Degrees(180)
	expected := Radian(math.Pi)
	result := DegreesToRadian(degrees)
	if result != expected {
		t.Errorf("DegreesToRadian(%v) = %v; want %v", degrees, result, expected)
	}
}

func TestRadianToDegrees(t *testing.T) {
	radian := Radian(math.Pi)
	expected := Degrees(180)
	result := RadianToDegrees(radian)
	if result != expected {
		t.Errorf("RadianToDegrees(%v) = %v; want %v", radian, result, expected)
	}
}

func TestAdd(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{3, 4}
	expected := Vector2D{4, 6}
	result := v1.Add(v2)
	if result != expected {
		t.Errorf("Add(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestSubtract(t *testing.T) {
	v1 := Vector2D{5, 6}
	v2 := Vector2D{3, 4}
	expected := Vector2D{2, 2}
	result := v1.Subtract(v2)
	if result != expected {
		t.Errorf("Subtract(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestMultiply(t *testing.T) {
	v := Vector2D{1, 2}
	scalar := 3.0
	expected := Vector2D{3, 6}
	result := v.Multiply(scalar)
	if result != expected {
		t.Errorf("Multiply(%v, %v) = %v; want %v", v, scalar, result, expected)
	}
}

func TestDivide(t *testing.T) {
	v := Vector2D{6, 8}
	scalar := 2.0
	expected := Vector2D{3, 4}
	result := v.Divide(scalar)
	if result != expected {
		t.Errorf("Divide(%v, %v) = %v; want %v", v, scalar, result, expected)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{3, 4}
	expected := 11.0
	result := v1.Dot(v2)
	if result != expected {
		t.Errorf("Dot(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestCross(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{3, 4}
	expected := -2.0
	result := v1.Cross(v2)
	if result != expected {
		t.Errorf("Cross(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector2D{3, 4}
	expected := 5.0
	result := v.Magnitude()
	if result != expected {
		t.Errorf("Magnitude(%v) = %v; want %v", v, result, expected)
	}
}

func TestAngleTo(t *testing.T) {
	v1 := Vector2D{1, 0}
	v2 := Vector2D{0, 1}
	expected := Radian(math.Pi / 2)
	result := v1.AngleTo(v2)
	if result != expected {
		t.Errorf("AngleTo(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestProjectOnto(t *testing.T) {
	v1 := Vector2D{3, 4}
	v2 := Vector2D{1, 0}
	expected := Vector2D{3, 0}
	result := v1.ProjectOnto(v2)
	if result != expected {
		t.Errorf("ProjectOnto(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestDistanceTo(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{4, 6}
	expected := 5.0
	result := v1.DistanceTo(v2)
	if result != expected {
		t.Errorf("DistanceTo(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestLerp(t *testing.T) {
	v1 := Vector2D{0, 0}
	v2 := Vector2D{10, 10}
	tValue := 0.5
	expected := Vector2D{5, 5}
	result := v1.Lerp(v2, tValue)
	if result != expected {
		t.Errorf("Lerp(%v, %v, %v) = %v; want %v", v1, v2, tValue, result, expected)
	}
}

func TestRotate(t *testing.T) {
	v := Vector2D{1, 0}
	angle := Radian(math.Pi / 2)
	expected := Vector2D{0, 1}
	result := v.Rotate(angle)
	if math.Abs(result.X-expected.X) > 1e-9 || math.Abs(result.Y-expected.Y) > 1e-9 {
		t.Errorf("Rotate(%v, %v) = %v; want %v", v, angle, result, expected)
	}
}

func TestReflectAcross(t *testing.T) {
	v := Vector2D{1, -1}
	normal := Vector2D{0, 1}
	expected := Vector2D{1, 1}
	result := v.ReflectAcross(normal)
	if result != expected {
		t.Errorf("ReflectAcross(%v, %v) = %v; want %v", v, normal, result, expected)
	}
}

func TestNormal(t *testing.T) {
	v := Vector2D{3, 4}
	expected := Vector2D{0.6, 0.8}
	result := v.Normal()
	if result != expected {
		t.Errorf("Normal(%v) = %v; want %v", v, result, expected)
	}
}
