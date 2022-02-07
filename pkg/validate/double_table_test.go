package validate

import (
	"fmt"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestValidate__DoubleTable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("bad number of prefs", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: []*[]core.MatchPreference{prefsSet[0]},
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		assert.Equal(t, "Internal error: expected exactly 2 Prefs and 2 Tables", err.Error())
	})

	t.Run("bad number of tables", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1], &tables[1]},
		}
		err := v.Validate()

		assert.Equal(t, "Internal error: expected exactly 2 Prefs and 2 Tables", err.Error())
	})

	t.Run("duplicate member name within a table", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
				{Name: "A", Preferences: []string{"L", "K", "M"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		msg := fmt.Sprintf("Member names must be unique. Found duplicate entry '%v'", "A")
		assert.Equal(t, msg, err.Error())
	})

	t.Run("duplicate member name across tables", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "B", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		msg := fmt.Sprintf("Tables must have distinct members. '%v' found in both tables", "B")
		assert.Equal(t, msg, err.Error())
	})

	t.Run("empty table", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		assert.Equal(t, "Table must be non-empty", err.Error())
	})

	t.Run("tables are differente sizes", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		assert.Equal(t, "Tables must be the same size", err.Error())
	})

	t.Run("empty member", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		assert.Equal(t, "All member names must non-blank", err.Error())
	})

	t.Run("member names are case sensitive", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "a", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "a", "A"}},
				{Name: "L", Preferences: []string{"A", "a", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "a"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("asymmetrical empty list", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "C", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "B")
		assert.Equal(t, wanted, err.Error())
	})

	t.Run("asymmetrical mismatched list", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "L")
		assert.Equal(t, wanted, err.Error())
	})

	t.Run("asymmetrical unknown member", func(t *testing.T) {
		prefsSet := []*[]core.MatchPreference{
			{
				{Name: "A", Preferences: []string{"K", "L", "M"}},
				{Name: "B", Preferences: []string{"L", "M", "K"}},
				{Name: "C", Preferences: []string{"M", "L", "K"}},
			},
			{
				{Name: "K", Preferences: []string{"B", "C", "A"}},
				{Name: "L", Preferences: []string{"A", "X", "B"}},
				{Name: "M", Preferences: []string{"A", "B", "C"}},
			},
		}

		tables := core.NewPreferenceTablePair(prefsSet[0], prefsSet[1])

		v := DoubleTableValidator{
			PrefsSet: prefsSet,
			Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
		}
		err := v.Validate()

		wanted := fmt.Sprintf("Preference list for '%v' contains at least one unknown member", "L")
		assert.Equal(t, wanted, err.Error())
	})
}
