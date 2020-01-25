package challenge12

import "math"

func CalculateTotalEnergy(moons []*Moon) int {
	totalEnergy := 0.0

	for _, moon := range moons {
		potentialEnergy := math.Abs(float64(moon.position.x)) + math.Abs(float64(moon.position.y)) + math.Abs(float64(moon.position.z))
		kineticEnergy := math.Abs(float64(moon.velocity.x)) + math.Abs(float64(moon.velocity.y)) + math.Abs(float64(moon.velocity.z))

		totalEnergy += potentialEnergy * kineticEnergy
	}

	return int(totalEnergy)
}
