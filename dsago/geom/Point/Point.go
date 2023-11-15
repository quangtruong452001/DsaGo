package Point

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const EPSILON float64 = 1e-8

// can export out of package

type Point struct {
	x, y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// Copy value to another point
func (p *Point) Copy() Point {
	return Point{p.x, p.y}
}

// public

// SetX sets the x coordinate of the Point.
func (p *Point) SetX(x float64) {
	p.x = x
}

// X returns the x coordinate of the Point.
func (p Point) X() float64 {
	return p.x
}

// SetY sets the y coordinate of the Point.
func (p *Point) SetY(y float64) {
	p.y = y
}

// Y returns the y coordinate of the Point.
func (p Point) Y() float64 {
	return p.y
}

// Radius return distance from point to (0, 0, 0)
func (p *Point) Radius() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// Add two point
func (p *Point) Add(rhs Point) Point {
	return Point{p.x + rhs.x, p.x + rhs.y}
}

// Sub two point
func (p *Point) Sub(rhs Point) Point {
	return Point{p.x - rhs.x, p.y - rhs.y}
}

// Point2Str converts a Point pointer to a string.
func (p *Point) Point2Str() string {
	return fmt.Sprintf("P(X: %.2f, Y: %.2f)", p.x, p.y)
}

// Println prints a Point object as a formatted string.
func (p *Point) Println() {
	println(p.Point2Str())
}

// GenPoints generates an array of Point objects with random coordinates.
func GenPoints(size int, minValue, maxValue float64, manualSeed bool, seedValue int64) []*Point {
	points := make([]*Point, size)

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
		points[idx] = NewPoint(x, y)
	}

	return points
}

// GenPointsNormal generates an array of Point objects with coordinates following a normal distribution.
func GenPointsNormal(size int, mu, sigma []float64, manualSeed bool, seedValue int64) []*Point {
	points := make([]*Point, size)

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
		points[idx] = NewPoint(x, y)
	}

	return points
}

// PrintPoints prints an array of Point objects as a formatted string.
func PrintPoints(points []*Point) {
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

// PointEQ compare to point and return boolean
func PointEQ(lhs, rhs *Point) bool {
	return math.Abs(lhs.x-rhs.x) < EPSILON &&
		math.Abs(lhs.y-rhs.y) < EPSILON
}

// PointEQ compare to point and return boolean
//func (p *Point) PointEQ(rhs *Point) bool {
//	return math.Abs(p.x-rhs.x) < EPSILON &&
//		math.Abs(p.y-rhs.y) < EPSILON
//}

func PointEqual(lhs, rhs interface{}) bool {
	return PointEQ(lhs.(*Point), rhs.(*Point))
}
