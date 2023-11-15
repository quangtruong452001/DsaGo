package Vector3D

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const EPSILON float64 = 1e-8

type Vector3D struct {
	x, y, z float64
}

func NewVector3D(x float64, y float64, z float64) *Vector3D {
	return &Vector3D{x: x, y: y, z: z}
}

// Copy value to another vector
func (p *Vector3D) Copy() Vector3D {
	return Vector3D{p.x, p.y, p.z}
}

// Equals compare to vector and return boolean
func (v Vector3D) Equals(rhs Vector3D) bool {
	return math.Abs(v.x-rhs.X()) < EPSILON &&
		math.Abs(v.y-rhs.Y()) < EPSILON &&
		math.Abs(v.z-rhs.Z()) < EPSILON
}

// SetX sets the x component of the Vector3D.
func (v *Vector3D) SetX(x float64) {
	v.x = x
}

// X returns the x component of the Vector3D.
func (v Vector3D) X() float64 {
	return v.x
}

// SetY sets the y component of the Vector3D.
func (v *Vector3D) SetY(y float64) {
	v.y = y
}

// Y returns the y component of the Vector3D.
func (v Vector3D) Y() float64 {
	return v.y
}

// SetZ sets the z component of the Vector3D.
func (v *Vector3D) SetZ(z float64) {
	v.z = z
}

// Z returns the z component of the Vector3D.
func (v Vector3D) Z() float64 {
	return v.z
}

// Add performs vector addition and returns a new Vector3D.
func (v *Vector3D) Add(other Vector3D) Vector3D {
	return Vector3D{v.x + other.x, v.y + other.y, v.z + other.z}
}

// Sub performs vector subtraction and returns a new Vector3D.
func (v *Vector3D) Sub(other Vector3D) Vector3D {
	return Vector3D{v.x - other.x, v.y - other.y, v.z - other.z}
}

// Negate negates the vector and returns a new Vector3D.
func (v *Vector3D) Negate() Vector3D {
	return Vector3D{-v.x, -v.y, -v.z}
}

// DotProduct calculates the dot product of two vectors and returns a float64.
func (v *Vector3D) DotProduct(other Vector3D) float64 {
	return v.x*other.x + v.y*other.y + v.z*other.z
}

// ScalarMultiply multiplies the vector by a scalar and returns a new Vector3D.
func (v *Vector3D) ScalarMultiply(scalar float64) Vector3D {
	return Vector3D{v.x * scalar, v.y * scalar, v.z * scalar}
}

// Length calculates the length (magnitude) of the vector and returns a float64.
func (v *Vector3D) Length() float64 {
	l2 := v.x*v.x + v.y*v.y + v.z*v.z
	return math.Sqrt(l2)
}

// Normalize normalizes the vector in-place and returns a reference to the modified Vector3D.
func (v *Vector3D) Normalize() *Vector3D {
	length := v.Length()
	v.x /= length
	v.y /= length
	v.z /= length
	return v
}

// Angle calculates the angle between two vectors and returns it in degrees.
func (v *Vector3D) Angle(other Vector3D) float64 {
	dotProduct := v.DotProduct(other)
	La := v.Length()
	Lb := other.Length()
	radianAngle := math.Acos(dotProduct / (La * Lb))
	return radianAngle * 180 / math.Pi
}

/*
 a: this
 b: other
 -------CROSS = DETERMINANT --------
 i    j     k
 a.x  a.y   a.z
 b.x  b.y   b.z
*/

// Cross calculates the cross product of two vectors and returns a new Vector3D.
func (v *Vector3D) Cross(other Vector3D) Vector3D {
	return Vector3D{
		v.y*other.z - other.y*v.z,
		-(v.x*other.z - other.x*v.z),
		v.x*other.y - other.x*v.y,
	}
}

// Vector2Str converts a Vector3D pointer to a string.
func (v *Vector3D) Vector2Str() string {
	return fmt.Sprintf("V(X: %.2f, Y: %.2f, Z: %.2f)", v.x, v.y, v.z)
}

// GenVectors generates an array of Vector3D objects with random coordinates.
func GenVectors(size int, minValue, maxValue float64, manualSeed bool, seedValue int64) []*Vector3D {
	vectors := make([]*Vector3D, size)
	var source rand.Source

	if manualSeed {
		source = rand.NewSource(seedValue)
	} else {
		source = rand.NewSource(time.Now().UnixNano())
	}

	rng := rand.New(source)

	for idx := 0; idx < size; idx++ {
		x := minValue + (maxValue-minValue)*rng.Float64()
		y := minValue + (maxValue-minValue)*rng.Float64()
		z := minValue + (maxValue-minValue)*rng.Float64()
		vectors[idx] = NewVector3D(x, y, z)
	}

	return vectors
}

// PrintVectors prints an array of Vector3D objects as a formatted string.
func PrintVectors(vectors []*Vector3D) {
	var builder strings.Builder
	builder.WriteString("[")
	for idx, vector := range vectors {
		builder.WriteString(vector.Vector2Str())

		if idx < len(vectors)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	fmt.Println(builder.String())
}
