package main

import (

	//chart "go-chart-master"
	//"go-chart-master/drawing"
	"encoding/csv"
	"os"
	"strconv"
)

// WriteToFile writes given array to a csv file
func WriteToFile(positionArray [][]float64) {
	stringArray := make([][]string, len(positionArray))
	// convert every value to string
	for index, row := range positionArray { //range over every row
		if row == nil { // skip nil rows
			continue
		}
		stringRow := make([]string, 4)
		stringRow[0] = strconv.FormatFloat(row[0], 'f', 1, 64)
		stringRow[1] = strconv.FormatFloat(row[1], 'f', 1, 64)
		stringRow[2] = strconv.FormatFloat(row[2], 'f', -1, 64)
		stringRow[3] = strconv.FormatFloat(row[3], 'f', -1, 64)
		stringArray[index] = stringRow
	}

	// make a CSV file to record cell positions at every time-point
	outFilePosition, err1 := os.Create("CellPosition.csv")

	if err1 != nil {
		panic("Error creating output csv file for positions.")
	}

	defer outFilePosition.Close()

	positionWriter := csv.NewWriter(outFilePosition)

	err2 := positionWriter.WriteAll(stringArray)

	if err2 != nil {
		panic("Error writing output csv file for positions.")
	}

}

// PlotGraph takes an array of positions in string form [timepoint, cell label, x, y] and plots both individual RMSD and average RMSD across all cells
func PlotGraph(positionArray [][]float64, numCells int) map[float64][][]float64 {

	// separate position arrays by cell
	positionList := SeparateByCell(positionArray, numCells)
	/*
		// Calculate Root Mean-Squared Deviation (RMSD) for every cell at every timepoint
		totalList := make([][]float64, numCells*2)
		for _, cellList := range positionList {
			xValues, yValues := GetRMSD(cellList) // returns time and RMSD values
			totalList = append(totalList, xValues, yValues...)
			PlotIndividualRMSD(xValues, yValues)
		}

		if numCells > 1 { // if there is more than 1 cell, let's look at the average RMSD
			meanXValues, meanYValues := GetMeanRMSD(totalList, numCells) // Calculate mean RMSD over all cells

			PlotMeanRMSD(meanXValues, meanYValues)
		}
	*/
	return positionList
}

// SeparateByCell takes a position array in the form [[timepoint1, cell label, x, y], [timepoint2, cell label, x, y]...] and converts it to [[cell1: timepoint, x, y], [cell2: timepoint, x, y]...]
func SeparateByCell(positionArray [][]float64, numCells int) map[float64][][]float64 {
	cellMap := make(map[float64][][]float64, numCells)
	for i := 1; i <= numCells; i++ { // we will make a new key for every cell
		counter := 0
		for _, row := range positionArray { // range over every row in position array
			if row == nil { // skip nil rows
				continue
			}
			// fmt.Println(row)
			if float64(i+1) == row[1] { // if current row corresponds to desired cell label
				if cellMap[float64(i)] == nil { // the key does not already exist in the map
					cellMap[float64(i)] = make([][]float64, len(positionArray)) // create an empty array
				}
				cellMap[float64(i)][counter] = make([]float64, 3)
				cellMap[float64(i)][counter][0] = row[0] // timepoint
				cellMap[float64(i)][counter][1] = row[2] // x-position
				cellMap[float64(i)][counter][2] = row[3] // y-position
				counter++
			}
		}
	}
	// fmt.Println(cellMap)
	return cellMap
}

/*
func GetRMSD(list [][]float64) ([]float64, []float64) {
	//
}

func GetMeanRMSD(totalList [][]float64, numCells int) ([]float64, []float64) {
	//
}

func PlotIndividualRMSD(xValues, yValues) {
	//
	graph := chart.Chart{
		Title:      "Root Mean Squared Deviation for Cell %.f",
		TitleStyle: GetTitleStyle(),
		Series: []chart.Series{
			Style: chart.Style{ // sets axes
				StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
				FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
			},
			chart.ContinuousSeries{
				XValues: x1Values,
				YValues: y1Values,
			},
		},
	}
}

func PlotMeanRMSD(xValues, yValues []float64) {
	//
	graph := chart.Chart{
		Title:      "Root Mean Squared Deviation for Cell %.f",
		TitleStyle: GetTitleStyle(),
		Series: []chart.Series{
			Style: chart.Style{ // sets axes
				StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
				FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
			},
			chart.ContinuousSeries{
				XValues: x1Values,
				YValues: y1Values,
			},
		},
	}
}

func GetTitleStyle() chart.Style {
	return chart.Style{
		Show:      true,
		FontColor: drawing.ColorBlue,
	}
}
*/
