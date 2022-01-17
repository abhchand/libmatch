package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPreferenceTable(t *testing.T) {
	entries := []MatchEntry{
		{Name: "A", Preferences: []string{"B", "C", "D"}},
		{Name: "B", Preferences: []string{"A", "C", "D"}},
		{Name: "C", Preferences: []string{"A", "B", "D"}},
		{Name: "D", Preferences: []string{"A", "B", "C"}},
	}

	wanted := PreferenceTable{
		"A": PreferenceList{Members: []string{"B", "C", "D"}},
		"B": PreferenceList{Members: []string{"A", "C", "D"}},
		"C": PreferenceList{Members: []string{"A", "B", "D"}},
		"D": PreferenceList{Members: []string{"A", "B", "C"}},
	}

	assert.Equal(t, wanted, NewPreferenceTable(&entries))
}

func TestNewPreferenceTable_EmptyTable(t *testing.T) {
	entries := []MatchEntry{}

	wanted := PreferenceTable{}

	assert.Equal(t, wanted, NewPreferenceTable(&entries))
}

func TestNewPreferenceTable_DuplicateEntries(t *testing.T) {
	entries := []MatchEntry{
		{Name: "A", Preferences: []string{"B", "C", "D"}},
		{Name: "B", Preferences: []string{"A", "C", "D"}},
		{Name: "C", Preferences: []string{"A", "B", "D"}},
		{Name: "A", Preferences: []string{"B", "C", "D"}},
	}

	wanted := PreferenceTable{
		"A": PreferenceList{Members: []string{"B", "C", "D"}},
		"B": PreferenceList{Members: []string{"A", "C", "D"}},
		"C": PreferenceList{Members: []string{"A", "B", "D"}},
	}

	assert.Equal(t, wanted, NewPreferenceTable(&entries))
}

func TestNewPreferenceTable_CaseSensitive(t *testing.T) {
	entries := []MatchEntry{
		{Name: "A", Preferences: []string{"B", "C", "a"}},
		{Name: "B", Preferences: []string{"A", "C", "a"}},
		{Name: "C", Preferences: []string{"A", "B", "a"}},
		{Name: "a", Preferences: []string{"A", "B", "C"}},
	}

	wanted := PreferenceTable{
		"A": PreferenceList{Members: []string{"B", "C", "a"}},
		"B": PreferenceList{Members: []string{"A", "C", "a"}},
		"C": PreferenceList{Members: []string{"A", "B", "a"}},
		"a": PreferenceList{Members: []string{"A", "B", "C"}},
	}

	assert.Equal(t, wanted, NewPreferenceTable(&entries))
}
