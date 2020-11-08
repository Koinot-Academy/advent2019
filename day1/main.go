package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const massesPath = "./module.masses"

// getMasses retrieves masses from a file located at filePath.
// File format should be one line equals one mass
// It returns a slice of masses. If an error happens, a nil slice is returned
func getMasses(filePath string) []int {
	var masses []int
	file, err := os.Open(filePath)

	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil
		}
		masses = append(masses, mass)
	}

	if scanner.Err() != nil {
		return nil
	}
	return masses
}

func main() {
	module := Module{12, 0}

	module.RequiredFuel = module.GetRequiredFuel()
	fmt.Println("Masses", getMasses(massesPath))
}
