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
	fmt.Println("Commands read in successfully.")

	initialECM := InitializeECM(numFibres, numCells, width, cellSpeed, stiffness)

	fmt.Println("ECM initialized. Beginning simulation.")

	timeFrames := SimulateCellMotility(initialECM, numGens, timeStep)
	fmt.Printf("Num Gens: %d, Num Fibres: %d, Num Cells: %d. Run Time: %s\n.", numGens, numFibres, numCells, time.Since(start).Truncate(time.Millisecond))

	fmt.Println("Simulation successful! Now drawing ECM.")

	frequency := 1
	canvasWidth := 2000

	imageList := DrawECM(timeFrames, canvasWidth, frequency, 1)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "gifs/CellMigration")
	fmt.Println("GIF drawn.")
}
