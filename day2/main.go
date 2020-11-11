package main

// import "github.com/elkoiko/day2/intcode"
import (
	"fmt"
	"io/ioutil"

	"github.com/elkoiko/day2/intcode"
)

const inputPath = "./input"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Reading entire input file and storing it into a string
	data, err := ioutil.ReadFile(inputPath)
	check(err)

	intCodes := intcode.GetIntCodes(string(data))
	intcode.ComputeIntCodes(intCodes)
	displayResult(intCodes)
}

func displayResult(intCodes []int) {
	for i, value := range intCodes {
		fmt.Printf("Position [%v]: %v\n", i, value)
	}
}
