package srp

import (
	"reflect"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pt := core.NewPreferenceTable(&[]core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		algoCtx := core.AlgorithmContext{
			PrimaryTable: &pt,
		}

		wanted := core.MatchResult{
			Mapping: map[string]string{
				"A": "B",
				"B": "A",
				"C": "D",
				"D": "C",
			},
		}

		result, err := Run(algoCtx)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("no stable solution exists", func(t *testing.T) {
		/*
					 * All other rooommates prefer "D" the least and prefer each other
			     * with equal priority. In this case D's preference list will get
			     * exhausted as no one prefers D to any other match
		*/
		pt := core.NewPreferenceTable(&[]core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"C", "A", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		})

		algoCtx := core.AlgorithmContext{
			PrimaryTable: &pt,
		}

		_, err := Run(algoCtx)

		assert.Equal(t, "No stable solution exists", err.Error())
	})
}
