package main

import "math"

// Module represents a module with its mass and required fuel
type Module struct {
	Mass         int
	RequiredFuel int
}

// GetRequiredFuelWithFuelMass calculates the required amount of taking into account m.Mass
// and the mass of the added fuel
func (m Module) GetRequiredFuelWithFuelMass() int {
	totalRequiredFuel := m.GetRequiredFuel()

	for requiredFuel := totalRequiredFuel; requiredFuel > 0; {
		requiredFuel = getFuelAmount(requiredFuel)
		if requiredFuel > 0 {
			totalRequiredFuel += requiredFuel
		}
	}
	return totalRequiredFuel
}

// GetRequiredFuel calculates the required fuel amount given m.Mass
func (m Module) GetRequiredFuel() int {
	return getFuelAmount(m.Mass)
}

// getFuelAmount calculates the amount of fuel needed given a mass m
func getFuelAmount(mass int) int {
	return int(math.Floor(float64(mass/3)) - 2)
}

// NewModule creates and initializes a Module with a given mass
// It initializes the required fuel without taking into account the fuel mass
func NewModule(mass int) Module {
	module := Module{mass, 0}

	module.RequiredFuel = module.GetRequiredFuel()
	return module
}

// NewModuleWithFuelMass creates a Module with a given mass
// It initializes the required fuel taking into account the fuel mass
func NewModuleWithFuelMass(mass int) Module {
	module := Module{mass, 0}

	module.RequiredFuel = module.GetRequiredFuelWithFuelMass()
	return module
}
