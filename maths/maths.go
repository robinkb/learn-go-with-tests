package maths

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

// Point represents a two-dimensional Carthesian coordinate
type Point struct {
	X, Y float64
}

type Clockface struct{}

// SecondHand is the unit vector of the second hand of analogue clock
// at time 't' represented as a Point.
func (c *Clockface) SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // Scale
	p = Point{p.X, -p.Y}                                      // Flip
	p = Point{p.X + clockCentreX, -p.Y + clockCentreY}        // Translate
	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
