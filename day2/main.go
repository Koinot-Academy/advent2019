package main

// import "github.com/elkoiko/day2/intcode"
import (
	"fmt"
	"io/ioutil"
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
	content := string(data)
	fmt.Println(content)
}
