package intcode

import (
	"reflect"
	"testing"
)

func TestGetIntCodes(t *testing.T) {
	input := "1,9,10,3,2,3,11,0,99,30,40,50"
	expected := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	output := GetIntCodes(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Got different output: %v, expected: %v", output, expected)
	}

	// Int codes with non numeric values
	input = "1,nn,10,3,2,3,11,0,99,30,@#,50"
	expected = nil

	output = GetIntCodes(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Got different output: %v, expected: %v", output, expected)
	}
}

func TestProcessOpCode(t *testing.T) {
	var result error

	_, result = ProcessOpCode(0, 1, 1)
	if result == nil {
		t.Error("Was expecting non nil error for opCode 0.")
	}

	_, result = ProcessOpCode(1000, 1, 1)
	if result == nil {
		t.Error("Was expecting non nil error for opCode out of range.")
	}

	_, result = ProcessOpCode(1, 1, 1)
	if result != nil {
		t.Error("Was expecting nil error for correct opCode.")
	}
}

func TestComputeIntCodes(t *testing.T) {
	var input []int
	var expected []int

	input = []int{1, 0, 0, 0, 99}
	expected = []int{2, 0, 0, 0, 99}
	ComputeIntCodes(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expecting %v, got %v", expected, input)
	}

	input = []int{2, 3, 0, 3, 99}
	expected = []int{2, 3, 0, 6, 99}
	ComputeIntCodes(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expecting %v, got %v", expected, input)
	}

	input = []int{2, 4, 4, 5, 99, 0}
	expected = []int{2, 4, 4, 5, 99, 9801}
	ComputeIntCodes(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expecting %v, got %v", expected, input)
	}

	input = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	ComputeIntCodes(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expecting %v, got %v", expected, input)
	}
}
