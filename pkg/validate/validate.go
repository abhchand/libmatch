package validate

import (
	"errors"
	"fmt"
	"sort"

	"github.com/abhchand/libmatch/pkg/core"
)

type Validator struct {
	PrimaryTable core.PreferenceTable
	Err          error
}

func (v *Validator) Validate() error {
	var err error

	err = v.validateSize(v.PrimaryTable)
	if err != nil {
		return err
	}

	members, err := v.validateMembers(v.PrimaryTable)
	if err != nil {
		return err
	}

	err = v.validateSymmetry(v.PrimaryTable, members)
	if err != nil {
		return err
	}

	return nil
}

func (v Validator) validateSize(table core.PreferenceTable) error {
	numMembers := len(table)

	if numMembers == 0 {
		return errors.New("Table must be non-empty")
	}

	if numMembers%2 != 0 {
		return errors.New("Table must have an even number of members")
	}

	return nil
}

func (v Validator) validateMembers(table core.PreferenceTable) ([]string, error) {

	members := make([]string, 0, len(table))

	for m, _ := range table {
		if m == "" {
			return members, errors.New("All member names must non-blank")
		}

		members = append(members, m)
	}

	return members, nil

}

/*
 * Check whether table is symmetrical
 *
 * That is, verify each member's preferences contains all the other members.
 */
func (v Validator) validateSymmetry(table core.PreferenceTable, members []string) error {
	for m, _ := range table {
		// Find index of member
		var idx int
		for i := range members {
			if members[i] == m {
				idx = i
				break
			}
		}

		/*
		 * Remove this member from the member list
		 * This result should be the expected preference list for this member
		 */
		expected := make([]string, len(members)-1)
		copy(expected[:idx], members[:idx])
		copy(expected[idx:], members[idx+1:])

		actual := make([]string, len(table[m].Members))
		copy(actual, table[m].Members)

		// Compare
		sort.Strings(actual)
		sort.Strings(expected)

		if !stringSlicesEqual(actual, expected) {
			return errors.New(
				fmt.Sprintf("Preference list for '%v' does not contain all the required members", m))
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
