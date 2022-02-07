package examples

/*
 * Example of solving the *Stable Roommates Problem* for some sample input
 */

import (
	"fmt"
	"os"

	"github.com/abhchand/libmatch"
)

func ExampleSolveSRP() {
	prefTable := []libmatch.MatchPreference{
		{Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
		{Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
		{Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
		{Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
		{Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
		{Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}},
	}

	// Call `libmatch`
	result, err := libmatch.SolveSRP(&prefTable)
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
	// B => E
	// C => D
	// D => C
	// E => B
	// F => A
}

func ExampleSolveSRP__no_stable_solution() {
	prefTable := []libmatch.MatchPreference{
		{Name: "A", Preferences: []string{"B", "E", "C", "F", "D"}},
		{Name: "B", Preferences: []string{"C", "F", "E", "A", "D"}},
		{Name: "C", Preferences: []string{"E", "A", "F", "D", "B"}},
		{Name: "D", Preferences: []string{"B", "A", "C", "F", "E"}},
		{Name: "E", Preferences: []string{"A", "C", "D", "B", "F"}},
		{Name: "F", Preferences: []string{"C", "A", "E", "B", "D"}},
	}

	// Call `libmatch`
	_, err := libmatch.SolveSRP(&prefTable)

	// The above input has no stable matching solution, and we expect it to return an error
	fmt.Println(err)

	// Output:
	// No stable solution exists
}
