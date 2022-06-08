package srp

import (
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestPhase1Proposal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pt := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
			{Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
			{Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
			{Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
			{Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
			{Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}},
		})

		wanted := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
			{Name: "B", Preferences: []string{"E", "F", "A", "C"}},
			{Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
			{Name: "D", Preferences: []string{"F", "C", "A", "E"}},
			{Name: "E", Preferences: []string{"C", "D", "B", "A"}},
			{Name: "F", Preferences: []string{"A", "B", "D", "C"}},
		})

		wanted["A"].Accept(wanted["F"])
		wanted["B"].Accept(wanted["A"])
		wanted["C"].Accept(wanted["E"])
		wanted["D"].Accept(wanted["C"])
		wanted["E"].Accept(wanted["B"])
		wanted["F"].Accept(wanted["D"])

		isStable := phase1Proposal(&pt)

		assert.True(t, isStable)
		assert.True(t, reflect.DeepEqual(wanted, pt))
	})

	t.Run("no stable solution exists", func(t *testing.T) {
		/*
					 * All other rooommates prefer "D" the least and prefer each other
			     * with equal priority. In this case D's preference list will get
			     * exhausted as no one prefers D to any other match
		*/
		pt := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		isStable := phase1Proposal(&pt)

		assert.False(t, isStable)
	})
}

func TestIsStable(t *testing.T) {
	pt := core.NewPreferenceTable(&[]core.MatchPreference{
		{Name: "A", Preferences: []string{"B", "C", "D"}},
		{Name: "B", Preferences: []string{"C", "A", "D"}},
		{Name: "C", Preferences: []string{"A", "B", "D"}},
		{Name: "D", Preferences: []string{"A", "B", "C"}},
	})

	pt["C"].Reject(pt["A"])
	pt["C"].Reject(pt["B"])
	assert.True(t, isStable(&pt))

	// Rejecting the last available preference makes the table unstable
	pt["C"].Reject(pt["D"])
	assert.False(t, isStable(&pt))
}

func TestSimulateProposal(t *testing.T) {
	t.Run("proposed has no accepted proposal", func(t *testing.T) {
		pt := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		// C proposes to A, who has no other accepted proposal and will accept
		simulateProposal(pt["C"], pt["A"])

		wanted := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		wanted["A"].Accept(wanted["C"])

		assert.True(t, reflect.DeepEqual(wanted, pt))
	})

	t.Run("proposed prefers new proposal to existing one", func(t *testing.T) {
		pt := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		// C proposes to A, then B proposes to A
		// A will prefer the newer proosal (B) and mutually reject the former proposal (C)
		simulateProposal(pt["C"], pt["A"])
		simulateProposal(pt["B"], pt["A"])

		wanted := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		wanted["A"].Accept(wanted["B"])

		assert.True(t, reflect.DeepEqual(wanted, pt))
	})

	t.Run("proposed doesn't prefer new proposal to existing one", func(t *testing.T) {
		pt := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		// C proposes to A, then D proposes to A
		// A will prefer the former proosal (C) and mutually reject the newer proposal (D)
		simulateProposal(pt["C"], pt["A"])
		simulateProposal(pt["D"], pt["A"])

		wanted := core.NewPreferenceTable(&[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"B", "C"}},
		})

		wanted["A"].Accept(wanted["C"])

		assert.True(t, reflect.DeepEqual(wanted, pt))
	})
}
