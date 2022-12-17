package main

var ECMwidth float64 = 500.0 // uM
var ECMstiffness float64 = 0.95
var CellSpeed float64 = 10.0 // uM per second

type ECM struct {
	// width     float64
	// stiffness float64
	fibres []*Fibre
	cells  []*Cell
}

type Cell struct {
	radius, height, integrin, shapeFactor, viscocity float64
	position, projection                             OrderedPair
	perimeterVertices                                []OrderedPair
	springs                                          []PseudoSpring
	label                                            int
}

type Fibre struct {
	length, width   float64
	position, pivot OrderedPair
	direction       OrderedPair
}

type OrderedPair struct {
	x, y float64
}

type PseudoSpring struct {
	end1, end2 *OrderedPair // Ordered Pair objects representing the ends of the spring.
	x0         float64      // the resting length of the "spring"
}
