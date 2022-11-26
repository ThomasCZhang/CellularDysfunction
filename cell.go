package main

import (
	"math"
	"math/rand"
)

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
	nearbyFibres := cell.FindNearbyFibres(threshold, fibres)    // returns a slice of nearest fibres within a certain threshold distance
	cell.projection = cell.CalculateNewProjection(nearbyFibres) // Normalized net force acting on the cell from all nearby fibres

	var changeMagnitude float64
	var indexMin int
	// Then we will loop through the fibres again to find the smallest change in projectiono
	for index, fibre := range nearbyFibres {
		delta_magnitude := FindAngleChange(cell.projection, fibre.direction)
		if index == 0 { // set default magnitude of change to first one
			changeMagnitude = math.Abs(delta_magnitude)
		} else if math.Abs(delta_magnitude) > changeMagnitude { // find the minimum magnitude of change
			changeMagnitude = delta_magnitude
			indexMin = index
		}
	}
	if len(nearbyFibres) > 0 {
		cell.UpdateProjection(nearbyFibres[indexMin].direction) // Update the projection vector
	}
	cell.UpdatePosition(time)
}

// FindNearbyFibres: Finds all fibres who's centers are within some threshold distance
// from the center of a cell.
// Input: currCell (*Cell) Pointer to a cell object. Checking distance of fibre's center
// to this cell's center.
// threshold (float64): The max distance in which a fibre can be considered "nearby"
// fibres ([]*Fibre) a slice of pointers to Fibre objects. These are the fibres that are in the ECM.
func (currCell *Cell) FindNearbyFibres(threshold float64, fibres []*Fibre) []*Fibre {
	var nearbyFibres []*Fibre

	for i := 0; i < len(fibres); i++ {
		if ComputeDistance(currCell.position, fibres[i].position) < threshold {
			nearbyFibres = append(nearbyFibres, fibres[i])
		}
	}

	return nearbyFibres
}

// CalculateNewProjection: Calculates the new projection vector of the cell based on nearby fibres.
// Input:
// c (*Cell) a pointer to the Cell object being acted upon
// fibres ([]*Slice) a slice of pointers to nearby fibres
// Output:
// (OrderedPair) The new normalized projection vector of the cell as an OrderedPair.
func (c *Cell) CalculateNewProjection(fibres []*Fibre) OrderedPair {
	var netForce OrderedPair
	if len(fibres) > 0 {
		netForce = c.CalculateNetForce(fibres)
		netForce.Normalize()
	}
	return netForce
}

// CalculateNetForce: Calculates the net force acting on the cell from all nearby fibres.
// Input:
// c (*Cell) a pointer to the Cell object being acted upon
// fibres ([]*Slice) a slice of pointers to nearby fibres
// Output:
// (OrderedPair) The net force acting on the cell as an OrderedPair.
func (c *Cell) CalculateNetForce(fibres []*Fibre) OrderedPair {
	var netForce OrderedPair
	for _, val := range fibres {
		sign := 1.0
		if DotProduct2D(c.projection, val.direction) < 0 {
			sign *= -1
		}
		noise := sign * (1 + rand.NormFloat64())
		projectionVector := ProjectVector(c.projection, MultiplyVectorByConstant2D(val.direction, noise))
		netForce.x += projectionVector.x
		netForce.y += projectionVector.y
	}
	return netForce
}

// FindAngleChange: Calculates the cosine of the angle between two vectors and
// returns the value as a float64. If the angle is > pi/2 radians, then takes
// then takes the negative change instead.
// Input:
// v1, v2 (OrderedPair) The two vectors that will be compared.
func FindAngleChange(v1, v2 OrderedPair) float64 {
	theta := CalculateAngleBetweenVectors2D(v1, v2)
	if theta > math.Pi/2 {
		theta -= math.Pi
	}
	return theta
}

// UpdateProjection: Updates the projection vector of a cell. (Direction of polarity)
// Input: currCell (*Cell) A pointer to the Cell object being updated.
// newProjection (OrderedPair), the new projection vector of the cell.
func (c *Cell) UpdateProjection(newProjection OrderedPair) {
	if DotProduct2D(c.projection, newProjection) < 0 {
		newProjection = MultiplyVectorByConstant2D(newProjection, -1.0)
	}
	// Update the projection vector of the cell
	c.projection.x = newProjection.x
	c.projection.y = newProjection.y
}

// UpdatePosition uses the updated projection vector to change the position of the cell.
func (currCell *Cell) UpdatePosition(time float64) {
	// The postion of the cell using the velocity and timeStep

	var drag OrderedPair
	// Calculate the drag force using the projection vectors from the fibres
	drag = currCell.ComputeDragForce()
	drag.Normalize()
	// Calculte the new position
	currCell.position.x += (drag.x) * CellSpeed * time
	currCell.position.y += (drag.y) * CellSpeed * time
	if currCell.position.x < 0 {
		currCell.position.x += ECMwidth
	} else if currCell.position.x > ECMwidth {
		currCell.position.x -= ECMwidth
	}
	if currCell.position.y < 0 {
		currCell.position.y += ECMwidth
	} else if currCell.position.y > ECMwidth {
		currCell.position.y -= ECMwidth
	}
}

// ComputeDragForce: Computes the drag force acting on a cell by all nearby fibres.
// Input: currCell (*Cell) a pointer to the Cell object.
// Output: (OrderedPair) The drag force.
func (currCell *Cell) ComputeDragForce() OrderedPair {
	// F = speed x shape factor (c) x fluid viscosity (n) x projection vector + noise
	var Drag, FinalForce OrderedPair
	var Noise float64

	// Calculate the noise
	Noise = rand.Float64() * 1.0

	// Compute the drag force
	Drag.x = CellSpeed * currCell.shapeFactor * currCell.viscocity * currCell.projection.x
	Drag.y = CellSpeed * currCell.shapeFactor * currCell.viscocity * currCell.projection.y

	// Add the noise to the drag force
	FinalForce.x = Drag.x + Noise
	FinalForce.y = Drag.y + Noise

	return FinalForce
}

// CopyCell: Returns a pointer to a copy of a Cell object.
func (c *Cell) CopyCell() *Cell {
	var newCell Cell
	newCell.radius = c.radius
	newCell.height = c.height
	// newCell.speed = c.speed
	newCell.integrin = c.integrin
	newCell.shapeFactor = c.shapeFactor
	newCell.viscocity = c.viscocity
	newCell.position = c.position
	newCell.projection = c.projection
	return &newCell
}

// Not quite sure what this is for.
// FindProjection finds the new projection vector caused by a fibre on the given cell
// Input: Current cell and fibre
// Output: The new projection vector of the cell caused by the fibre based on the direction of the fibre and random noise.
// func FindProjection(cell *Cell, fibre *Fibre) OrderedPair {
// 	var newProjection OrderedPair

// 	newProjection.x = cell.projection.x + fibre.projection.x
// 	newProjection.y = cell.projection.y + fibre.projection.y

// 	return newProjection
// }
