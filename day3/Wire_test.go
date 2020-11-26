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

func TestGetIntersections(t *testing.T) {
	fWirePoints := []Point{{1, 1}, {9, 1}, {9, 6}, {4, 6}, {4, 3}}
	sWirePoints := []Point{{1, 1}, {1, 8}, {7, 8}, {7, 4}, {3, 4}}
	fWire := NewWire(fWirePoints)
	sWire := NewWire(sWirePoints)
	expected := []Point{{7, 6}, {4, 4}}

	actual := GetIntersections(fWire, sWire)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Wrong crossing points. Expected %v, got %v", expected, actual)
	}
}

func TestGetNumberOfSteps(t *testing.T) {
	var expected, actual int
	var wire Wire

	wire = Wire{Segments: []Segment{
		{Point{1, 1}, Point{9, 1}},
		{Point{9, 1}, Point{9, 6}},
		{Point{9, 6}, Point{4, 6}},
		{Point{4, 6}, Point{4, 3}},
	}}
	expected = 20
	actual = wire.GetNumberOfSteps(Point{4, 4})
	if actual != expected {
		t.Errorf("Wrong number of steps. Expected %v, got %v", expected, actual)
	}

	wire = Wire{Segments: []Segment{
		{Point{1, 1}, Point{1, 8}},
		{Point{1, 8}, Point{7, 8}},
		{Point{7, 8}, Point{7, 4}},
		{Point{7, 4}, Point{3, 4}},
	}}
	expected = 20
	actual = wire.GetNumberOfSteps(Point{4, 4})
	if actual != expected {
		t.Errorf("Wrong number of steps. Expected %v, got %v", expected, actual)
	}

	wire = Wire{Segments: []Segment{
		{Point{1, 1}, Point{9, 1}},
		{Point{9, 1}, Point{9, 6}},
		{Point{9, 6}, Point{4, 6}},
		{Point{4, 6}, Point{4, 3}},
	}}
	expected = 15
	actual = wire.GetNumberOfSteps(Point{7, 6})
	if actual != expected {
		t.Errorf("Wrong number of steps. Expected %v, got %v", expected, actual)
	}

	wire = Wire{Segments: []Segment{
		{Point{1, 1}, Point{1, 8}},
		{Point{1, 8}, Point{7, 8}},
		{Point{7, 8}, Point{7, 4}},
		{Point{7, 4}, Point{3, 4}},
	}}
	expected = 15
	actual = wire.GetNumberOfSteps(Point{7, 6})
	if actual != expected {
		t.Errorf("Wrong number of steps. Expected %v, got %v", expected, actual)
	}
}
