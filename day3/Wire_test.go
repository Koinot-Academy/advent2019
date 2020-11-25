package main

import (
	"reflect"
	"testing"
)

func TestNewWire(t *testing.T) {
	var points []Point
	var expected, actual Wire

	points = []Point{{1, 1}, {9, 1}, {9, 6}, {4, 6}, {4, 3}}
	expected = Wire{Segments: []Segment{
		{Point{1, 1}, Point{9, 1}},
		{Point{9, 1}, Point{9, 6}},
		{Point{9, 6}, Point{4, 6}},
		{Point{4, 6}, Point{4, 3}},
	}}
	actual = NewWire(points)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Wrong wire initialization. Got %v expected %v.", actual, expected)
	}
}

func TestGetCrossingPoints(t *testing.T) {
	fWirePoints := []Point{{1, 1}, {9, 1}, {9, 6}, {4, 6}, {4, 3}}
	sWirePoints := []Point{{1, 1}, {1, 8}, {7, 8}, {7, 4}, {3, 4}}
	fWire := NewWire(fWirePoints)
	sWire := NewWire(sWirePoints)
	expected := []Point{{7, 6}, {4, 4}}

	actual := GetCrossingPoints(fWire, sWire)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Wrong crossing points. Expected %v, got %v", expected, actual)
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
