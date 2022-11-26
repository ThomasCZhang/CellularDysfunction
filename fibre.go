package main

/*
type Fibre struct {
	length, width         float64
	position, pivot       OrderedPair
	direction, projection OrderedPair
}
*/

import (
	"math"
)

// UpdateFibre calculates the new angle of rotation and changes the position of the fibre
// Input: Nearest cell, matrix stiffness float64 value
func (fibre *Fibre) UpdateFibre(cell *Cell, stiffness float64) {
	fibre.FindPivot(cell)
	phi := fibre.ComputePhi(cell, stiffness) // compute angle of rotation
	fibre.UpdateDirection(phi)               // updates direction vector of fibre using phi
	fibre.UpdatePosition()
}

// FindPivot:  Determines which end of the fibre will be considered the pivot with regards to a cell.
// The pivot is the end of the fibre that is furthest away from the center of the cell.
// Additionally sets the direction vector such that pivot + direction vector = center of fibre.
// Input:
// fibre (*Fibre) pointer to the fibre being analyzed.
// cell (*Cell) pointer to the cell object being analyzed.
func (fibre *Fibre) FindPivot(cell *Cell) {
	endpoint1, endpoint2 := fibre.GetEndpoints()
	//calculate the distance between the cell and each endpoint to identify the pivot
	distance1 := ComputeDistance(cell.position, endpoint1)
	distance2 := ComputeDistance(cell.position, endpoint2)
	if distance1 > distance2 {
		// If endpoint1 is the pivot
		fibre.pivot.x = endpoint1.x
		fibre.pivot.y = endpoint1.y
		fibre.direction.x *= -1
		fibre.direction.y *= -1
	} else {
		// If endpoint 2 is the pivot
		fibre.pivot.x = endpoint2.x
		fibre.pivot.y = endpoint2.y
	}
}

// ComputePhi: Calculates the angle of rotation between the fibre pivot and the center of the cell
// Input:
// cell (*Cell): Pointer to the cell object that is acting on the fibre.
// S (float64): Matrix stiffness.
// Output:
// (float64) The angle that the fibre needs to be rotated about its pivot.
func (f *Fibre) ComputePhi(cell *Cell, S float64) float64 {
	// phi (angle of rotation) = theta (angle of cell from pivot) - arcsin[(1 - 0.1*integrins*(1-stiffness)*perpendicular distance D) / hypotenuse]

	D := f.FindPerpendicularDistance(cell)       // perpendicular distance from cell to fibre
	d := ComputeDistance(f.pivot, cell.position) // distance from pivot point to cell
	I := cell.integrin                           // Percentage of integrins expressed by the cell
	theta := f.FindTheta(D, d)                   // Angle between fibre and line from pivot to cell.

	alignFactor := (1 - 0.1*I*(1-S))
	phi := (theta - math.Asin(alignFactor*D/d))
	rotationSign := f.DetermineRotationDirection(cell)
	return float64((rotationSign)) * phi
}

// Determines whether the fibre needs to be rotated in the clockwise (negative) or counter-clockwise (positive) direction
// to align the fibre closer to the center of the cell acting on the fibre. Returns 1 if rotation should be counter-clockwise,
// Otherwise returns 0.
// Input:
// f (*Fibre) Pointer to the fibre being acted on
// cell (*Cell) Pointer to the cell acting on the fibre
// Output:
// (int) -1 if rotation should be clockwise, 1 if rotation should be clockwise.
func (f *Fibre) DetermineRotationDirection(cell *Cell) int {
	nonPivot := f.GetNonPivot() // non-pivot end of the fibre

	m, b := FindLine(f.pivot, cell.position) // y = mx + b of line between pivot and cell center
	y := EvaluateLineAtX(nonPivot.x, m, b)

	rotation := 1
	if f.pivot.x < cell.position.x { // If pivot is to the left of cell.
		if nonPivot.y > y { // Non-Pivot end is above the line from pivot to cell
			rotation *= -1
		}
	} else { // If Pivot is to the right of the cell.
		if nonPivot.y < y { // Non-Pivot end is below the line from pivot to cell
			rotation *= -1
		}
	}
	return rotation
}

// UpdateDirection: Updates the direction of a fibre
// Input: f (*Fibre) pointer to a Fibre object.
func (f *Fibre) UpdateDirection(phi float64) {
	x := f.direction.x
	y := f.direction.y
	f.direction.x = x*math.Cos(phi) - y*math.Sin(phi)
	f.direction.y = x*math.Sin(phi) + y*math.Cos(phi)
}

