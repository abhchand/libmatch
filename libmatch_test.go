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

	t.Run("no stable solution", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B","E","C","F","D"}},
			{Name: "B", Preferences: []string{"C","F","E","A","D"}},
			{Name: "C", Preferences: []string{"E","A","F","D","B"}},
			{Name: "D", Preferences: []string{"B","A","C","F","E"}},
			{Name: "E", Preferences: []string{"A","C","D","B","F"}},
			{Name: "F", Preferences: []string{"C","A","E","B","D"}},
		}

		_, err := SolveSRP(&entries)

		assert.Equal(t, "No stable solution exists", err.Error())
	})

	t.Run("is not dependent on order", func(t *testing.T) {
		entries := []core.MatchEntry{
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

func BenchmarkSolveSRP(b *testing.B) {
	entries := []core.MatchEntry{
		{Name:"A", Preferences: []string{"H","J","E","B","D","I","C","G","F"}},
		{Name:"B", Preferences: []string{"E","I","G","D","A","J","C","F","H"}},
		{Name:"C", Preferences: []string{"J","A","B","H","F","I","G","D","E"}},
		{Name:"D", Preferences: []string{"I","C","E","G","B","A","J","F","H"}},
		{Name:"E", Preferences: []string{"F","J","G","B","C","H","A","D","I"}},
		{Name:"F", Preferences: []string{"C","I","D","E","G","H","A","J","B"}},
		{Name:"G", Preferences: []string{"E","H","F","A","J","C","D","B","I"}},
		{Name:"H", Preferences: []string{"F","G","J","B","I","E","C","A","D"}},
		{Name:"I", Preferences: []string{"J","G","B","D","A","C","E","F","H"}},
		{Name:"J", Preferences: []string{"C","F","A","B","I","G","H","D","E"}},
	}

  for i := 0; i < b.N; i++ {
    SolveSRP(&entries)
  }
}
