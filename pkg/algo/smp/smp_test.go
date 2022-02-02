package smp

import (
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entriesList := []*[]core.MatchEntry{
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

		tables := core.NewPreferenceTablePair(entriesList[0], entriesList[1])

		algoCtx := core.AlgorithmContext{
			PrimaryTable: &tables[0],
			PartnerTable: &tables[1],
		}

		wanted := core.MatchResult{
			Mapping: map[string]string{
				"A": "K",
				"B": "J",
				"C": "L",
				"D": "I",
				"E": "H",
				"F": "M",
				"K": "A",
				"J": "B",
				"L": "C",
				"I": "D",
				"M": "F",
				"H": "E",
			},
		}

		result, err := Run(algoCtx)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("table order is reversible", func(t *testing.T) {
		entriesList := []*[]core.MatchEntry{
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

		tables := core.NewPreferenceTablePair(entriesList[0], entriesList[1])

		algoCtx := core.AlgorithmContext{
			PrimaryTable: &tables[1],
			PartnerTable: &tables[0],
		}

		wanted := core.MatchResult{
			Mapping: map[string]string{
				"A": "K",
				"B": "J",
				"C": "L",
				"D": "I",
				"E": "M",
				"F": "H",
				"K": "A",
				"J": "B",
				"L": "C",
				"I": "D",
				"M": "E",
				"H": "F",
			},
		}

		result, err := Run(algoCtx)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})
}
