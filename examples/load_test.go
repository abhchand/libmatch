package examples

import (
	"bufio"
	"fmt"
	"os"

	"github.com/abhchand/libmatch"
)

// ExampleLoad loads preferences from a JSON file and uses it as input to the
// Stable Roommates Problem (SRP).
func ExampleLoad() {

	// Read file contents as `io.Reader`
	file, err := os.Open("./preferences.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Load input preferences
	reader := bufio.NewReader(file)
	prefTable, err := libmatch.Load(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call `libmatch`
	result, err := libmatch.SolveSRP(prefTable)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Iterate through the results
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
