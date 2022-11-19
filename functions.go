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
		fibre.UpdateFibre(nearestCell)
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
// Input: Nearest cell
func (fibre *Fibre) UpdateFibre(cell *Cell) {

	// phi (angle of rotation) = theta (angle of cell from pivot) - arcsin[(1 - 0.1*integrins*(1-stiffness)*perpendicular distance D) / hypotenuse]

	d := fibre.FindHypotenuse(cell) // find the hypotenuse of the cell to the pivot point of the fibre
	D := fibre.FindPerpendicularDistance(cell) // find the perpendicular distance of the cell to the fibre
	theta := ComputeTheta(d,D) // theta = arcsin(d / D)

	phi := ComputePhi(theta,d,D,cell) // compute angle of rotation
	fibre.UpdateDirection(phi) // updates direction vector of fibre using phi

}


func (fibre *Fibre) UpdateDirection() {

}

func FindHypotenuse() {

}

func FindPerpendicularDistance() {

}
