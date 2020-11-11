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
