package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var runningFuelTotal = 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Calculating fuel required")

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("err %s", err)
			continue
		}

		fuel := calculateFuel(i)
		fuelFuel := calculateFuelFuel(fuel)
		runningFuelTotal += (fuel + fuelFuel)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Printf("Total Fuel Required %d\n", runningFuelTotal)
}
