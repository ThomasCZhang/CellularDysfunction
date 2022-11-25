package main

import "math/rand"

/*
type Cell struct {
	radius, height, speed, integrin, shapeFactor, viscocity float64
	position, projection                                    OrderedPair
}
*/

// UpdateCell finds the new direction of a given cell based on the fibre that causes the least change in direction
// Input: cell object and a list of updated fibres
// Output: cell with updated projection and position
func (cell *Cell) UpdateCell(fibres []*Fibre, threshold float64, time float64) {
	// range over all fibres and compute projection vectors caused by all fibres on the cell
	// to make this easier, we can only pick fibres that are within a certain critical distance to the cell

	nearestFibres := cell.FindNearbyFibres(threshold, fibres) // returns a slice of nearest fibres within a certain threshold distance
	var changeMagnitude float64
	var netProjection OrderedPair
	// First we loop through all the fibres, to find the net projection due to all the NearestFibres
	for _, fibre := range nearestFibres {
		netProjection.x += fibre.projection.x
		netProjection.y += fibre.projection.y
	}

	netProjection.x = netProjection.x / float64(len(nearestFibres))
	netProjection.y = netProjection.y / float64(len(nearestFibres))

	var indexMin int
	// Then we will loop through the fibres again to find the smallest change in projection
	for index, fibre := range nearestFibres {
		newProjection := FindProjection(cell, fibre)
		delta_magnitude := FindMagnitudeChange(newProjection, cell.projection)
		if index == 0 { // set default magnitude of change to first one
			changeMagnitude = delta_magnitude
		}
		if delta_magnitude > changeMagnitude { // find the minimum magnitude of change
			changeMagnitude = delta_magnitude
			indexMin = index
		}
	}
	cell.UpdateProjection(nearestFibres[indexMin].direction) // Update the projection vector
	cell.UpdatePosition(time)
}

// FindNearbyFibres: Finds all fibres who's centers are within some threshold distance
// from teh center of a cell.
// Input: currCell (*Cell) Pointer to a cell object. Checking distance of fibre's center
// to this cell's center.
// threshold (float64): The max distance in which a fibre can be considered "nearby"
// fibres ([]*Fibre) a slice of pointers to Fibre objects. These are the fibres that are in the ECM.
func (currCell *Cell) FindNearbyFibres(threshold float64, fibres []*Fibre) []*Fibre {
	var nearestFibres []*Fibre

	for i := 0; i < len(fibres); i++ {
		if ComputeDistance(currCell.position, fibres[i].position) < threshold {
			nearestFibres = append(nearestFibres, fibres[i])
		}
	}

	return nearestFibres
}

// FindProjection finds the new projection vector caused by a fibre on the given cell
// Input: Current cell and fibre
// Output: The new projection vector of the cell caused by the fibre based on the direction of the fibre and random noise.
func FindProjection(cell *Cell, fibre *Fibre) OrderedPair {
	var newProjection OrderedPair

	newProjection.x = cell.projection.x + fibre.projection.x
	newProjection.y = cell.projection.y + fibre.projection.y

	return newProjection
}

// FindMagnitudeChange: Calculates the cosine of the angle between two vectors and
// returns the value as a float64.
// Input: projectionA, projectionB (OrderedPair) The two vectors that will be compared.
func FindMagnitudeChange(projectionA, projectionB OrderedPair) float64 {
	var cos float64
	// Find the dot product of the two vectors
	dot := (projectionA.x * projectionB.x) + (projectionA.y * projectionB.y)

	// Find the magnitude of the two vectors
	magA := projectionA.Magnitude()
	magB := projectionB.Magnitude()

	// Find the cos theta value
	cos = dot / (magA * magB)

	// return the cos theta value
	return cos
}

// UpdateProjection: Updates the projection vector of a cell. (Direction of polarity)
// Input: currCell (*Cell) A pointer to the Cell object being updated.
// newProjection (OrderedPair), the new projection vector of the cell.
func (currCell *Cell) UpdateProjection(newProjection OrderedPair) {

	// Update the projection vector of the cell
	currCell.projection.x = newProjection.x
	currCell.projection.y = newProjection.y
}

// ComputeDragForce: Computes the drag force acting on a cell by all nearby fibres.
// Input: currCell (*Cell) a pointer to the Cell object.
func (currCell *Cell) ComputeDragForce() OrderedPair {
	// F = speed x shape factor (c) x fluid viscosity (n) x projection vector + noise
	var Drag, FinalForce OrderedPair
	var Noise float64

	// Calculate the noise
	Noise = rand.Float64() * 1.0

	// Compute the drag force
	Drag.x = currCell.speed * currCell.shapeFactor * currCell.viscocity * currCell.projection.x
	Drag.y = currCell.speed * currCell.shapeFactor * currCell.viscocity * currCell.projection.y

	// Add the noise to the drag force
	FinalForce.x = Drag.x + Noise
	FinalForce.y = Drag.y + Noise

	return FinalForce
}

// UpdatePosition uses the updated projection vector to change the position of the cell.
func (currCell *Cell) UpdatePosition(time float64) {
	// The postion of the cell using the velocity and timeStep

	var newPos, drag OrderedPair
	// Calculate the drag force using the projection vectors from the fibres
	drag = currCell.ComputeDragForce()

	// Calculte the new position
	newPos.x = currCell.position.x + (drag.x)*time
	newPos.y = currCell.position.y + (drag.y)*time
}
