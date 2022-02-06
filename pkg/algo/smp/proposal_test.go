package smp

import (
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestPhase1Proposal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		actualEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
				{Name: "B", Preferences: []string{"L", "J", "K", "M", "I", "H"}},
				{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
				{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
				{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
				{Name: "F", Preferences: []string{"J", "K", "L", "M", "H", "I"}},
			},
			{
				{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
				{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
				{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
				{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
				{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
				{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
			},
		}

		wantedEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"K", "I", "M"}},
				{Name: "B", Preferences: []string{"J", "K", "M", "I", "H"}},
				{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
				{Name: "D", Preferences: []string{"I", "M", "H"}},
				{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
				{Name: "F", Preferences: []string{"M", "H", "I"}},
			},
			{
				{Name: "H", Preferences: []string{"F", "E", "C", "D", "B"}},
				{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
				{Name: "J", Preferences: []string{"B", "E", "C"}},
				{Name: "K", Preferences: []string{"A", "E", "C", "B"}},
				{Name: "L", Preferences: []string{"C", "E"}},
				{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
			},
		}

		actualTables := core.NewPreferenceTablePair(actualEntries[0], actualEntries[1])
		wantedTables := core.NewPreferenceTablePair(wantedEntries[0], wantedEntries[1])

		wantedTables[0]["A"].AcceptMutually(wantedTables[1]["K"])
		wantedTables[0]["B"].AcceptMutually(wantedTables[1]["J"])
		wantedTables[0]["C"].AcceptMutually(wantedTables[1]["L"])
		wantedTables[0]["D"].AcceptMutually(wantedTables[1]["I"])
		wantedTables[0]["E"].AcceptMutually(wantedTables[1]["H"])
		wantedTables[0]["F"].AcceptMutually(wantedTables[1]["M"])

		phase1Proposal(&actualTables[0], &actualTables[1])
		assert.True(t, reflect.DeepEqual(wantedTables[0], actualTables[0]))
		assert.True(t, reflect.DeepEqual(wantedTables[1], actualTables[1]))
	})

	t.Run("table order is reversible", func(t *testing.T) {
		actualEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
				{Name: "B", Preferences: []string{"L", "J", "K", "M", "I", "H"}},
				{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
				{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
				{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
				{Name: "F", Preferences: []string{"J", "K", "L", "M", "H", "I"}},
			},
			{
				{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
				{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
				{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
				{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
				{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
				{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
			},
		}

		wantedEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
				{Name: "B", Preferences: []string{"L", "J", "K", "H"}},
				{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
				{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
				{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
				{Name: "F", Preferences: []string{"J", "K", "L", "M", "H", "I"}},
			},
			{
				{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
				{Name: "I", Preferences: []string{"D", "A", "E", "C", "F"}},
				{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
				{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
				{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
				{Name: "M", Preferences: []string{"E", "D", "F", "C", "A"}},
			},
		}

		actualTables := core.NewPreferenceTablePair(actualEntries[0], actualEntries[1])
		wantedTables := core.NewPreferenceTablePair(wantedEntries[0], wantedEntries[1])

		wantedTables[0]["A"].AcceptMutually(wantedTables[1]["K"])
		wantedTables[0]["B"].AcceptMutually(wantedTables[1]["J"])
		wantedTables[0]["C"].AcceptMutually(wantedTables[1]["L"])
		wantedTables[0]["D"].AcceptMutually(wantedTables[1]["I"])
		wantedTables[0]["E"].AcceptMutually(wantedTables[1]["M"])
		wantedTables[0]["F"].AcceptMutually(wantedTables[1]["H"])

		phase1Proposal(&actualTables[1], &actualTables[0])
		assert.True(t, reflect.DeepEqual(wantedTables[1], actualTables[1]))
		assert.True(t, reflect.DeepEqual(wantedTables[0], actualTables[0]))
	})
}

func TestSimulateProposal(t *testing.T) {
	t.Run("proposed has no accepted proposal", func(t *testing.T) {
		actualEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"H", "I", "J"}},
				{Name: "B", Preferences: []string{"H", "I", "J"}},
				{Name: "C", Preferences: []string{"H", "I", "J"}},
			},
			{
				{Name: "H", Preferences: []string{"A", "B", "C"}},
				{Name: "I", Preferences: []string{"A", "B", "C"}},
				{Name: "J", Preferences: []string{"A", "B", "C"}},
			},
		}

		actualTables := core.NewPreferenceTablePair(actualEntries[0], actualEntries[1])

		// C proposes to I, who has no other accepted proposal and will accept
		simulateProposal(actualTables[0]["C"], actualTables[1]["I"])

		wantedEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"H", "I", "J"}},
				{Name: "B", Preferences: []string{"H", "I", "J"}},
				{Name: "C", Preferences: []string{"H", "I", "J"}},
			},
			{
				{Name: "H", Preferences: []string{"A", "B", "C"}},
				{Name: "I", Preferences: []string{"A", "B", "C"}},
				{Name: "J", Preferences: []string{"A", "B", "C"}},
			},
		}

		wantedTables := core.NewPreferenceTablePair(wantedEntries[0], wantedEntries[1])

		wantedTables[0]["C"].AcceptMutually(wantedTables[1]["I"])

		assert.True(t, reflect.DeepEqual(wantedTables, actualTables))
	})

	t.Run("proposed prefers new proposal to existing one", func(t *testing.T) {
		actualEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"H", "I", "J"}},
				{Name: "B", Preferences: []string{"H", "I", "J"}},
				{Name: "C", Preferences: []string{"H", "I", "J"}},
			},
			{
				{Name: "H", Preferences: []string{"A", "B", "C"}},
				{Name: "I", Preferences: []string{"A", "B", "C"}},
				{Name: "J", Preferences: []string{"A", "B", "C"}},
			},
		}

		actualTables := core.NewPreferenceTablePair(actualEntries[0], actualEntries[1])

		// B proposes to H, then A proposes to H
		// H will prefer the newer proosal (A) and mutually reject the former proposal (B)
		simulateProposal(actualTables[0]["B"], actualTables[1]["H"])
		simulateProposal(actualTables[0]["A"], actualTables[1]["H"])

		wantedEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"H", "I", "J"}},
				{Name: "B", Preferences: []string{"I", "J"}},
				{Name: "C", Preferences: []string{"H", "I", "J"}},
			},
			{
				{Name: "H", Preferences: []string{"A", "C"}},
				{Name: "I", Preferences: []string{"A", "B", "C"}},
				{Name: "J", Preferences: []string{"A", "B", "C"}},
			},
		}

		wantedTables := core.NewPreferenceTablePair(wantedEntries[0], wantedEntries[1])

		wantedTables[1]["H"].AcceptMutually(wantedTables[0]["A"])

		assert.True(t, reflect.DeepEqual(wantedTables[1], actualTables[1]))
	})

	t.Run("proposed doesn't prefer new proposal to existing one", func(t *testing.T) {
		actualEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"H", "I", "J"}},
				{Name: "B", Preferences: []string{"H", "I", "J"}},
				{Name: "C", Preferences: []string{"H", "I", "J"}},
			},
			{
				{Name: "H", Preferences: []string{"A", "B", "C"}},
				{Name: "I", Preferences: []string{"A", "B", "C"}},
				{Name: "J", Preferences: []string{"A", "B", "C"}},
			},
		}

		actualTables := core.NewPreferenceTablePair(actualEntries[0], actualEntries[1])

		// A proposes to H, then B proposes to H
		// H will prefer the formaer proosal (A) and mutually reject the newer proposal (B)
		simulateProposal(actualTables[0]["A"], actualTables[1]["H"])
		simulateProposal(actualTables[0]["B"], actualTables[1]["H"])

		wantedEntries := []*[]core.MatchEntry{
			{
				{Name: "A", Preferences: []string{"H", "I", "J"}},
				{Name: "B", Preferences: []string{"I", "J"}},
				{Name: "C", Preferences: []string{"H", "I", "J"}},
			},
			{
				{Name: "H", Preferences: []string{"A", "C"}},
				{Name: "I", Preferences: []string{"A", "B", "C"}},
				{Name: "J", Preferences: []string{"A", "B", "C"}},
			},
		}

		wantedTables := core.NewPreferenceTablePair(wantedEntries[0], wantedEntries[1])

		wantedTables[1]["H"].AcceptMutually(wantedTables[0]["A"])

		assert.True(t, reflect.DeepEqual(wantedTables[1], actualTables[1]))
	})
}
