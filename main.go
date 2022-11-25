package main

import (
	"fmt"
	"gifhelper"
)

func main() {

	// arguments: number of generations (int), number of cells (int), number of fibres (int)

	numGens := 15
	numCells := 1
	numFibres := 15000
	width := 1000.0   // the dimensions of a square ECM in micrometres
	stiffness := 0.95 // a value between 0 and 1
	cellSpeed := 20.0 // in micrometres/ hour

	timeStep := 3.0

	fmt.Println("Commands read in successfully.")

	initialECM := InitializeECM(numFibres, numCells, width, cellSpeed, stiffness)

	fmt.Println("ECM initialized. Beginning simulation.")

	timeFrames := SimulateCellMotility(initialECM, numGens, timeStep)

	fmt.Println("Simulation successful! Now drawing ECM.")

	frequency := 1

	canvasWidth := 2000
	imageList := DrawECM(timeFrames, canvasWidth, frequency, width)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "CellMigration")
	fmt.Println("GIF drawn.")
}
