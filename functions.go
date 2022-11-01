package main

func SimulateCellMotility() {
	// range over some number of generations
	UpdateECM()
}

func UpdateECM() {
	// range over all cells on ECM
	UpdateCell()
	UpdateFibre()
}

func (cell *Cell) UpdateCell() {
	cell.UpdateProjection()
	cell.UpdatePosition()
}

func (cell *Cell) UpdateProjection() {
	ComputeNoise()
	ComputeDragForce()
	ComputePolarity()
	FindNearestFibre()
}

func (cell *Cell) UpdatePosition() {

}

func (fibre *Fibre) UpdateFibre() {
	fibre.UpdatePosition()
	fibre.UpdateDirection()
}

func (fibre *Fibre) UpdatePosition() {

}

func (fibre *Fibre) UpdateDirection() {

}
