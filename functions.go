package main

func SimulateCellMotility() {
	// range over some number of generations
	UpdateECM()
}

func UpdateECM() {
	// range over all cells on ECM
	// Place holder. This is just here to get rid of compiler errors.
	var cell Cell
	var fibre Fibre
	////
	cell.UpdateCell()
	fibre.UpdateFibre()
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
	fibre.UpdatePosition() // Maybe we don't need this.
	fibre.UpdateDirection()
	FindHypotenuse()
	FindPerpendicularDistance()
}

func (fibre *Fibre) UpdatePosition() {

}

func (fibre *Fibre) UpdateDirection() {

}

func ComputeNoise() {

}

func ComputeDragForce() {

}

func FindNearestFibre() {

}

func FindHypotenuse() {

}

func FindPerpendicularDistance() {

}

func ComputePolarity() {

}
