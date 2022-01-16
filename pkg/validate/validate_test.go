package validate

import (
	"fmt"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	table := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	v := Validator{PrimaryTable: table}
	err := v.Validate()

	assert.Nil(t, err)
}

func TestValidate_EmptyTable(t *testing.T) {
	table := core.PreferenceTable{}

	v := Validator{PrimaryTable: table}
	err := v.Validate()

	if assert.NotNil(t, err) {
		assert.Equal(t, "Table must be non-empty", err.Error())
	}
}

func TestValidate_OddNumberOfMembers(t *testing.T) {
	table := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
	}

	v := Validator{PrimaryTable: table}
	err := v.Validate()

	if assert.NotNil(t, err) {
		assert.Equal(t, "Table must have an even number of members", err.Error())
	}
}

func TestValidate_EmptyMember(t *testing.T) {
	table := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"":  core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	v := Validator{PrimaryTable: table}
	err := v.Validate()

	if assert.NotNil(t, err) {
		assert.Equal(t, "All member names must non-blank", err.Error())
	}
}

func TestValidate_AsymmetricalEmptyList(t *testing.T) {
	table := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	v := Validator{PrimaryTable: table}
	err := v.Validate()

	if assert.NotNil(t, err) {
		wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "A")
		assert.Equal(t, wanted, err.Error())
	}
}

func TestValidate_AsymmetricalMismatchedList(t *testing.T) {
	table := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	v := Validator{PrimaryTable: table}
	err := v.Validate()

	if assert.NotNil(t, err) {
		wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "A")
		assert.Equal(t, wanted, err.Error())
	}
}
