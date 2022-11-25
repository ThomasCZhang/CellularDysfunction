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
	var thresh float64            // NEED TO GIVE THIS AN ACTUAL VALUE. Threshold should probably be length of fibre/2.

	for _, fibre := range newECM.fibres {
		nearestCell := fibre.FindNearestCell(newECM.cells) // returns a nearest cell
		fibre.UpdateFibre(nearestCell, newECM.stiffness)
	}

	for _, cell := range newECM.cells {
		cell.UpdateCell(newECM.fibres, thresh, time)
	}
	return newECM
}

// CopyECM creates a deep copy of the given ECM
func CopyECM(currentECM *ECM) *ECM {
	var newECM ECM

	newECM.width = currentECM.width
	newECM.stiffness = currentECM.stiffness

	totalFibres := len(currentECM.fibres)
	totalCells := len(currentECM.cells)

	newECM.fibres = make([]*Fibre, totalFibres)
	newECM.cells = make([]*Cell, totalCells)

	// For fibres
	for i := 0; i < totalFibres; i++ {
		newECM.fibres[i].direction = currentECM.fibres[i].direction
		newECM.fibres[i].length = currentECM.fibres[i].length
		newECM.fibres[i].position = currentECM.fibres[i].position
		newECM.fibres[i].width = currentECM.fibres[i].width
	}

	// For Cells
	for i := 0; i < totalCells; i++ {
		newECM.cells[i].radius = currentECM.cells[i].radius
		newECM.cells[i].height = currentECM.cells[i].height
		newECM.cells[i].speed = currentECM.cells[i].speed
		newECM.cells[i].integrin = currentECM.cells[i].integrin
		newECM.cells[i].shapeFactor = currentECM.cells[i].shapeFactor
		newECM.cells[i].viscocity = currentECM.cells[i].viscocity
		newECM.cells[i].position = currentECM.cells[i].position
		newECM.cells[i].projection = currentECM.cells[i].projection
	}

	return &newECM
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

// FindYOnLine: Calculates the y value resulting from plugging x into "y = m*x + b"
// Input: x (float64) the x-value to use
// m , b (float64) m and b in the line equation "y = mx+b"
func FindYOnLine(x, m, b float64) float64 {
	return (m*x + b)
}
