package validate

import (
	"errors"
	"fmt"

	"github.com/abhchand/libmatch/pkg/core"
)

type DoubleTableValidator struct {
	EntriesList []*[]core.MatchEntry
	Tables      []*core.PreferenceTable
	Err         error
}

func (v DoubleTableValidator) Validate() error {
	var err error

	// This should already be verified upstream
	if len(v.EntriesList) != 2 || len(v.Tables) != 2 {
		return errors.New("Internal error: expected exactly 2 Entries and 2 Tables")
	}

	err = v.validateEntriesUniqueness()
	if err != nil {
		return err
	}

	err = v.validateTableUniqueness()
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

func (v DoubleTableValidator) validateEntriesUniqueness() error {
	caches := make([]map[string]bool, 2)

	for e := range v.EntriesList {
		caches[e] = make(map[string]bool, 0)

		for i := range *v.EntriesList[e] {
			name := (*v.EntriesList[e])[i].Name

			if caches[e][name] {
				msg := fmt.Sprintf("Member names must be unique. Found duplicate entry '%v'", name)
				return errors.New(msg)
			}

			caches[e][name] = true
		}
	}

	return nil
}

func (v DoubleTableValidator) validateTableUniqueness() error {
	for t := range v.Tables {
		table := v.Tables[t]
		otherTable := v.Tables[1-t]

		for name := range *table {
			if (*otherTable)[name] != nil {
				msg := fmt.Sprintf("Tables must have distinct members. '%v' found in both tables", name)
				return errors.New(msg)
			}
		}
	}

	return nil
}

func (v DoubleTableValidator) validateSize() error {
	sizes := make([]int, 2)

	for t := range v.Tables {
		sizes[t] = len(*v.Tables[t])

		if sizes[t] == 0 {
			return errors.New("Table must be non-empty")
		}
	}

	if sizes[0] != sizes[1] {
		errors.New("Tables must be the same size")
	}

	return nil
}

func (v DoubleTableValidator) validateMembers() error {
	for t := range v.Tables {
		if (*v.Tables[t])[""] != nil {
			return errors.New("All member names must non-blank")
		}
	}

	return nil
}

/*
 * Check whether the tables are symmetrical
 *
 * That is, verify each member's preferences contains all members of the other
 * table
 */
func (v DoubleTableValidator) validateSymmetry() error {

	// Build a list of member names of both tables

	memberNames := make([][]string, 2)

	for t := range v.Tables {
		memberNames[t] = make([]string, len(*v.Tables[t]))

		i := 0
		for name := range *v.Tables[t] {
			memberNames[t][i] = name
			i++
		}
	}

	// Verify each member's preference list across both tables

	for t := range v.Tables {
		table := v.Tables[t]

		for name, member := range *table {
			prefs := member.PreferenceList().Members()

			actual := make([]string, len(prefs))
			for p := range prefs {
				/*
				 * The only way a PreferenceList member would be `nil` is if it referenced
				 * a member that does not exist. That is, no `Member` value could be determined
				 * when constructing the PreferenceTable.
				 */
				if prefs[p] == nil {
					return errors.New(
						fmt.Sprintf("Preference list for '%v' contains at least one unknown member", name))
				}
				actual[p] = prefs[p].Name()
			}

			expected := memberNames[1-t]

			if !stringSlicesMatch(actual, expected) {
				return errors.New(
					fmt.Sprintf("Preference list for '%v' does not contain all the required members", name))
			}
		}
	}

	return nil
}