// UpdatePosition: Updates the position of a fibre.
// Input: fibre (*Fibre) location of the center of the fibre.
// Output: None.
func (f *Fibre) UpdatePosition() {
	magnitude := math.Sqrt(f.direction.x*f.direction.x + f.direction.y*f.direction.y)
	f.position.x = f.pivot.x + f.direction.x*0.5*f.length/magnitude
	f.position.y = f.pivot.y + f.direction.y*0.5*f.length/magnitude
}

// FindPerpendicularDistance: Finds shortest distance from the center of a cell to a fibre.
// Input:
// fibre (*Fibre) Pointer to a fibre object.
// cell (*Cell) Pointer to a cell object.
// Output:
// (float64) The distance from the center of the cell to the fibre.
func (fibre *Fibre) FindPerpendicularDistance(cell *Cell) float64 {
	// need to find coordinates of pivot point
	A, B, C := FindHomogenousLine(fibre.position, fibre.pivot)
	numerator := math.Abs(A*cell.position.x + B*cell.position.y + C)
	denominator := math.Sqrt(A*A + B*B)
	return numerator / denominator
}

// FindNearestCell: Finds the cell nearest to the center of a Fibre object.
// Input: fibre (*Fibre) a pointer to the Fibre object.
// cells ([]*Cell) A slice of pointers to cell objects. This slice contains all the cells in the ECM.
func (fibre *Fibre) FindNearestCell(cells []*Cell) *Cell {
	// Placeholder function.
	nearestCell := cells[0]
	currentDistance := fibre.FindPerpendicularDistance(cells[0])
	for _, cell := range cells {
		newDistance := fibre.FindPerpendicularDistance(cell)
		if newDistance < currentDistance {
			nearestCell = cell
		}
	}
	return nearestCell
}

// GetEndpoints: Gets the coordinates of the two ends of a fibre object and returns them as OrderedPairs
// Input: fibre (*Fibre) A pointer to the Fibre object
// Output: endpoint1, endpoint2 (OrderedPair): The two ends of the Fibre object.
func (fibre *Fibre) GetEndpoints() (OrderedPair, OrderedPair) {
	var endpoint1, endpoint2 OrderedPair
	magnitude := fibre.direction.Magnitude()
	//calculate the ends of the fibres
	endpoint1.x = fibre.position.x + fibre.direction.x*0.5*fibre.length/magnitude
	endpoint1.y = fibre.position.y + fibre.direction.y*0.5*fibre.length/magnitude
	endpoint2.x = fibre.position.x - fibre.direction.x*0.5*fibre.length/magnitude
	endpoint2.y = fibre.position.y - fibre.direction.y*0.5*fibre.length/magnitude

	return endpoint1, endpoint2
}

// GetLine: Determines the line that describes a fibre in the form y = mx + b.
// Returns m and b as float64.
// Output: m, b (float64) the m and b values from the equation "y = mx + b"
func (fibre *Fibre) GetLine() (float64, float64) {
	endpoint1, endpoint2 := fibre.GetEndpoints()
	m, b := FindLine(endpoint1, endpoint2)
	return m, b
}

// GetNonPivot: Returns the non-pivot end of a fibre as an OrderedPair
// Output: (OrderedPair) The non-pivot end of the fibre.
func (fibre *Fibre) GetNonPivot() OrderedPair {
	endpoint1, endpoint2 := fibre.GetEndpoints()
	if fibre.pivot != endpoint1 {
		return endpoint1
	}
	return endpoint2
}

// FindTheta: Find the angle between the line describing the fibre and the line between the pivot and the cell
func (f *Fibre) FindTheta(opposite, hypotenuse float64) float64 {
	theta := math.Asin(opposite / hypotenuse) // theta = arcsin(opposite / hypotenuse)
	if theta != theta {
		panic("Error in FindTheta: arcsin returned NaN.")
	}
	return theta
}

// CopyFibre: Returns a pointer to a copy of a Fibre object.
func (f *Fibre) CopyFibre() *Fibre {
	var newFibre Fibre
	newFibre.direction = f.direction
	newFibre.length = f.length
	newFibre.position = f.position
	newFibre.width = f.width
	newFibre.pivot = f.pivot
	return &newFibre
}

// This is just a dead function... Wth is it even doing?
func (fibre *Fibre) FindHypotenuse(currCell *Cell) float64 {
	// Takes the position of the fibre
	posFibre := fibre.position
	var fibreCoord OrderedPair

	fibreCoord.x = posFibre.x + (fibre.length)/2
	fibreCoord.y = posFibre.y + (fibre.length)/2

	// Get the hypotenuse
	hypotenuse := ComputeDistance(fibreCoord, currCell.position)
	return hypotenuse
}
