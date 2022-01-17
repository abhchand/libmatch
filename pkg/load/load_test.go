package load

import (
	"strings"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromIO(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B", "C", "D"] },
    { "name":"B", "preferences": ["A", "C", "D"] },
    { "name":"C", "preferences": ["A", "B", "D"] },
    { "name":"D", "preferences": ["A", "B", "C"] }
  ]
	`

	got, err := LoadFromIO(strings.NewReader(body))

	wanted := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	assert.Nil(t, err)
	assert.Equal(t, wanted, got)
}

func TestLoadFromDataType(t *testing.T) {
	entries := []core.MatchEntry{
		core.MatchEntry{Name: "A", Preferences: []string{"B", "C", "D"}},
		core.MatchEntry{Name: "B", Preferences: []string{"A", "C", "D"}},
		core.MatchEntry{Name: "C", Preferences: []string{"A", "B", "D"}},
		core.MatchEntry{Name: "D", Preferences: []string{"A", "B", "C"}},
	}

	wanted := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	assert.Equal(t, wanted, LoadFromDataType(&entries))
}

func TestLoadFromDataType_EmptyTable(t *testing.T) {
	entries := []core.MatchEntry{}

	wanted := core.PreferenceTable{}

	assert.Equal(t, wanted, LoadFromDataType(&entries))
}

func TestLoadFromDataType_DuplicateEntries(t *testing.T) {
	entries := []core.MatchEntry{
		core.MatchEntry{Name: "A", Preferences: []string{"B", "C", "D"}},
		core.MatchEntry{Name: "B", Preferences: []string{"A", "C", "D"}},
		core.MatchEntry{Name: "C", Preferences: []string{"A", "B", "D"}},
		core.MatchEntry{Name: "A", Preferences: []string{"B", "C", "D"}},
	}

	wanted := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
	}

	assert.Equal(t, wanted, LoadFromDataType(&entries))
}

func TestLoadFromDataType_CaseSensitive(t *testing.T) {
	entries := []core.MatchEntry{
		core.MatchEntry{Name: "A", Preferences: []string{"B", "C", "a"}},
		core.MatchEntry{Name: "B", Preferences: []string{"A", "C", "a"}},
		core.MatchEntry{Name: "C", Preferences: []string{"A", "B", "a"}},
		core.MatchEntry{Name: "a", Preferences: []string{"A", "B", "C"}},
	}

	wanted := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "a"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "a"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "a"}},
		"a": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	assert.Equal(t, wanted, LoadFromDataType(&entries))
}
