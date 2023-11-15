package Point3D

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const EPSILON float64 = 1e-8

// can export out of package

type Point3D struct {
	x, y, z float64
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	return &Point3D{x: x, y: y, z: z}
}

// Copy value to another point
func (p *Point3D) Copy() Point3D {
	return Point3D{p.x, p.y, p.z}
}

// public

// SetX sets the x coordinate of the Point3D.
func (p *Point3D) SetX(x float64) {
	p.x = x
}

// X returns the x coordinate of the Point3D.
func (p Point3D) X() float64 {
	return p.x
}

// SetY sets the y coordinate of the Point3D.
func (p *Point3D) SetY(y float64) {
	p.y = y
}

// Y returns the y coordinate of the Point3D.
func (p Point3D) Y() float64 {
	return p.y
}

// SetZ sets the z coordinate of the Point3D.
func (p *Point3D) SetZ(z float64) {
	p.z = z
}

// Z returns the z coordinate of the Point3D.
func (p Point3D) Z() float64 {
	return p.z
}

// Radius return distance from point to (0, 0, 0)
func (p *Point3D) Radius() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

// PointEQ compare to point and return boolean
func (p *Point3D) PointEQ(rhs Point3D) bool {
	return math.Abs(p.x-rhs.x) < EPSILON &&
		math.Abs(p.y-rhs.y) < EPSILON &&
		math.Abs(p.z-rhs.z) < EPSILON
}

// Add two point
func (p *Point3D) Add(rhs Point3D) Point3D {
	return Point3D{p.x + rhs.x, p.x + rhs.y, p.z + rhs.z}
}

// Sub two point
func (p *Point3D) Sub(rhs Point3D) Point3D {
	return Point3D{p.x - rhs.x, p.y - rhs.y, p.z - rhs.z}
}

// Point2Str converts a Point3D pointer to a string.
func (p *Point3D) Point2Str() string {
	return fmt.Sprintf("P(X: %.2f, Y: %.2f, Z: %.2f)", p.x, p.y, p.z)
}

func (p *Point3D) Println() {
	println(p.Point2Str())
}

// GenPoints generates an array of Point3D objects with random coordinates.
func GenPoints(size int, minValue, maxValue float64, manualSeed bool, seedValue int64) []*Point3D {
	points := make([]*Point3D, size)

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
		points[idx] = NewPoint3D(x, y, z)
	}

	return points
}

// GenPointsNormal generates an array of Point3D objects with coordinates following a normal distribution.
func GenPointsNormal(size int, mu, sigma []float64, manualSeed bool, seedValue int64) []*Point3D {
	points := make([]*Point3D, size)

	var source rand.Source

	if manualSeed {
		source = rand.NewSource(seedValue)
	} else {
		source = rand.NewSource(time.Now().UnixNano())
	}

	rng := rand.New(source)

	for idx := 0; idx < size; idx++ {
		x := mu[0] + sigma[0]*rng.NormFloat64()
		y := mu[1] + sigma[1]*rng.NormFloat64()
		z := mu[2] + sigma[2]*rng.NormFloat64()
		points[idx] = NewPoint3D(x, y, z)
	}

	return points
}

// PrintPoints prints an array of Point3D objects as a formatted string.
func PrintPoints(points []*Point3D) {
	var builder strings.Builder
	builder.WriteString("[")
	for idx, point := range points {
		//builder.WriteString(fmt.Sprintf("X=%.2f, Y=%.2f, Z=%.2f", point.x, point.y, point.z))
		builder.WriteString(point.Point2Str())

		if idx < len(points)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	fmt.Println(builder.String())
}
