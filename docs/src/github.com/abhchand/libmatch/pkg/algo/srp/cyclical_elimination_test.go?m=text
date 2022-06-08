package srp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestPhase3CyclicalElimnation(t *testing.T) {
	wanted := core.NewPreferenceTable(&[]core.MatchPreference{
		{Name: "A", Preferences: []string{"F"}},
		{Name: "B", Preferences: []string{"E"}},
		{Name: "C", Preferences: []string{"D"}},
		{Name: "D", Preferences: []string{"C"}},
		{Name: "E", Preferences: []string{"B"}},
		{Name: "F", Preferences: []string{"A"}},
	})

	testCases := []string{"A", "B", "C", "D", "E", "F", ""}

	for tc := range testCases {
		title := fmt.Sprintf("With seed: %v", testCases[tc])

		t.Run(title, func(t *testing.T) {
			pt := core.NewPreferenceTable(&[]core.MatchPreference{
				{Name: "A", Preferences: []string{"B", "F"}},
				{Name: "B", Preferences: []string{"E", "F", "A"}},
				{Name: "C", Preferences: []string{"D", "E"}},
				{Name: "D", Preferences: []string{"F", "C"}},
				{Name: "E", Preferences: []string{"C", "B"}},
				{Name: "F", Preferences: []string{"A", "B", "D"}},
			})

			phase3CyclicalElimnationWithSeed(&pt, testCases[tc])

			assert.True(t, reflect.DeepEqual(wanted, pt))
		})
	}

	t.Run("table is already complete", func(t *testing.T) {
		pt := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"F"}},
			{Name: "B", Preferences: []string{"E"}},
			{Name: "C", Preferences: []string{"D"}},
			{Name: "D", Preferences: []string{"C"}},
			{Name: "E", Preferences: []string{"B"}},
			{Name: "F", Preferences: []string{"A"}},
		})

		phase3CyclicalElimnation(&pt)

		assert.True(t, reflect.DeepEqual(wanted, pt))
	})
}
