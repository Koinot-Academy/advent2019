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

// getTotalRequiredFuel counts the sum of all required fuel of each module and returns it
func getTotalRequiredFuel(modules []Module) int {
	total := 0

	for _, module := range modules {
		total += module.RequiredFuel
	}
	return total
}

func main() {
	masses := getMasses(massesPath)
	var modules []Module
	var modulesWithFuelMass []Module

	// Create modules for each retrieved mass
	// Each modules has its corresponding fuel required amount
	for _, mass := range masses {
		modules = append(modules, NewModule(mass))
		modulesWithFuelMass = append(modulesWithFuelMass, NewModuleWithFuelMass(mass))
	}
	fmt.Println("Total fuel required:", getTotalRequiredFuel(modules))
	fmt.Println("Total fuel required when counting fuel mass:", getTotalRequiredFuel(modulesWithFuelMass))
}
