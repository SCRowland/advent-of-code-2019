package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Running Challenge Three\n")

	instructions := []string{}

	for scanner.Scan() {
		instruction := scanner.Text()
		instructions = append(instructions, instruction)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	wireOne := instructions[0]
	wireTwo := instructions[1]

	distance := NearestJunction(wireOne, wireTwo)
	fmt.Printf("Distance %d\n", distance)
}
