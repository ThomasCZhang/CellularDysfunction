package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {

	// arguments: number of generations (int), number of cells (int), number of fibres (int)

	numGens := 200
	timeStep := 1.0

	numCells := 6
	numFibres := 7500
	stiffness := 0.95 // a value between 0 and 1
	cellSpeed := 10.0 // in micrometres/ hour

	width := 500.0 // the dimensions of a square ECM in micrometres

	fmt.Println("Commands read in successfully.")

	initialECM := InitializeECM(numFibres, numCells, width, cellSpeed, stiffness)

	fmt.Println("ECM initialized. Beginning simulation.")

	start := time.Now()
	timeFrames := SimulateCellMotility(initialECM, numGens, timeStep)
	fmt.Printf("Num Gens: %d, Num Fibres: %d, Num Cells: %d. Run Time: %s\n.", numGens, numFibres, numCells, time.Since(start).Truncate(time.Millisecond))

	fmt.Println("Simulation successful! Now drawing ECM.")

	frequency := 1
	canvasWidth := 2000

	imageList := DrawECM(timeFrames, canvasWidth, frequency, 1)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "CellMigration")
	fmt.Println("GIF drawn.")
}
