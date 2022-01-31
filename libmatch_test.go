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

		wanted := &[]core.MatchEntry{
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

func TestSolveSRP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entries := []core.MatchEntry{
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

		result, err := SolveSRP(&entries)

		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(wanted, result))
	})

	t.Run("validates match entries", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "A", Preferences: []string{"C", "B", "D"}},
		}

		_, err := SolveSRP(&entries)

		assert.Equal(t, "Member names must be unique. Found duplicate entry 'A'", err.Error())
	})

	t.Run("validates preference table", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
		}

		_, err := SolveSRP(&entries)

		assert.Equal(t, "Table must have an even number of members", err.Error())
	})
}
