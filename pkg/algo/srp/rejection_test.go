package srp

import (
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestPhase2Rejection(t *testing.T) {
	pt := core.NewPreferenceTable(&[]core.MatchEntry{
		{Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
		{Name: "B", Preferences: []string{"E", "F", "A", "C"}},
		{Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
		{Name: "D", Preferences: []string{"F", "C", "A", "E"}},
		{Name: "E", Preferences: []string{"C", "D", "B", "A"}},
		{Name: "F", Preferences: []string{"A", "B", "D", "C"}},
	})

	pt["A"].Accept(pt["F"])
	pt["B"].Accept(pt["A"])
	pt["C"].Accept(pt["E"])
	pt["D"].Accept(pt["C"])
	pt["E"].Accept(pt["B"])
	pt["F"].Accept(pt["D"])

	wanted := core.NewPreferenceTable(&[]core.MatchEntry{
		{Name: "A", Preferences: []string{"B", "F"}},
		{Name: "B", Preferences: []string{"E", "F", "A"}},
		{Name: "C", Preferences: []string{"D", "E"}},
		{Name: "D", Preferences: []string{"F", "C"}},
		{Name: "E", Preferences: []string{"C", "B"}},
		{Name: "F", Preferences: []string{"A", "B", "D"}},
	})

	wanted["A"].Accept(wanted["F"])
	wanted["B"].Accept(wanted["A"])
	wanted["C"].Accept(wanted["E"])
	wanted["D"].Accept(wanted["C"])
	wanted["E"].Accept(wanted["B"])
	wanted["F"].Accept(wanted["D"])

	phase2Rejection(&pt)

	assert.True(t, reflect.DeepEqual(wanted, pt))
}
