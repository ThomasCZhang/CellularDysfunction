package main

import (
	"fmt"
	"math"
	"math/rand"
)

// SimulateCellMotility takes a ECM object of cells and fibres and updates it over certain number of generations with a specified timestep.
// Input: a initialECM, numGens and a timestep
// Output: A slice of numGens+1 ECM objects that model cell and fibre movement.
func SimulateCellMotility(initialECM *ECM, numGens int, time float64) []*ECM {
	// range over some number of generations
	timeFrames := make([]*ECM, numGens+1)
	timeFrames[0] = initialECM
	for gen := 1; gen <= numGens; gen++ {
		timeFrames[gen] = UpdateECM(timeFrames[gen-1], time)
	}
	return timeFrames
}

// UpdateECM takes a current ECM object and updates it by the given time step.
// Input: currentECM and a time step
// Output: A new ECM object with  updated cell and fibre positions
func UpdateECM(currentECM *ECM, time float64) *ECM {
	// range over all cells on ECM

	newECM := CopyECM(currentECM) // makes a deep copy of the ECM

	for _, fibre := range newECM.fibres {
		nearestCell := fibre.FindNearestCell(newECM.cells) // returns a nearest cell
		fibre.UpdateFibre(nearestCell, newECM.stiffness)
	}

	for _, cell := range newECM.cells {
		cell.UpdateCell()
	}
	return newECM
}

// CopyECM creates a deep copy of the given ECM
func CopyECM(currentECM *ECM) *ECM {
	return
}

// UpdateCell finds the new direction of a given cell based on the fibre that causes the least change in direction
// Input: cell object and a list of updated fibres
// Output: cell with updated projection and position
func (cell *Cell) UpdateCell(fibres []*fibre, threshold float64) {

	// range over all fibres and compute projection vectors caused by all fibres on the cell
	// to make this easier, we can only pick fibres that are within a certain critical distance to the cell

	NearestFibres := FindNearestFibres(cell, threshold) // returns a slice of nearest fibres within a certain threshold distance
	var changeMagnitude float64
	for index, fibre := NearestFibres {
		newProjection := FindProjection(cell, fibre)
		delta_magnitude := FindMagnitudeChange(newProjection, cell.projection)
		if index == 0 { // set default magnitude of change to first one
			changeMagnitude = delta_magnitude
		}
		if delta_magnitude < changeMagnitude { // find the minimum magnitude of change
			changeMagnitude = delta_magnitude
		}
		cell.UpdateProjection(newProjection) // Update the projection vector
	}

	cell.UpdatePosition()
}



// FindProjection finds the new projection vector caused by a fibre on the given cell
// Input: Current cell and fibre
// Output: The new projection vector of the cell caused by the fibre based on the direction of the fibre and random noise.
func FindProjection(cell *Cell, fibre *Fibre) OrderedPair {
}


// UpdateProjection updates the projection vector of the cell using a new projection vector
func (cell *Cell) UpdateProjection(newProjection OrderedPair) {
	cell.ComputeDragForce() // computes F = speed x shape factor (c) x fluid viscosity (n) x projection vector + noise
}

func ComputeDragForce() {
	ComputeNoise()
}

func ComputeNoise() {

}

func FindNearestFibre() {

}
// UpdatePosition uses the updated projection vector to change the position of the cell.
func (cell *Cell) UpdatePosition() {

}

// UpdateFibre calculates the new angle of rotation and changes the position of the fibre
// Input: Nearest cell, matrix stiffness float64 value
func (fibre *Fibre) UpdateFibre(cell *Cell, S float64) {

	// phi (angle of rotation) = theta (angle of cell from pivot) - arcsin[(1 - 0.1*integrins*(1-stiffness)*perpendicular distance D) / hypotenuse]
	otherEnd := fibre.FindPivot(cell)
	d := ComputeDistance(fibre.pivot, cell.position) // find the hypotenuse of the cell to the pivot point of the fibre
	D := fibre.FindPerpendicularDistance(cell) // find the perpendicular distance of the cell to the fibre
	theta := math.Asin(d/D) // theta = arcsin(d / D)

	phi := ComputePhi(theta, d, D, S, fibre, cell) // compute angle of rotation
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

func (fibre *Fibre) FindPivot(cell *Cell) OrderedPair {
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
		return endpoint1
	} else {
		fibre.pivot.x = endpoint1.x
		fibre.pivot.y = endpoint1.y
		fibre.direction.x *= -1
		fibre.direction.y *= -1
		return endpoint2
	}
}

func ComputePhi(theta, d, D, S float64, fibre *Fibre, cell *Cell) float64 {
	alignFactor := (1 - 0.1*cell.integrin*(1-S))
	phi := (theta - math.Asin(alignFactor*D/d))
	return phi
}

func FindLine(p1, p2 OrderedPair) (float64, float64) {
	m := (p2.y - p1.y) / (p2.x - p1.x)
	b := p1.y - m*p1.x
	return m, b
}

func FindYOnLine(x, m, b float64) float64 {
	return (m*x + b)
}


func (fibre *Fibre) UpdateDirection(phi float64) {
	x := fibre.direction.x
	y := fibre.direction.y
	fibre.direction.x = x*math.Cos(phi) - y*math.Sin(phi)
	fibre.direction.y = x*math.Sin(phi) + y*math.Cos(phi)
}

func (fibre *Fibre) UpdatePosition() {
	magnitude := math.Sqrt(fibre.direction.x*fibre.direction.x + fibre.direction.y*fibre.direction.y)
	fibre.position.x = fibre.pivot.x + fibre.direction.x*0.5*fibre.length/magnitude
	fibre.position.y = fibre.pivot.y + fibre.direction.y*0.5*fibre.length/magnitude
}

func (fibre *Fibre) FindPerpendicularDistance(cell *Cell) float64 {
	// need to find coordinates of pivot point
	A := fibre.pivot.y - fibre.position.y
	B := fibre.position.x - fibre.pivot.x
	C := fibre.pivot.x*fibre.position.y - fibre.position.x*fibre.pivot.y
	numerator := math.Abs(A*cell.position.x + B*cell.position.y + C)
	denominator := math.Sqrt(A*A + B*B)
	return numerator/denominator
}

func ComputeDistance(p1, p2 OrderedPair) float64 {
	xComp := p2.x - p1.x
	yComp := p2.y - p1.y
	return math.Sqrt(xComp*xComp + yComp*yComp)
}
