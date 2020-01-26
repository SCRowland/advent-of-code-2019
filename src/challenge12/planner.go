package challenge12

import "math"

func CalculateTotalEnergy(moons []*Moon) int {
	totalEnergy := 0.0

	for _, moon := range moons {
		potentialEnergy := math.Abs(float64(moon.Position.X)) + math.Abs(float64(moon.Position.Y)) + math.Abs(float64(moon.Position.Z))
		kineticEnergy := math.Abs(float64(moon.Velocity.X)) + math.Abs(float64(moon.Velocity.Y)) + math.Abs(float64(moon.Velocity.Z))

		totalEnergy += potentialEnergy * kineticEnergy
	}

	return int(totalEnergy)
}
