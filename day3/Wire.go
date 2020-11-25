package main

import (
	"strconv"
	"strings"
)

// Wire represents multiple segments connected with each other
type Wire struct {
	Segments []Segment
}

// NewWire initializes a wire with a slice of points
func NewWire(points []Point) Wire {
	var wire Wire

	wire.Segments = make([]Segment, 0, len(points)-1)
	for i, pointsLen := 1, len(points); i < pointsLen; i++ {
		wire.Segments = append(wire.Segments, NewSegment(points[i-1], points[i]))
	}
	return wire
}

// GetCrossingPoints returns a slice of points where both wires cross
// Returns nil if the two wires do not cross
func GetCrossingPoints(fWire, sWire Wire) []Point {
	var crossingPoints []Point

	for _, seg1 := range fWire.Segments {
		for _, seg2 := range sWire.Segments {
			// If cross but not in the starting point
			if crossingPoint := seg1.Crosses(seg2); crossingPoint != nil && seg1.FirstPoint.X != 1 && seg1.FirstPoint.Y != 1 {
				crossingPoints = append(crossingPoints, *crossingPoint)
			}
		}
	}
	return crossingPoints
}

// GetPointsFromPath slices path into directions to deduct a set of points
// starting at start point.
// path expects this format: "R100,D10,U30,L87"
// Returns an empty Point slice if any error happens
func GetPointsFromPath(path string, start Point) []Point {
	var points []Point
	var directions []string

	directions = strings.Split(path, ",")
	// Init points slice with capacity for all points counting start one
	points = make([]Point, 1, len(directions)+1)
	points[0] = start
	for i, direction := range directions {
		var xFactor, yFactor int
		previousPoint := points[i]
		offset, err := strconv.Atoi(direction[1:])

		if err != nil {
			return make([]Point, 0)
		}
		// Getting factors for each direction
		switch direction[0] {
		case 'R':
			xFactor = 1
		case 'L':
			xFactor = -1
		case 'D':
			yFactor = -1
		case 'U':
			yFactor = 1
		}
		// One of both factors will have a zero value, so only one axis value will be modified
		points = append(points, Point{previousPoint.X + offset*xFactor, previousPoint.Y + offset*yFactor})
	}
	return points
}
