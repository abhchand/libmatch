package validate

import (
	"errors"
	"fmt"
	"sort"

	"github.com/abhchand/libmatch/pkg/core"
)

type PreferenceTableValidator struct {
	PrimaryTable core.PreferenceTable
	Err          error
}

func (v *PreferenceTableValidator) Validate() error {
	var err error

	err = v.validateSize(v.PrimaryTable)
	if err != nil {
		return err
	}

	memberNames, err := v.validateMembers(v.PrimaryTable)
	if err != nil {
		return err
	}

	err = v.validateSymmetry(v.PrimaryTable, memberNames)
	if err != nil {
		return err
	}

	return nil
}

func (v PreferenceTableValidator) validateSize(table core.PreferenceTable) error {
	numMembers := len(table)

	if numMembers == 0 {
		return errors.New("Table must be non-empty")
	}

	if numMembers%2 != 0 {
		return errors.New("Table must have an even number of members")
	}

	return nil
}

func (v PreferenceTableValidator) validateMembers(table core.PreferenceTable) ([]string, error) {

	memberNames := make([]string, 0, len(table))

	for name := range table {
		if name == "" {
			return memberNames, errors.New("All member names must non-blank")
		}

		memberNames = append(memberNames, name)
	}

	return memberNames, nil

}

/*
 * Check whether table is symmetrical
 *
 * That is, verify each member's preferences contains all the other members.
 */
func (v PreferenceTableValidator) validateSymmetry(table core.PreferenceTable, memberNames []string) error {
	for name := range table {
		// Find index of this member's name
		var idx int
		for i := range memberNames {
			if memberNames[i] == name {
				idx = i
				break
			}
		}

		/*
		 * Remove this member from the member name list
		 * This result should be the expected preference list (names) for this member
		 */
		expected := make([]string, len(memberNames)-1)
		copy(expected[:idx], memberNames[:idx])
		copy(expected[idx:], memberNames[idx+1:])

		// Determine the actual list of preference list (names) for this member
		prefs := table[name].PreferenceList().Members()
		actual := make([]string, len(prefs))
		for i := range prefs {
			actual[i] = prefs[i].Name()
		}

		// Compare
		sort.Strings(actual)
		sort.Strings(expected)

		if !stringSlicesEqual(actual, expected) {
			return errors.New(
				fmt.Sprintf("Preference list for '%v' does not contain all the required members", name))
		}
	}

	return nil
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
