package main

import (
	"math"
	"math/rand"
)

// InitializeECM generates a new ECM object
// Input: number of fibres, number of cells, width of ECM, speed of cells, stiffness of matrix
// Output: pointer to ECM object made using given parameters
func InitializeECM(numFibres, numCells int, width, speed float64, stiffness float64) *ECM {

	ECMwidth = width
	ECMstiffness = stiffness
	CellSpeed = speed
	var newECM ECM
	newECM.fibres = InitializeFibres(numFibres, width)
	newECM.cells = InitializeCells(numCells, width)
	return &newECM
}

// InitializeFibres generates an array of identical fibres that only vary in position and direction
// Input: number of fibres and ECM width
// Output: a slice of pointers to distinct fibre objects with unique positions and directions
func InitializeFibres(numFibres int, width float64) []*Fibre {

	FibreArray := make([]*Fibre, numFibres)

	for i := 0; i <= numFibres-1; i++ {

		var newFibre Fibre

		newFibre.length = rand.NormFloat64()*5.0 + 75.0 // the length is normally distributed with a mean of 75 micrometres and sd of 5 micrometres

		newFibre.width = 0.2 // the width is 200nm = 0.2 micrometres

		// place fibres randomly on ECM. This value represents centre of the fibre
		newFibre.position.x = rand.Float64() * width
		newFibre.position.y = rand.Float64() * width

		// randomly assign x-direction and calculate y-direction such that the vector is a unit vector (length = 1)
		newFibre.direction.x = ((rand.Float64() - 0.5) * 2) // some random float in the interval [-1.0, 1.0)
		newFibre.direction.y = GenerateYDirection(newFibre.direction.x)

		FibreArray[i] = &newFibre
	}
	return FibreArray
}

// InitializeCells generates an array of identical cells that only vary in position and projection
// Input: number of cells
// Output: a slice of pointers to distinct cell objects with unique positions and directions
func InitializeCells(numCells int, width float64) []*Cell {

	CellArray := make([]*Cell, numCells)

	for i := 0; i <= numCells-1; i++ {

		var newCell Cell
		newCell.label = i + 1

		newCell.radius = 15.0 // in micrometres
		newCell.height = 2.6  // in micrometres
		// newCell.speed = cellSpeed
		newCell.integrin = 50                                                     // in %
		newCell.shapeFactor = 16.7 * math.Sqrt(0.5*newCell.radius*newCell.height) // In Eqn S3, c = 16.7 * sqrt(0.5 * r * h)
		newCell.viscocity = 100                                                   // in Poise

		// place cell randomly on ECM

		// newCell.position.x = width/4 + rand.Float64()*width/2
		// newCell.position.y = width/4 + rand.Float64()*width/2
		n := 0.125
		newCell.position.x = width*n + rand.Float64()*width*(1-2*n)
		newCell.position.y = width*n + rand.Float64()*width*(1-2*n)

		// generate random direction for cell
		newCell.projection.x = ((rand.Float64() - 0.5) * 2) // some random float in the interval [-1.0, 1.0)
		newCell.projection.y = GenerateYDirection(newCell.position.x)

		CellArray[i] = &newCell
	}
	return CellArray
}

// GenerateYDirection uses the x-direction value to generate a y-direction value such that the resulting direction is a unit vector
// Input: x value of a direction vector
// Output: y value of a direction vector
func GenerateYDirection(xDirection float64) float64 {

	// determine sign of y randomly
	someInt := rand.Intn(2)
	var sign float64
	if someInt%2 == 0 {
		sign = 1.0
	} else {
		sign = -1.0
	}

	// use pythagorean theorem to ensure magnitude of (x,y) is 1
	y := sign * math.Sqrt(1-math.Pow(xDirection, 2))

	return y
}

// InitializePositionArray creates a 2-dimensional array for storing the positions of cells during the simulation
// Input: initial ECM object and number of generations
// Output: A 2-dimensional array where there is a new []float64 array for every cell at every generation, containing the time point value, the cell label as well as the x and y coordinates.
func InitializePositionArray(initialECM *ECM, numGens int) [][]float64 {
	newArray := make([][]float64, (numGens+1)*len(initialECM.cells))
	initialTime := 0.0
	for _, cell := range initialECM.cells {
		values := make([]float64, 4)
		values[0] = initialTime
		values[1] = float64(cell.label)
		values[2] = cell.position.x
		values[3] = cell.position.y
	}
	return newArray
}
