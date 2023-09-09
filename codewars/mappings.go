package codewars

import (
	"fmt"
	"strconv"
	"strings"
)

func GoodVsEvil(good, evil string) (string, error) {
	goodForces, err := parseInput(good)
	if err != nil {
		return "", fmt.Errorf("unable to parse input: %w", err)
	}
	evilForces, err := parseInput(evil)
	if err != nil {
		return "", fmt.Errorf("unable to parse input: %w", err)
	}

	goodMapping := []int{1, 2, 3, 3, 4, 10}
	evilMapping := []int{1, 2, 2, 2, 3, 5, 10}

	if calcStrength(goodForces, goodMapping) >
		calcStrength(evilForces, evilMapping) {
		return fmt.Sprintln("Battle Result: Good triumphs over Evil"), nil
	} else if calcStrength(goodForces, goodMapping) <
		calcStrength(evilForces, evilMapping) {
		return fmt.Sprintln("Battle Result: Evil eradicates all trace of Good"), nil
	} else {
		return fmt.Sprintln("Battle Result: No victor on this battle field"), nil
	}
}

func parseInput(good string) ([]int, error) {
	parts := strings.Split(good, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		i, err := strconv.ParseInt(p, 10, 0)
		if err != nil {
			return nil, err
		}
		out = append(out, int(i))
	}
	return out, nil
}

func calcStrength(forces, mapping []int) int {
	var out int
	for idx, v := range forces {
		out += v * mapping[idx]
	}
	return out
}

type Person struct {
	Name string
	Age  uint8
}
