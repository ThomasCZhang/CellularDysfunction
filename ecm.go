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
func (e *ECM) UpdateECM(time, timePoint float64, positionArray [][]float64) (float64, *ECM, [][]float64) {
	// range over all cells on ECM
	newECM := e.CopyECM() // makes a deep copy of the ECM
	var thresh float64    // NEED TO GIVE THIS AN ACTUAL VALUE. Threshold should probably be length of fibre/2.
	thresh = 40.0

	for _, fibre := range newECM.fibres {
		nearestCell := fibre.FindNearestCell(newECM.cells) // returns a nearest cell
		if ComputeDistance(nearestCell.position, fibre.position) <= thresh {
			fibre.UpdateFibre(nearestCell, ECMstiffness)
		}
	}

	timePoint += time // update time point by time step

	for i, cell := range newECM.cells {

		cell.UpdateCell(e.cells[i], newECM.fibres, thresh, time)

		// add position and time values to array as string
		newValues := make([]float64, 4)
		newValues[0] = float64(timePoint)
		newValues[1] = float64(cell.label)
		newValues[2] = cell.position.x
		newValues[3] = cell.position.y
		positionArray = append(positionArray, newValues)

	}
	return timePoint, newECM, positionArray
}

// CopyECM creates a deep copy of the given ECM
func (e *ECM) CopyECM() *ECM {
	var newECM ECM

	// newECM.width = e.width
	// newECM.stiffness = e.stiffness

	totalFibres := len(e.fibres)
	totalCells := len(e.cells)

	newECM.fibres = make([]*Fibre, totalFibres)
	newECM.cells = make([]*Cell, totalCells)

	// For fibres
	for i := 0; i < totalFibres; i++ {
		newECM.fibres[i] = e.fibres[i].CopyFibre()
	}

	// For Cells
	for i := 0; i < totalCells; i++ {
		newECM.cells[i] = e.cells[i].CopyCell()
	}

	return &newECM
}
