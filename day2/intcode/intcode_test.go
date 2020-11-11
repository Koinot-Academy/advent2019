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
