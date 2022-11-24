package main

type ECM struct {
	width     float64
	stiffness float64
	fibres    []*Fibre
	cells     []*Cell
}

type Cell struct {
	radius, height, speed, integrin, shapeFactor, viscocity float64
	position, projection                                    OrderedPair
}

type Fibre struct {
	length, width float64
	position, direction, pivot OrderedPair
}

type OrderedPair struct {
	x, y float64
}
