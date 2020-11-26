package main

import (
	"math"
	"strconv"
	"strings"
)

// Point represents a point in a cartesian coordinate system
type Point struct {
	X int
	Y int
}

// GetManhattanDistance calculates the Manhattan distance between two points
func GetManhattanDistance(fPoint, sPoint Point) int {
	xDistance := int(math.Abs(float64(fPoint.X - sPoint.X)))
	yDistance := int(math.Abs(float64(fPoint.Y - sPoint.Y)))
	return xDistance + yDistance
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

// GetMinManhattanDistance returns the minimal Manhattan distance between
// a set of points and the reference point
func GetMinManhattanDistance(points []Point, reference Point) int {
	var minDistance int

	for i, point := range points {
		tmpDistance := GetManhattanDistance(reference, point)
		if i == 0 || tmpDistance < minDistance {
			minDistance = tmpDistance
		}
	}
	return minDistance
}
