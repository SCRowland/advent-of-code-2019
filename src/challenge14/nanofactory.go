package challenge14

import (
	"math"
	"strconv"
	"strings"
)

// Component is a single reagent and amount
type Component struct {
	amount  int64
	element string
}

// Reaction is a reaction with a list of components to generate an output
type Reaction struct {
	output *Component
	inputs []*Component
}

// ReactionMap keeps a map of component names to reactions
type ReactionMap map[string]Reaction

func parseComponent(compStr string) *Component {
	compStr = strings.TrimSpace(compStr)

	tokens := strings.Split(compStr, " ")
	i, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		panic(err)
	}

	component := &Component{
		amount:  int64(i),
		element: tokens[1],
	}

	return component
}

func parseReactionMap(input string) ReactionMap {
	rm := make(ReactionMap)

	for _, line := range strings.Split(input, "\n") {
		equationLR := strings.Split(line, "=>")

		r := Reaction{
			output: parseComponent(equationLR[1]),
		}

		for _, elem := range strings.Split(equationLR[0], ",") {
			r.inputs = append(r.inputs, parseComponent(elem))
		}

		rm[r.output.element] = r
	}

	return rm
}

func calcMinOre(rm ReactionMap, required *Component) int64 {
	needed := map[string]int64{
		required.element: required.amount,
	}

	for true {
		processedAnElement := false

		for elem, amountNeeded := range needed {
			if amountNeeded > 0 && elem != "ORE" {
				processedAnElement = true

				reaction := rm[elem]
				unitsWeCanMake := reaction.output.amount

				var multiplier int64 = 1

				if amountNeeded <= unitsWeCanMake {
					// you can't partially run a reaction
					multiplier = 1
				}

				if amountNeeded > unitsWeCanMake {
					// e.g. we need 7, and we can make 2, so we need to make 4 lots (8)
					fractionalMultiplier := float64(amountNeeded / unitsWeCanMake)
					multiplier = int64(math.Ceil(fractionalMultiplier))
				}

				needed[elem] -= unitsWeCanMake * multiplier

				for _, comp := range reaction.inputs {
					needed[comp.element] += comp.amount * multiplier
				}
			}
		}

		if !processedAnElement {
			break
		}
	}

	return needed["ORE"]
}

func maxFuel(rm ReactionMap, oreAmount int64) int64 {
	// binary search the likely answer range

	var high int64 = oreAmount
	var low int64 = 1

	for true {

		midPoint := ((high - low) / 2) + low

		maxOre := calcMinOre(rm, &Component{midPoint, "FUEL"})
		if high == low {
			return midPoint
		}
		if maxOre > oreAmount {
			high = midPoint - 1
		}
		if maxOre < oreAmount {
			low = midPoint + 1
		}

	}

	return 0
}
