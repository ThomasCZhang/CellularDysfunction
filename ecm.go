package main

/*
type ECM struct {
	width     float64
	stiffness float64
	fibres    []*Fibre
	cells     []*Cell
}
*/

// UpdateECM takes a current ECM object and updates it by the given time step.
// Input: currentECM and a time step
// Output: A new ECM object with  updated cell and fibre positions
func (e *ECM) UpdateECM(time float64) *ECM {
	// range over all cells on ECM
	newECM := e.CopyECM() // makes a deep copy of the ECM
	var thresh float64    // NEED TO GIVE THIS AN ACTUAL VALUE. Threshold should probably be length of fibre/2.
	thresh = 40.0

	for _, fibre := range newECM.fibres {
		nearestCell := fibre.FindNearestCell(newECM.cells) // returns a nearest cell
		if ComputeDistance(nearestCell.position, fibre.position) <= thresh {
			fibre.UpdateFibre(nearestCell, newECM.stiffness)
		}
	}

	for _, cell := range newECM.cells {
		cell.UpdateCell(newECM.fibres, thresh, time)
	}
	return newECM
}

// CopyECM creates a deep copy of the given ECM
func (e *ECM) CopyECM() *ECM {
	var newECM ECM

	newECM.width = e.width
	newECM.stiffness = e.stiffness

	totalFibres := len(e.fibres)
	totalCells := len(e.cells)

	newECM.fibres = make([]*Fibre, totalFibres)
	newECM.cells = make([]*Cell, totalCells)

	// For fibres
	for i := 0; i < totalFibres; i++ {
		var tempFibre Fibre
		newECM.fibres[i] = &tempFibre
		tempFibre.direction = e.fibres[i].direction
		tempFibre.length = e.fibres[i].length
		tempFibre.position = e.fibres[i].position
		tempFibre.width = e.fibres[i].width
	}

	// For Cells
	for i := 0; i < totalCells; i++ {
		var tempCell Cell
		newECM.cells[i] = &tempCell
		tempCell.radius = e.cells[i].radius
		tempCell.height = e.cells[i].height
		tempCell.speed = e.cells[i].speed
		tempCell.integrin = e.cells[i].integrin
		tempCell.shapeFactor = e.cells[i].shapeFactor
		tempCell.viscocity = e.cells[i].viscocity
		tempCell.position = e.cells[i].position
		tempCell.projection = e.cells[i].projection
	}

	return &newECM
}
