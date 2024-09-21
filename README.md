# VectorsInGo

Hey there! Welcome to VectorsInGo, a super cool project that helps you work with 2D and 3D vectors in Go. Whether you're adding, subtracting, or rotating vectors, we've got you covered. This project includes two packages: `vector2d` and `vector3d`. Let's dive in and see what each package offers!

## vector2d Package

The `vector2d` package provides a `Vector2D` type and various methods for 2D vector operations. Here's a quick overview of what you can do with it.

### Vector2D Type

```go
type Vector2D struct {
    X, Y float64
}
```

### Methods

- `DegreesToRadian(angle Degrees) Radian`: Converts an angle from degrees to radians.
- `RadianToDegrees(radian Radian) Degrees`: Converts an angle from radians to degrees.
- `Add(v2 Vector2D) Vector2D`: Returns the vector addition of v and v2.
- `Subtract(v2 Vector2D) Vector2D`: Returns the vector subtraction of v2 from v.
- `Multiply(scalar float64) Vector2D`: Returns the vector v scaled by the given scalar.
- `Divide(scalar float64) Vector2D`: Returns the vector v divided by the given scalar.
- `Dot(v2 Vector2D) float64`: Returns the dot product of v and v2.
- `Cross(v2 Vector2D) float64`: Returns the cross product of v and v2.
- `Magnitude() float64`: Returns the magnitude (length) of the vector v.
- `AngleTo(v2 Vector2D) Radian`: Returns the angle in radians between v and v2.
- `ProjectOnto(v2 Vector2D) Vector2D`: Returns the projection of v onto v2.
- `DistanceTo(v2 Vector2D) float64`: Returns the distance between v and v2.
- `Lerp(v2 Vector2D, t float64) Vector2D`: Performs linear interpolation between v and v2 by the factor t.
- `Rotate(angle Radian) Vector2D`: Rotates the vector v by the given angle in radians.
- `ReflectAcross(normal Vector2D) Vector2D`: Returns the reflection of v across the given normal vector.
- `Normal() Vector2D`: Returns the normalized (unit) vector of v.

## vector3d Package

The `vector3d` package provides a `Vector3D` type and various methods for 3D vector operations. Here's what you can do with it.

### Vector3D Type

```go
type Vector3D struct {
    X, Y, Z float64
}
```

### Methods

- `DegreesToRadian(angle Degrees) Radian`: Converts an angle from degrees to radians.
- `RadianToDegrees(radian Radian) Degrees`: Converts an angle from radians to degrees.
- `Add(v2 Vector3D) Vector3D`: Returns the vector addition of v and v2.
- `Subtract(v2 Vector3D) Vector3D`: Returns the vector subtraction of v2 from v.
- `Multiply(scalar float64) Vector3D`: Returns the vector v scaled by the given scalar.
- `Divide(scalar float64) Vector3D`: Returns the vector v divided by the given scalar.
- `Dot(v2 Vector3D) float64`: Returns the dot product of v and v2.
- `Cross(v2 Vector3D) Vector3D`: Returns the cross product of v and v2.
- `Magnitude() float64`: Returns the magnitude (length) of the vector v.
- `AngleTo(v2 Vector3D) Radian`: Returns the angle in radians between v and v2.
- `ProjectOnto(v2 Vector3D) Vector3D`: Returns the projection of v onto v2.
- `DistanceTo(v2 Vector3D) float64`: Returns the distance between v and v2.
- `Lerp(v2 Vector3D, t float64) Vector3D`: Performs linear interpolation between v and v2 by the factor t.
- `ReflectAcross(normal Vector3D) Vector3D`: Returns the reflection of v across the given normal vector.
- `Normal() Vector3D`: Returns the normalized (unit) vector of v.

## How to Use

To use these packages, simply import them into your Go project and start working with vectors. Here's a quick example to get you started:

```go
package main

import (
    "fmt"
    "github.com/selectr4u/VectorsInGo/vector2d"
    "github.com/selectr4u/VectorsInGo/vector3d"
)

func main() {
    v1 := vector2d.Vector2D{X: 1, Y: 2}
    v2 := vector2d.Vector2D{X: 3, Y: 4}
    result := v1.Add(v2)
    fmt.Println("2D Vector Addition:", result)

    v3 := vector3d.Vector3D{X: 1, Y: 2, Z: 3}
    v4 := vector3d.Vector3D{X: 4, Y: 5, Z: 6}
    result3D := v3.Add(v4)
    fmt.Println("3D Vector Addition:", result3D)
}
```

Happy coding! 