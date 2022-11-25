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

// UpdateFibre calculates the new angle of rotation and changes the position of the fibre
// Input: Nearest cell, matrix stiffness float64 value
func (fibre *Fibre) UpdateFibre(cell *Cell, S float64) {

	// phi (angle of rotation) = theta (angle of cell from pivot) - arcsin[(1 - 0.1*integrins*(1-stiffness)*perpendicular distance D) / hypotenuse]
	// otherEnd := fibre.FindPivot(cell)
	fibre.FindPivot(cell)
	d := ComputeDistance(fibre.pivot, cell.position) // find the hypotenuse of the cell to the pivot point of the fibre
	D := fibre.FindPerpendicularDistance(cell)       // find the perpendicular distance of the cell to the fibre
	theta := math.Asin(d / D)                        // theta = arcsin(d / D)

	phi := ComputePhi(fibre, cell, theta, d, D, S) // compute angle of rotation
	m, b := FindLine(fibre.pivot, cell.position)
	yCoord := FindYOnLine(cell.position.x, m, b)
	if fibre.pivot.x < cell.position.x {
		if yCoord < cell.position.y {
			phi *= -1
		}
	} else {
		if yCoord > cell.position.y {
			phi *= -1
		}
	}
	fibre.UpdateDirection(phi) // updates direction vector of fibre using phi
	fibre.UpdatePosition()
}

// FindPivot:  Determines which end of the fibre will be considered the pivot with regards to a cell
// The pivot is the end of the fibre that is furthest away from the center of the cell.
// Also returns the non-pivot end as an ordered pair.
// Input: fibre (*Fibre) pointer to the fibre being analyzed.
// cell (*Cell) pointer to the cell object being analyzed.
// Output: OrderedPair, the end of the fibre that is not a pivot.
// func (fibre *Fibre) FindPivot(cell *Cell) OrderedPair {
func (fibre *Fibre) FindPivot(cell *Cell) {
	var endpoint1, endpoint2 OrderedPair
	magnitude := math.Sqrt(fibre.direction.x*fibre.direction.x + fibre.direction.y*fibre.direction.y)
	//calculate the ends of the fibres
	endpoint1.x = fibre.position.x + fibre.direction.x*0.5*fibre.length/magnitude
	endpoint1.y = fibre.position.y + fibre.direction.y*0.5*fibre.length/magnitude
	endpoint2.x = fibre.position.x - fibre.direction.x*0.5*fibre.length/magnitude
	endpoint2.y = fibre.position.y - fibre.direction.y*0.5*fibre.length/magnitude
	//calculate the distance between the cell and each endpoint to identify the pivot
	distance1 := ComputeDistance(cell.position, endpoint1)
	distance2 := ComputeDistance(cell.position, endpoint2)
	if distance1 < distance2 {
		fibre.pivot.x = endpoint2.x
		fibre.pivot.y = endpoint2.y
		// return endpoint1
	} else {
		fibre.pivot.x = endpoint1.x
		fibre.pivot.y = endpoint1.y
		fibre.direction.x *= -1
		fibre.direction.y *= -1
		// return endpoint2
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
	A := fibre.pivot.y - fibre.position.y
	B := fibre.position.x - fibre.pivot.x
	C := fibre.pivot.x*fibre.position.y - fibre.position.x*fibre.pivot.y
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
