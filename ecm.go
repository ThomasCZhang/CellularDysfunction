package main

import (
	"strconv"
)

// UpdateECM takes a current ECM object and updates it by the given time step.
// Input: currentECM and a time step
// Output: A new ECM object with  updated cell and fibre positions
func (e *ECM) UpdateECM(time, timePoint float64, positionArray [][]string, identityArray [][]string) (float64, *ECM, [][]string, [][]string) {
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

	for _, cell := range newECM.cells {

		cell.UpdateCell(newECM.fibres, thresh, time)

		// add position and time values to array as string
		newValues := make([]string, 3)
		newValues[0] = strconv.FormatFloat(timePoint, 'f', -1, 64)
		newValues[1] = strconv.FormatFloat(cell.position.x, 'f', -1, 64)
		newValues[2] = strconv.FormatFloat(cell.position.y, 'f', -1, 64)
		positionArray = append(positionArray, newValues)

		// add cell label to array
		newLabel := make([]string, 1)
		newLabel[0] = cell.label
		identityArray = append(identityArray, newLabel)
	}
	return timePoint, newECM, positionArray, identityArray
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
