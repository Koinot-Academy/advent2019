package intcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var opCodes = [...]func(int, int) int{nil, add, multiply}

/*
* Operations
 */

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

// GetIntCodes transforms strCodes into a slice of integers
// The input must look like: 1,0,0,2
func GetIntCodes(strCodes string) []int {
	var intcodes []int

	codes := strings.Split(strCodes, ",")
	for _, strCode := range codes {
		code, err := strconv.Atoi(strCode)
		if err != nil {
			return nil
		}
		intcodes = append(intcodes, code)
	}
	return intcodes
}

// ProcessOpCode processes the given opCode on firstInput and secInput and returns its result
// If no opCode implementation is found, an error is returned
func ProcessOpCode(opCode int, firstInput int, secInput int) (int, error) {
	var strError = fmt.Sprintf("No implementation has been found for %v.", opCode)

	if opCode >= len(opCodes) {
		return 0, errors.New(strError)
	}

	opCodeImplementation := opCodes[opCode]
	if opCodeImplementation == nil {
		return 0, errors.New(strError)
	}
	return opCodeImplementation(firstInput, secInput), nil
}

// ComputeIntCodes modifies intCodes values following opCodes instructions
func ComputeIntCodes(intCodes []int) {

	for opCodeIndex := 0; ; opCodeIndex += 4 {
		opCode := intCodes[opCodeIndex]
		if opCode == 99 {
			// We halt the program
			return
		}

		firstInputIndex := intCodes[opCodeIndex+1]
		secInputIndex := intCodes[opCodeIndex+2]
		outputPosition := intCodes[opCodeIndex+3]

		output, err := ProcessOpCode(opCode, intCodes[firstInputIndex], intCodes[secInputIndex])
		if err != nil {
			return
		}
		intCodes[outputPosition] = output
	}
}

// FindInputPair tests all possible values for noun (Position 1) and verb (Position 2) between 0 and 99 included
// to get an output at the given position.
// Returns a pair of negative values if no pair give the expected output
func FindInputPair(baseIntCodes []int, output int, position int) (int, int) {
	var noun, verb int
	intCodesCpy := make([]int, len(baseIntCodes))

	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			copy(intCodesCpy, baseIntCodes)
			intCodesCpy[1] = noun
			intCodesCpy[2] = verb
			ComputeIntCodes(intCodesCpy)

			if intCodesCpy[position] == output {
				return noun, verb
			}
		}
	}
	return -1, -1
}
