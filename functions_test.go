package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func testFindPerpendicularDistance(t *testing.T) {
	type test struct {
		cell   Cell
		fibre  Fibre
		answer float64
	}

	val, err := os.Getwd()
	if err != nil {
		panic("Unable to find working directory.")
	}
	inputDirectory := val + "\\tests\\Input\\PerpendicularDistance\\"
	inputFiles := ReadFilesFromDirectory(inputDirectory)

	outputDirectory := val + "\\tests\\Output\\PerpendicularDistance\\"
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	if len(inputFiles) != len(outputFiles) {
		panic("Number of test inputs does not equal number of test outputs.")
	}

	trials := make([]test, len(inputFiles))
	for i := 0; i < len(inputFiles); i++ {
		ordPairs := ReadOrderedPairsFromDirectory(inputDirectory, inputFiles[i])
		trials[i].cell.position = ordPairs[0]
		trials[i].fibre.position = ordPairs[1]
		trials[i].fibre.pivot = ordPairs[2]
		answers := ReadFloatsFromDirectory(outputDirectory, outputFiles[i])
		trials[i].answer = answers[0][0]
	}

	for _, trial := range trials {
		solution := trial.fibre.FindPerpendicularDistance(&trial.cell)
		if RoundFloat(solution, 3) != RoundFloat(trial.answer, 3) {
			t.Errorf("Error: When the input is:\n"+
				"fibre pivot = (%.3e, %.3e)\n"+
				"fibre center = (%.3e, %.3e)\n"+
				"cell center = (%.3e, %.3e)\n "+
				"The distance should be %.3e, however the function returned %.3e.",
				trial.fibre.pivot.x, trial.fibre.pivot.y,
				trial.fibre.position.x, trial.fibre.position.y,
				trial.cell.position.x, trial.cell.position.y,
				trial.answer, solution)
		} else {
			fmt.Printf("Correct: When the input is:\n"+
				"fibre pivot = (%.3e, %.3e)\n"+
				"fibre center = (%.3e, %.3e)\n"+
				"cell center = (%.3e, %.3e)\n "+
				"The distance is %.3e.",
				trial.fibre.pivot.x, trial.fibre.pivot.y,
				trial.fibre.position.x, trial.fibre.position.y,
				trial.cell.position.x, trial.cell.position.y,
				trial.answer)
		}
	}

}

func ReadOrderedPairsFromDirectory(directory string, inputFile os.DirEntry) []OrderedPair {
	floatSlice := ReadFloatsFromDirectory(directory, inputFile)
	ordPair := make([]OrderedPair, len(floatSlice))
	for i := range ordPair {
		ordPair[i].x = floatSlice[i][0]
		ordPair[i].y = floatSlice[i][1]
	}
	return ordPair
}

// Reads a text file containing numbers separated by spaces and returns the numbers in a 2 dimensional slice of ints.
// Input:
// directory (string) The location of the text file.
// inputFile (os.DirEntry) information on the file to be read.
// Output:
// ([][]int) A 2d slice of float64.
func ReadIntsFromFile(directory string, inputFile os.DirEntry) [][]int {
	fileName := inputFile.Name()

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	var intSlice [][]int
	for i, line := range inputLines {
		lineValues := strings.Split(line, " ")
		intSlice[i] = make([]int, len(lineValues))
		for j, value := range lineValues {
			intSlice[i][j], err = strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error in ReadIntsFromFile!")
				fmt.Println(err)
				panic("Program Terminated!")
			}
		}
	}
	return intSlice
}

// Reads a text file containing numbers separated by spaces and returns the numbers in a 2 dimensional slice of floats.
// Input:
// directory (string) The location of the text file.
// inputFile (os.DirEntry) information on the file to be read.
// Output:
// ([][]float64) A 2d slice of float64.
func ReadFloatsFromDirectory(directory string, inputFile os.DirEntry) [][]float64 {
	fileName := inputFile.Name()

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	var floatSlice [][]float64
	for i, line := range inputLines {
		lineValues := strings.Split(line, " ")
		floatSlice[i] = make([]float64, len(lineValues))
		for j, value := range lineValues {
			floatSlice[i][j], err = strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Println("Error in ReadIntsFromFile!")
				fmt.Println(err)
				panic("Program Terminated!")
			}
		}
	}
	return floatSlice
}

func ReadFilesFromDirectory(directory string) []os.DirEntry {
	dirContents, err := os.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}

func AssertEqualAndNonzero(length0, length1 int) {
	if length0 == 0 {
		panic("No files present in input directory.")
	}
	if length1 == 0 {
		panic("No files present in output directory.")
	}
	if length0 != length1 {
		panic("Number of files in directories doesn't match.")
	}
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func StringToFloat(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic("Error in StringToFloat, unable to parse " + str)
	}
	return val
}
