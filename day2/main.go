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

	// Instructions tell to replace those values
	intCodes[1] = 12
	intCodes[2] = 2

	intcode.ComputeIntCodes(intCodes)
	displayResult(intCodes)

	// For Part 2
	valueToFind := 19690720
	data, err = ioutil.ReadFile(inputPath)
	check(err)
	intCodes = intcode.GetIntCodes(string(data))

	noun, verb := intcode.FindInputPair(intCodes, valueToFind, 0)
	fmt.Printf("\nPART TWO\nTo find %v, we must have noun=%v and verb=%v\n", valueToFind, noun, verb)
	fmt.Println("Part two answer is:", 100*noun+verb)
}

func displayResult(intCodes []int) {
	for i, value := range intCodes {
		fmt.Printf("Position [%v]: %v\n", i, value)
	}
}
