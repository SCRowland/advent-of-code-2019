package main

// calculate fuel required
func calculateFuel(moduleMass int) int {
	// Fuel required to launch a given module is based on its mass.
	// Specifically, to find the fuel required for a module,
	// take its mass, divide by three, round down, and subtract 2.
	oneThirdModuleMass := int(moduleMass / 3)
	minusTwo := oneThirdModuleMass - 2

	if minusTwo < 0 {
		return 0
	}

	return minusTwo
}

// calculate fuel required to carry fuel
func calculateFuelFuel(fuelMass int) int {
	runningFuelTotal := 0

	fuelRequired := calculateFuel(fuelMass)
	if fuelRequired < 0 {
		return 0
	}

	runningFuelTotal += fuelRequired

	if fuelRequired > 0 {
		runningFuelTotal += calculateFuelFuel(fuelRequired)
	}

	return runningFuelTotal
}
