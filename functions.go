package main

import (
	"math"
)

// SimulateCellMotility takes a ECM object of cells and fibres and updates it over certain number of generations with a specified timestep.
// Input: a initialECM, numGens and a timestep
// Output: A slice of numGens+1 ECM objects that model cell and fibre movement.
func SimulateCellMotility(initialECM *ECM, numGens int, time float64) []*ECM {
	// range over some number of generations
	timeFrames := make([]*ECM, numGens+1)
	timeFrames[0] = initialECM
	for gen := 1; gen <= numGens; gen++ {
		timeFrames[gen] = timeFrames[gen-1].UpdateECM(time)
	}
	return timeFrames
}

// Distance: Takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func ComputeDistance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// Magnitude: Calculate the magnitude of an ordered pair.
// Formula for magnitude = (x^2+y^2)^(1/2)
func (p1 *OrderedPair) Magnitude() float64 {
	var mag float64
	mag = math.Sqrt((p1.x * p1.x) + (p1.y * p1.y))
	return mag
}

// FindLine: Find the equation of the line between two Ordered Pairs.
// Calculates the equation of the line in the form "y = mx+b"
// Input:
// p1, p2 (Ordered Pair) The two ordered pairs that the line should pass through
// Output:
// m, b (float64) the m and b values of the equation for the line "y = mx + b"
func FindLine(p1, p2 OrderedPair) (float64, float64) {
	m := (p2.y - p1.y) / (p2.x - p1.x)
	b := p1.y - m*p1.x
	return m, b
}

// FindLine: Find the equation of the line between two points.
// Calculates the equation of the line in the form "Ax + By + C = 0" and returns A B and C.
// Input:
// p1, p2 (Ordered Pair) The two ordered pairs that the line should pass through
// Output:
// A, B, C (float64) the values of A, B and C in the homogenous line equation Ax+By+C = 0
func FindHomogenousLine(p1, p2 OrderedPair) (float64, float64, float64) {
	A := p2.y - p1.y
	B := p1.x - p2.x
	C := p2.x*p1.y - p1.x*p2.y
	return A, B, C
}

// EvaluateLineAtX: Calculates the y value resulting from plugging x into "y = m*x + b"
// Input: x (float64) the x-value to use
// m , b (float64) m and b in the line equation "y = mx+b"
func EvaluateLineAtX(x, m, b float64) float64 {
	return (m*x + b)
}
