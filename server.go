package main

import (
	"fmt"
	htemplate "html/template"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

// For handling directories
const Plots = "gifs"
const PlotRoot = "/" + Plots + "/"

// Page a small struct to help us use HTML templates
type Page struct {
	Title    string
	Contents htemplate.HTML
}

// RunWebApp: For creating the web app server.
func RunWebApp() {
	fmt.Println("Running Web App.")
	fmt.Println("http://localhost:5000")

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/Inputs/", inputHandler)

	http.Handle(PlotRoot, http.StripPrefix(PlotRoot, http.FileServer(http.Dir("./"+Plots))))
	http.ListenAndServe(":5000", nil)
}

// MainHandler: Handler that loads the html right when server is built.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("inputs.html")
	// t, err := template.ParseFiles("inputs.html")
	if err != nil {
		fmt.Println("ERROR")
		panic("Error in mainHandler.")
	}
	t.Execute(w, nil)
}

// inputHandler: Handler for when someone hits the "submit" button.
func inputHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, val := range r.Form {
		fmt.Println("key", key, "val", val)
	}
	numGens, err := strconv.Atoi(r.Form["numGens"][0])
	if err != nil {
		panic("Failure in inputHandler.")
	}
	numCells, err := strconv.Atoi(r.Form["numCells"][0])
	if err != nil {
		panic("Failure in inputHandler.")
	}
	numFibres, err := strconv.Atoi(r.Form["numFibres"][0])
	if err != nil {
		panic("Failure in inputHandler.")
	}
	timeStep, err := strconv.ParseFloat(r.Form["timeStep"][0], 0)
	if err != nil {
		panic("Failure in inputHandler.")
	}
	stiffness, err := strconv.ParseFloat(r.Form["stiffness"][0], 0)
	if err != nil {
		panic("Failure in inputHandler.")
	}
	cellSpeed, err := strconv.ParseFloat(r.Form["cellSpeed"][0], 0)
	if err != nil {
		panic("Failure in inputHandler.")
	}
	width, err := strconv.ParseFloat(r.Form["width"][0], 0)
	if err != nil {
		panic("Failure in inputHandler.")
	}

	RunSimulation(numGens, numCells, numFibres, timeStep, width, cellSpeed, stiffness)

	t, err := template.ParseFiles("inputs.html")
	if err != nil {
		panic("Error while adding gif to page.")
	}

	var page Page
	plotFile := "CellMigration"
	plotPath := path.Join(Plots, plotFile) + ".out.gif"
	page = Page{
		Title: "ECM Gif",
		Contents: htemplate.HTML(fmt.Sprintf(
			"<img src='/%s' class='rounded' alt='skew' style='width:600px;height:auto;'>", plotPath,
		)),
	}
	t.Execute(w, page)
}
