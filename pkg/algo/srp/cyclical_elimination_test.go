package srp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestPhase3CyclicalElimnation(t *testing.T) {
	wanted := core.MatchResult{
		Mapping: map[string]string{
			"A": "F",
			"B": "E",
			"C": "D",
			"D": "C",
			"E": "B",
			"F": "A",
		},
	}

	testCases := []string{"A", "B", "C", "D", "E", "F", ""}

	for tc := range testCases {
		title := fmt.Sprintf("With seed: %v", testCases[tc])

		t.Run(title, func(t *testing.T) {
			pt := core.NewPreferenceTable(&[]core.MatchEntry{
				{Name: "A", Preferences: []string{"B", "F"}},
				{Name: "B", Preferences: []string{"E", "F", "A"}},
				{Name: "C", Preferences: []string{"D", "E"}},
				{Name: "D", Preferences: []string{"F", "C"}},
				{Name: "E", Preferences: []string{"C", "B"}},
				{Name: "F", Preferences: []string{"A", "B", "D"}},
			})

			result := phase3CyclicalElimnationWithSeed(&pt, testCases[tc])

			assert.True(t, reflect.DeepEqual(wanted, result))
		})
	}
}
