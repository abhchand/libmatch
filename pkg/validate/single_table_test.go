package validate

import (
	"fmt"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestValidate__SingleTable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("duplicate member name", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "A", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			msg := fmt.Sprintf("Member names must be unique. Found duplicate entry '%v'", "A")
			assert.Equal(t, msg, err.Error())
		}
	})

	t.Run("empty table", func(t *testing.T) {
		entries := []core.MatchEntry{}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			assert.Equal(t, "Table must be non-empty", err.Error())
		}
	})

	t.Run("odd number of members", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			assert.Equal(t, "Table must have an even number of members", err.Error())
		}
	})

	t.Run("empty member", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			assert.Equal(t, "All member names must non-blank", err.Error())
		}
	})

	t.Run("member names are case sensitive", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "a"}},
			{Name: "B", Preferences: []string{"A", "C", "a"}},
			{Name: "C", Preferences: []string{"A", "B", "a"}},
			{Name: "a", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("asymmetrical empty list", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "B")
			assert.Equal(t, wanted, err.Error())
		}
	})

	t.Run("asymmetrical mismatched list", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "B")
			assert.Equal(t, wanted, err.Error())
		}
	})

	t.Run("asymmetrical unknown member", func(t *testing.T) {
		entries := []core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "X"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		table := core.NewPreferenceTable(&entries)

		v := SingleTableValidator{Entries: &entries, Table: &table}
		err := v.Validate()

		if assert.NotNil(t, err) {
			wanted := fmt.Sprintf("Preference list for '%v' contains at least one unknown member", "B")
			assert.Equal(t, wanted, err.Error())
		}
	})
}
