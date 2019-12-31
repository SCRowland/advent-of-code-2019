package main

import "testing"

var calculateFuelTests = []struct {
	mass int
	fuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestCalculateFuel(t *testing.T) {
	for _, tt := range calculateFuelTests {
		got := calculateFuel(tt.mass)
		if got != tt.fuel {
			t.Errorf("calculateFuel(%d) = %d, want %d", tt.mass, got, tt.fuel)
		}
	}
}

var calculateFuelFuelTests = []struct {
	fuelMass int
	fuel     int
}{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func TestCalculateFuelFuel(t *testing.T) {
	for _, tt := range calculateFuelFuelTests {
		got := calculateFuelFuel(tt.fuelMass)
		if got != tt.fuel {
			t.Errorf("calculateFuelFuel(%d) = %d, want %d", tt.fuelMass, got, tt.fuel)
		}
	}
}
