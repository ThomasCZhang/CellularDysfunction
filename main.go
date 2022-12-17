package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	RunWebApp()
}

// RunSimulation: Simulates cells on an ECM matrix for a given number of generations
// Input:
// numGens (int): Number of generations to simulate the ECM.
// numCells (int): Number of cells to put on the ECM.
// numFibres (int): Number of fibres to put on the ECM.
// timeStep (float64): Time passed per generation in hours.
// width (float64): The width and length of the ECM "board".
// cellSpeed (float64): The speed at which cells travel on the ECM.
// stiffness (float64): The stiffness of the ECM matrix.
func RunSimulation(numGens, numCells, numFibres int, timeStep, width, cellSpeed, stiffness float64) {
	// arguments: number of generations (int), number of cells (int), number of fibres (int)

	fmt.Println("Commands read in successfully.")

	initialECM := InitializeECM(numFibres, numCells, width, cellSpeed, stiffness)

	fmt.Println("ECM initialized. Beginning simulation.")

	start := time.Now()

	timeFrames, positionArray := SimulateCellMotility(initialECM, numGens, timeStep)

	fmt.Printf("Num Gens: %d, Time Step: %4.3f, Num Cells: %d, Num Fibres: %d, "+
		" Stiffness: %4.3f, Cell Speed: %4.3f, Run Time: %s.\n",
		numGens, timeStep, numCells,
		numFibres, stiffness, cellSpeed,
		time.Since(start).Truncate(time.Millisecond))
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
