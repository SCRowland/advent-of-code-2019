package main

import "fmt"

func main() {
	start := 356261
	end := 846303

	got := NumberPotentialPasswords(start, end)

	fmt.Printf("Answer == %d\n", got)
}
