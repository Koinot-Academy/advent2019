package main

import "testing"

func TestGetRequiredFuel(t *testing.T) {
	var module Module
	var expectedFuel int
	var fuel int

	module.Mass = 12
	expectedFuel = 2
	fuel = module.GetRequiredFuel()
	if fuel != expectedFuel {
		t.Errorf("Fuel should be %v, got %v", expectedFuel, fuel)
	}

	module.Mass = 14
	expectedFuel = 2
	fuel = module.GetRequiredFuel()
	if fuel != expectedFuel {
		t.Errorf("Fuel should be %v, got %v", expectedFuel, fuel)
	}

	module.Mass = 1969
	expectedFuel = 654
	fuel = module.GetRequiredFuel()
	if fuel != expectedFuel {
		t.Errorf("Fuel should be %v, got %v", expectedFuel, fuel)
	}

	module.Mass = 100756
	expectedFuel = 33583
	fuel = module.GetRequiredFuel()
	if fuel != expectedFuel {
		t.Errorf("Fuel should be %v, got %v", expectedFuel, fuel)
	}

}

func TestGetRequiredFuelWithFuelMass(t *testing.T) {
	var module Module
	var expectedTotalFuel int
	var totalFuel int

	module.Mass = 14
	expectedTotalFuel = 2
	totalFuel = module.GetRequiredFuelWithFuelMass()
	if totalFuel != expectedTotalFuel {
		t.Errorf("Total fuel should be %v, got %v", expectedTotalFuel, totalFuel)
	}

	module.Mass = 1969
	expectedTotalFuel = 966
	totalFuel = module.GetRequiredFuelWithFuelMass()
	if totalFuel != expectedTotalFuel {
		t.Errorf("Total fuel should be %v, got %v", expectedTotalFuel, totalFuel)
	}

	module.Mass = 100756
	expectedTotalFuel = 50346
	totalFuel = module.GetRequiredFuelWithFuelMass()
	if totalFuel != expectedTotalFuel {
		t.Errorf("Total fuel should be %v, got %v", expectedTotalFuel, totalFuel)
	}
}
