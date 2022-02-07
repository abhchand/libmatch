package validate

import (
	"errors"
	"fmt"

	"github.com/abhchand/libmatch/pkg/core"
)

type SingleTableValidator struct {
	Prefs *[]core.MatchPreference
	Table *core.PreferenceTable
	Err   error
}

func (v SingleTableValidator) Validate() error {
	var err error

	err = v.validateUniqueness()
	if err != nil {
		return err
	}

	err = v.validateSize()
	if err != nil {
		return err
	}

	err = v.validateMembers()
	if err != nil {
		return err
	}

	err = v.validateSymmetry()
	if err != nil {
		return err
	}

	return nil
}

func (v SingleTableValidator) validateUniqueness() error {
	cache := make(map[string]bool, 0)

	for i := range *v.Prefs {
		name := (*v.Prefs)[i].Name

		if cache[name] {
			msg := fmt.Sprintf("Member names must be unique. Found duplicate entry '%v'", name)
			return errors.New(msg)
		}

		cache[name] = true
	}

	return nil
}

func (v SingleTableValidator) validateSize() error {
	numMembers := len(*v.Table)

	if numMembers == 0 {
		return errors.New("Table must be non-empty")
	}

	if numMembers%2 != 0 {
		return errors.New("Table must have an even number of members")
	}

	return nil
}

func (v SingleTableValidator) validateMembers() error {
	if (*v.Table)[""] != nil {
		return errors.New("All member names must non-blank")
	}

	return nil
}

/*
 * Check whether table is symmetrical
 *
 * That is, verify each member's preferences contains all the other members.
 */
func (v SingleTableValidator) validateSymmetry() error {

	// Build a list of member names

	memberNames := make([]string, len(*v.Table))
	i := 0
	for name := range *v.Table {
		memberNames[i] = name
		i++
	}

	// Verify each member's preference list

	for name := range *v.Table {
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
		prefs := (*v.Table)[name].PreferenceList().Members()
		actual := make([]string, len(prefs))
		for i := range prefs {
			/*
			 * The only way a PreferenceList member would be `nil` is if it referenced
			 * a member that does not exist. That is, no `Member` value could be determined
			 * when constructing the PreferenceTable.
			 */
			if prefs[i] == nil {
				return errors.New(
					fmt.Sprintf("Preference list for '%v' contains at least one unknown member", name))
			}
			actual[i] = prefs[i].Name()
		}

		// Compare
		if !stringSlicesMatch(actual, expected) {
			return errors.New(
				fmt.Sprintf("Preference list for '%v' does not contain all the required members", name))
		}
	}

	return nil
}
