package main

func SimulateCellMotility() {
	// range over some number of generations
	UpdateECM()
}

func UpdateECM() {
	// range over all cells on ECM
	UpdateCell()
}

func UpdateCell() {
	UpdateAcceleration()
	UpdateVelocity()
	UpdatePosition()
}

func UpdateAcceleration() {

}

func UpdateVelocity() {

}

func UpdatePosition() {

}
