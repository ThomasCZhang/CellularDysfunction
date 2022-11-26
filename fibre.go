package main

/*
type Fibre struct {
	length, width              float64
	position, direction, pivot OrderedPair
}
*/

import (
	"math"
)

// UpdateFibre calculates the new angle of rotation and changes the position of the fibre
// Input: Nearest cell, matrix stiffness float64 value
func (fibre *Fibre) UpdateFibre(cell *Cell, S float64) {

	// phi (angle of rotation) = theta (angle of cell from pivot) - arcsin[(1 - 0.1*integrins*(1-stiffness)*perpendicular distance D) / hypotenuse]
	// otherEnd := fibre.FindPivot(cell)
	fibre.FindPivot(cell)
	d := ComputeDistance(fibre.pivot, cell.position) // find the hypotenuse of the cell to the pivot point of the fibre
	D := fibre.FindPerpendicularDistance(cell)       // find the perpendicular distance of the cell to the fibre
	theta := math.Asin(D / d)                        // theta = arcsin(d / D)

	phi := ComputePhi(fibre, cell, theta, d, D, S) // compute angle of rotation
	m, b := fibre.GetLine()
	y := FindYOnLine(cell.position.x, m, b)
	if fibre.pivot.x < cell.position.x {
		// If pivot is to the left of cell.
		if cell.position.y < y {
			// And the cell is below the line describing the fibre.
			phi *= -1
		}
	} else {
		// If Pivot is to the right of the cell.
		if cell.position.y > y {
			// And the cell is above the line describing the fibre.
			phi *= -1
		}
	}

	fibre.UpdateDirection(phi) // updates direction vector of fibre using phi
	fibre.UpdatePosition()
	if math.IsNaN(fibre.direction.x) ||
		math.IsNaN(fibre.direction.y) ||
		math.IsNaN(fibre.position.x) ||
		math.IsNaN(fibre.position.y) {
		panic("Error in UpdateFibre. Position or Direction are NaN (Not a Number).")
	}
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

// ComputePhi: Calculates the new angle between the fibre pivot and the center of the cell
// Input:
// theta (float64): The initial angle between the fibre pivot and the center of the cell
func ComputePhi(fibre *Fibre, cell *Cell, theta, d, D, S float64) float64 {
	alignFactor := (1 - 0.1*cell.integrin*(1-S))
	phi := (theta - math.Asin(alignFactor*D/d))
	return phi
}

// UpdateDirection: Updates the direction of a fibre
// Input: fibre (*Fibre) pointer to a Fibre object.
func (fibre *Fibre) UpdateDirection(phi float64) {
	x := fibre.direction.x
	y := fibre.direction.y
	fibre.direction.x = x*math.Cos(phi) - y*math.Sin(phi)
	fibre.direction.y = x*math.Sin(phi) + y*math.Cos(phi)
}

//	UpdatePosition: Updates the position of a fibre.
//
// Input: fibre (*Fibre) location of the center of the fibre.
// Output: None.
func (fibre *Fibre) UpdatePosition() {
	magnitude := math.Sqrt(fibre.direction.x*fibre.direction.x + fibre.direction.y*fibre.direction.y)
	fibre.position.x = fibre.pivot.x + fibre.direction.x*0.5*fibre.length/magnitude
	fibre.position.y = fibre.pivot.y + fibre.direction.y*0.5*fibre.length/magnitude
}

// FindPerpendicularDistance: Finds shortest distance from the center of a cell to a fibre.
// Input:
// fibre (*Fibre) Pointer to a fibre object.
// cell (*Cell) Pointer to a cell object.
// Output: (float64) The distance from the center of the cell to the fibre.
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
	if fibre.pivot == endpoint1 {
		return endpoint1
	}
	return endpoint2
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
