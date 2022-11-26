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
		// fmt.Println("Generation i: ", timeFrames[gen].cells[0].position)
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

// Normalize: Normalizes an ordered pair so that its magnitude = 1.
// Input: p1 (*OrderedPair) a pointer to the OrderedPair to be normalized.
func (p1 *OrderedPair) Normalize() {
	magnitude := p1.Magnitude()
	p1.x /= magnitude
	p1.y /= magnitude
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

// ProjectVector: Returns the projection of one vector onto another.
// Input:
// v1 (OrderedPair): The vector be projected
// v2 (OrderedPair): The vector being projected onto.
// Output: (OrderedPair) The resulting projection vector.
func ProjectVector(v1, v2 OrderedPair) OrderedPair {
	scalingFactor := DotProduct2D(v1, v2) / (v2.Magnitude() * v2.Magnitude())
	newVector := MultiplyVectorByConstant2D(v2, scalingFactor)
	return newVector
}

// DotProduct2D: Returns the dot product between two vectors in R2.
// Input: v1, v2 (OrderedPair) the vectors to be dotted.
// Output: (float64) The dot product of v1 and v2.
func DotProduct2D(v1, v2 OrderedPair) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

// MultiplyVectorByConstant2D: Multiplies a vector by some constant value.
// Input:
// v1 (OrderedPair): The vector to be multiplied
// constant (float64): The consant to multiply the vector with
// Output: (OrderedPair): The new vector as an OrderedPair object.
func MultiplyVectorByConstant2D(v1 OrderedPair, constant float64) OrderedPair {
	newVector := v1
	newVector.x *= constant
	newVector.y *= constant
	return newVector
}

// CalculateAngleBetweenVectors2D: Finds the angle between two vectors.
// Input:
// v1, v2 (OrderedPair): The two vectors being compared.
// Output:
// (float64) The angle between v1 and v2 in radians.
func CalculateAngleBetweenVectors2D(v1, v2 OrderedPair) float64 {
	cosTheta := DotProduct2D(v1, v2) / (v1.Magnitude() * v2.Magnitude())
	theta := math.Acos(cosTheta)
	return theta
}
