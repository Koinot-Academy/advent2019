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

	minManhattanDistance := GetMinManhattanDistance(GetCrossingPoints(fWire, sWire), startPoint)
	fmt.Printf("Minimal Manhattan distance is: %v.\n", minManhattanDistance)
}
