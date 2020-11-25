package main

import (
	"reflect"
	"testing"
)

func TestGetManhattanDistance(t *testing.T) {
	var expected, actual int

	expected = 6
	actual = GetManhattanDistance(Point{1, 1}, Point{4, 4})
	if actual != expected {
		t.Errorf("Got %v distance, expected %v.", actual, expected)
	}
}

func TestGetPointsFromPath(t *testing.T) {
	var path string
	start := Point{1, 1}
	var expected, actual []Point

	path = "R8,U5,L5,D3"
	expected = []Point{start, {9, 1}, {9, 6}, {4, 6}, {4, 3}}
	actual = GetPointsFromPath(path, start)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrong points. Got %v expected %v", actual, expected)
	}

	path = "U7,R6,D4,L4"
	expected = []Point{start, {1, 8}, {7, 8}, {7, 4}, {3, 4}}
	actual = GetPointsFromPath(path, start)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrong points. Got %v expected %v", actual, expected)
	}

	// Invalid offsets
	path = "R8@è,U5,L5,D3"
	expected = make([]Point, 0)
	actual = GetPointsFromPath(path, start)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrong points. Got %v expected %v", actual, expected)
	}
}

func TestGetMinManhattanDistance(t *testing.T) {
	var expected, actual int
	referencePoint := Point{1, 1}

	expected = 159
	actual = GetMinManhattanDistance([]Point{{159, -11}, {147, 47}, {156, 5}, {156, 12}}, referencePoint)
	if actual != expected {
		t.Errorf("Wrong distance. Got %v, expected %v.", actual, expected)
	}

	expected = 135
	actual = GetMinManhattanDistance([]Point{{108, 48}, {125, 12}, {158, 19}, {108, 72}, {108, 52}}, referencePoint)
	if actual != expected {
		t.Errorf("Wrong distance. Got %v, expected %v.", actual, expected)
	}

}
