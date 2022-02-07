package libmatch

import (
	"reflect"
	"strings"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		body := `
	  [
	    { "name":"A", "preferences": ["B", "C", "D"] },
	    { "name":"B", "preferences": ["A", "C", "D"] },
	    { "name":"C", "preferences": ["A", "B", "D"] },
	    { "name":"D", "preferences": ["A", "B", "C"] }
	  ]
	  `

		result, err := Load(strings.NewReader(body))

		wanted := &[]core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("error", func(t *testing.T) {
		// Note missing `:` on final row
		body := `
	  [
	    { "name":"A", "preferences": ["B", "C", "D"] },
	    { "name":"B", "preferences": ["A", "C", "D"] },
	    { "name":"C", "preferences": ["A", "B", "D"] },
	    { "name":"D", "preferences" ["A", "B", "C"] }
	  ]
		`

		_, err := Load(strings.NewReader(body))

		assert.Equal(t, "invalid character '[' after object key", err.Error())
	})
}

func TestSolveSMP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		prefsA := []core.MatchPreference{
			{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
			{Name: "B", Preferences: []string{"L", "J", "K", "M", "I", "H"}},
			{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
			{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
			{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
			{Name: "F", Preferences: []string{"J", "K", "L", "M", "H", "I"}},
		}

		prefsB := []core.MatchPreference{
			{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
			{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
			{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
			{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
			{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
			{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
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

		result, err := SolveSMP(&prefsA, &prefsB)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("table order is reversible", func(t *testing.T) {
		prefsA := []core.MatchPreference{
			{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
			{Name: "B", Preferences: []string{"L", "J", "K", "M", "I", "H"}},
			{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
			{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
			{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
			{Name: "F", Preferences: []string{"J", "K", "L", "M", "H", "I"}},
		}

		prefsB := []core.MatchPreference{
			{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
			{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
			{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
			{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
			{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
			{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
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

		result, err := SolveSMP(&prefsB, &prefsA)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("validates match prefs", func(t *testing.T) {
		prefsA := []core.MatchPreference{
			{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
			{Name: "B", Preferences: []string{"L", "J", "K", "M", "I", "H"}},
			{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
			{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
			{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
			{Name: "F", Preferences: []string{"J", "K", "L", "M", "H", "I"}},
			{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
		}

		prefsB := []core.MatchPreference{
			{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
			{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
			{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
			{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
			{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
			{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
		}

		_, err := SolveSMP(&prefsA, &prefsB)

		assert.Equal(t, "Member names must be unique. Found duplicate entry 'A'", err.Error())
	})

	t.Run("validates preference table", func(t *testing.T) {
		prefsA := []core.MatchPreference{
			{Name: "A", Preferences: []string{"L", "J", "H", "K", "I", "M"}},
			{Name: "B", Preferences: []string{"L", "J", "K", "M", "I", "H"}},
			{Name: "C", Preferences: []string{"L", "J", "M", "I", "K", "H"}},
			{Name: "D", Preferences: []string{"L", "K", "J", "I", "M", "H"}},
			{Name: "E", Preferences: []string{"H", "I", "L", "K", "M", "J"}},
		}

		prefsB := []core.MatchPreference{
			{Name: "H", Preferences: []string{"F", "E", "C", "A", "D", "B"}},
			{Name: "I", Preferences: []string{"B", "D", "A", "E", "C", "F"}},
			{Name: "J", Preferences: []string{"B", "A", "F", "E", "D", "C"}},
			{Name: "K", Preferences: []string{"A", "E", "C", "F", "D", "B"}},
			{Name: "L", Preferences: []string{"C", "F", "E", "B", "D", "A"}},
			{Name: "M", Preferences: []string{"B", "E", "D", "F", "C", "A"}},
		}

		_, err := SolveSMP(&prefsA, &prefsB)

		assert.Equal(t, "Tables must be the same size", err.Error())
	})
}

func TestSolveSRP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		prefs := []core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		wanted := core.MatchResult{
			Mapping: map[string]string{
				"A": "B",
				"B": "A",
				"C": "D",
				"D": "C",
			},
		}

		result, err := SolveSRP(&prefs)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("no stable solution", func(t *testing.T) {
		prefs := []core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "E", "C", "F", "D"}},
			{Name: "B", Preferences: []string{"C", "F", "E", "A", "D"}},
			{Name: "C", Preferences: []string{"E", "A", "F", "D", "B"}},
			{Name: "D", Preferences: []string{"B", "A", "C", "F", "E"}},
			{Name: "E", Preferences: []string{"A", "C", "D", "B", "F"}},
			{Name: "F", Preferences: []string{"C", "A", "E", "B", "D"}},
		}

		_, err := SolveSRP(&prefs)

		assert.Equal(t, "No stable solution exists", err.Error())
	})

	t.Run("is not dependent on order", func(t *testing.T) {
		prefs := []core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
		}

		wanted := core.MatchResult{
			Mapping: map[string]string{
				"A": "B",
				"B": "A",
				"C": "D",
				"D": "C",
			},
		}

		result, err := SolveSRP(&prefs)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("validates match prefs", func(t *testing.T) {
		prefs := []core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "A", Preferences: []string{"C", "B", "D"}},
		}

		_, err := SolveSRP(&prefs)

		assert.Equal(t, "Member names must be unique. Found duplicate entry 'A'", err.Error())
	})

	t.Run("validates preference table", func(t *testing.T) {
		prefs := []core.MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
		}

		_, err := SolveSRP(&prefs)

		assert.Equal(t, "Table must have an even number of members", err.Error())
	})
}
