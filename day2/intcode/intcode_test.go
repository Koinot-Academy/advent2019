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
