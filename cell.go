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
func (cell *Cell) UpdateCell(oldCell *Cell, fibres []*Fibre, threshold float64, time float64) {
	// This was for the attempt to model the cell as a soft body object. It didn't work unfortunately.
	// cell.UpdateShape(oldCell)

	// range over all fibres and compute projection vectors caused by all fibres on the cell
	// to make this easier, we can only pick fibres that are within a certain critical distance to the cell
	nearbyFibres := cell.FindNearbyFibres(threshold, fibres)    // returns a slice of nearest fibres within a certain threshold distance
	cell.projection = cell.CalculateNewProjection(nearbyFibres) // Normalized net force acting on the cell from all nearby fibres

	var changeMagnitude float64
	var indexMin int
	// Then we will loop through the fibres again to find the smallest change in projection
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

// UpdateProjection: Updates the projection vector of a cell. (Direction of polarity)
// Input: currCell (*Cell) A pointer to the Cell object being updated.
// newProjection (OrderedPair), the new projection vector of the cell.
func (c *Cell) UpdateProjection(newProjection OrderedPair) {
	if DotProduct2D(c.projection, newProjection) < 0 { // Making sure the cell doesn't do a 180.
		newProjection = MultiplyVectorByConstant2D(newProjection, -1.0)
	}
	// Update the projection vector of the cell
	c.projection.x = newProjection.x
	c.projection.y = newProjection.y
}

// UpdatePosition uses the updated projection vector to change the position of the cell.
// Input: The time step (in hours).
func (currCell *Cell) UpdatePosition(time float64) {
	// The postion of the cell using the velocity and timeStep

	var drag OrderedPair
	// Calculate the drag force using the projection vectors from the fibres
	drag = currCell.ComputeDragForce()
	drag.Normalize()

	// Calculte the new position
	currCell.position.x += (drag.x) * CellSpeed * time
	currCell.position.y += (drag.y) * CellSpeed * time

	// Putting the cells on a torus
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
	Noise = rand.Float64()*2.0 - 1.0

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
	newCell.perimeterVertices = make([]OrderedPair, len(c.perimeterVertices))
	newCell.springs = make([]PseudoSpring, len(c.springs))
	for i := range c.perimeterVertices {
		newCell.perimeterVertices[i].x = c.perimeterVertices[i].x
		newCell.perimeterVertices[i].y = c.perimeterVertices[i].y

		// spring between perimeter verticies
		end1 := &(newCell.perimeterVertices[i])
		var end2 *OrderedPair
		if i == len(c.perimeterVertices)-1 {
			end2 = &(newCell.perimeterVertices[0])
		} else {
			end2 = &(newCell.perimeterVertices[i+1])
		}
		newCell.springs[i].end1 = end1
		newCell.springs[i].end2 = end2
		newCell.springs[i].x0 = c.springs[i].x0
		// spring between perimeter and center
		newCell.springs[len(c.perimeterVertices)+i].end1 = end1
		newCell.springs[len(c.perimeterVertices)+i].end2 = &(newCell.position)
		newCell.springs[len(c.perimeterVertices)+i].x0 = c.springs[len(c.perimeterVertices)+i].x0

	}
	return &newCell
}

// UpdatesShape updates the perimeter vertices of a Cell object.
// Input:
// c (*Cell): The cell whose shape is updated by this function.
// oldCell (*Cell): The cell that represents "c" from the previous generation.
func (c *Cell) UpdateShape(oldCell *Cell) {
	for i := range c.springs {
		movementMagnitude := oldCell.springs[i].CalculateSpringMoveMagnitude()
		// Calculate the direction the ends of the spring should move
		var delta OrderedPair
		delta.x = oldCell.springs[i].end2.x - oldCell.springs[i].end1.x
		delta.y = oldCell.springs[i].end2.y - oldCell.springs[i].end1.y
		deltaMag := delta.Magnitude()
		delta.x *= movementMagnitude / deltaMag
		delta.y *= movementMagnitude / deltaMag
		c.springs[i].end1.x += delta.x
		c.springs[i].end1.y += delta.y
		c.springs[i].end2.x -= delta.x
		c.springs[i].end2.y -= delta.y
	}
}

// CalculateSpringMoveMagnitude: Calculates how much the "weights" at each end of a
// spring should move towards each other. Neither end is thought to be fixed. Note: This is a pseudospring,
// so it doesn't update using the force equation (F = kx).
// Input: scalingFactor (float64) a factor used to determine how much the ends of the spring should move.
func (s PseudoSpring) CalculateSpringMoveMagnitude() float64 {
	dx := s.end1.x - s.end2.x
	dy := s.end1.y - s.end2.y
	distance := math.Sqrt(dx*dx + dy*dy)
	numerator := distance - s.x0
	magnitude := math.Abs(numerator) * numerator / (4 * s.x0)
	// if numerator > s.x0 {
	// 	magnitude = numerator / 2
	// }
	// if magnitude < 0 {
	// 	fmt.Println(magnitude, distance, s.x0)
	// }
	// fmt.Println(magnitude)
	return magnitude
}

// FindClosestPerimVertex Finds the closest vertex in the direction of the drag force.
// Input: drag (OrderedPair) drag is the force acting on the cell
// currCell (*Cell) is the current cell being analyzed.
func (currCell *Cell) FindClosestPerimVertex(drag OrderedPair) *OrderedPair {
	// Find the perimeter vertex that's in the closest direction as drag.
	var closestPerim *OrderedPair
	maxDotProd := 0.0
	for i := range currCell.perimeterVertices {
		var delta OrderedPair
		delta.x = currCell.perimeterVertices[i].x - currCell.position.x
		delta.y = currCell.perimeterVertices[i].y - currCell.position.y
		delta.Normalize()
		if DotProduct2D(delta, drag) > maxDotProd {
			closestPerim = &(currCell.perimeterVertices[i])
			maxDotProd = DotProduct2D(delta, drag)
		}
	}
	return closestPerim
}
