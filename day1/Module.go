package main

import "math"

// Module represents a module with its mass and required fuel
type Module struct {
	Mass         int
	RequiredFuel int
}

// GetRequiredFuel calculates the required fuel amount given m.Mass
func (m Module) GetRequiredFuel() int {
	return int(math.Floor(float64(m.Mass/3)) - 2)
}
