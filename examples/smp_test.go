package examples

import (
	"fmt"
	"os"

	"github.com/abhchand/libmatch"
)

// ExampleSolveSMP solves the "Stable Matching Problem" for some sample input
func ExampleSolveSMP() {

	prefTableA := []libmatch.MatchPreference{
		{Name: "A", Preferences: []string{"F", "J", "H", "G", "I"}},
		{Name: "B", Preferences: []string{"F", "J", "H", "G", "I"}},
		{Name: "C", Preferences: []string{"F", "G", "H", "J", "I"}},
		{Name: "D", Preferences: []string{"H", "J", "F", "I", "G"}},
		{Name: "E", Preferences: []string{"H", "F", "G", "I", "J"}},
	}

	prefTableB := []libmatch.MatchPreference{
		{Name: "F", Preferences: []string{"A", "E", "C", "B", "D"}},
		{Name: "G", Preferences: []string{"D", "E", "C", "B", "A"}},
		{Name: "H", Preferences: []string{"A", "B", "C", "D", "E"}},
		{Name: "I", Preferences: []string{"B", "E", "C", "D", "A"}},
		{Name: "J", Preferences: []string{"E", "A", "D", "B", "C"}},
	}

	// Call `libmatch`
	result, err := libmatch.SolveSMP(&prefTableA, &prefTableB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Iterate through the result mapping
	for x, y := range result.Mapping {
		fmt.Printf("%v => %v\n", x, y)
	}

	// Unordered output:
	// A => F
	// B => H
	// C => I
	// D => J
	// E => G
	// F => A
	// G => E
	// H => B
	// I => C
	// J => D
}
