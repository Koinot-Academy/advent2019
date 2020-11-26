package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var paths []string
	var scanner *bufio.Scanner
	var fWire, sWire Wire
	startPoint := Point{1, 1}
	var intersections []Point

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}
	file.Close()

	// Creating wires from paths
	fWire = NewWire(GetPointsFromPath(paths[0], startPoint))
	sWire = NewWire(GetPointsFromPath(paths[1], startPoint))
	// Gettings crossing points between wires
	intersections = GetIntersections(fWire, sWire)

	// Minimum Manhattan distance
	minManhattanDistance := GetMinManhattanDistance(intersections, startPoint)
	fmt.Printf("Minimal Manhattan distance is: %v.\n", minManhattanDistance)

	var minCombinedSteps int
	// Getting fewer combined steps between each intersection
	for i, point := range intersections {
		totalSteps := fWire.GetNumberOfSteps(point) + sWire.GetNumberOfSteps(point)
		if i == 0 || totalSteps < minCombinedSteps {
			minCombinedSteps = totalSteps
		}
	}
	fmt.Printf("Fewer steps to reach an intersection is: %v\n", minCombinedSteps)
}
