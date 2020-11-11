package intcode

import (
	"strconv"
	"strings"
)

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
