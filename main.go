package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	RunWebApp()
}

func RunSimulation(numGens, numCells, numFibres int, timeStep, width, cellSpeed, stiffness float64) {
	// arguments: number of generations (int), number of cells (int), number of fibres (int)

	fmt.Println("Commands read in successfully.")

	initialECM := InitializeECM(numFibres, numCells, width, cellSpeed, stiffness)

	fmt.Println("ECM initialized. Beginning simulation.")

	start := time.Now()

	timeFrames, positionArray := SimulateCellMotility(initialECM, numGens, timeStep)

	fmt.Printf("Num Gens: %d, Num Fibres: %d, Num Cells: %d, Run Time: %s, "+
		"Time Step: %4.3f, Cell Speed: %4.3f, Stiffness: %4.3f.\n",
		numGens, numFibres, numCells, time.Since(start).Truncate(time.Millisecond),
		timeStep, cellSpeed, stiffness)
	// write data to files
	WriteToFile(positionArray)

	// generate graph of mean-squared deviation from results
	PlotGraph(positionArray, numCells)

	fmt.Println("Simulation successful! Now drawing ECM.")

	frequency := 1
	canvasWidth := 2000

	imageList := DrawECM(timeFrames, canvasWidth, frequency, 1)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "gifs/CellMigration")
	fmt.Println("GIF drawn.")
}
