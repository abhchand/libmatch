package core

import (
	"fmt"
	"regexp"
	"sort"
)

// PreferenceTable models a preference table for a set of members. In theory
// it maps a member name to its preferene list. In reality, the Member data
// model already contains a member's name and preference list. So this data
// model is a glorified lookup table that maps a member's string name to its
// Member struct object.
//
// This could have been modeled without a dedicated struct, but using a struct
// also provides the convenience of defining additional methods that apply to
// the table itself.
type PreferenceTable map[string]*Member

// NewPreferenceTable creates a new preference table given a list of match
// preferences (likely loaded from JSON data). Each member will have a
// preference list of all other members in the same set.
func NewPreferenceTable(prefs *[]MatchPreference) PreferenceTable {
	p := *prefs

	table := make(PreferenceTable, len(*prefs))

	// First pass: build a list of members as a lookup table
	for i := range p {
		m := NewMember(p[i].Name)
		table[p[i].Name] = &m
	}

	// Second pass, build preference list for each member
	// that contains references to other members
	for i := range p {
		name := p[i].Name
		m := table[name]
		plMembers := make([]*Member, len(p[i].Preferences))

		for j := range p[i].Preferences {
			prefName := p[i].Preferences[j]
			pref := table[prefName]
			plMembers[j] = pref
		}

		m.preferenceList = &PreferenceList{members: plMembers}
		table[name] = m
	}

	return table
}

// NewPreferenceTablePair creates a pair of preference tables where each
// member has a preference list of members in the *other* set.
func NewPreferenceTablePair(prefsA, prefsB *[]MatchPreference) []PreferenceTable {
	prefsSet := []*[]MatchPreference{prefsA, prefsB}

	tables := make([]PreferenceTable, 2)
	tables[0] = make(PreferenceTable, len(*prefsA))
	tables[1] = make(PreferenceTable, len(*prefsB))

	// First pass: build a list of members as a lookup table
	for i := range prefsSet {
		prefs := prefsSet[i]

		for j := range *prefs {
			name := (*prefs)[j].Name
			m := NewMember(name)
			tables[i][name] = &m
		}
	}

	// Second pass, build preference list for each member
	// that contains references to the other table's members
	for i := range prefsSet {
		p := *prefsSet[i]

		table := tables[i]
		otherTable := tables[1-i]

		for j := range p {
			name := p[j].Name
			m := table[name]
			plMembers := make([]*Member, len(p[j].Preferences))

			for k := range p[j].Preferences {
				prefName := p[j].Preferences[k]
				pref := otherTable[prefName]
				plMembers[k] = pref
			}

			m.preferenceList = &PreferenceList{members: plMembers}
			tables[i][name] = m
		}
	}

	return tables
}

// String returns a human readable representation of this preference table
func (pt PreferenceTable) String() string {
	var str string

	// Sort map keys so we can iterate over the map below deterministically
	keys := make([]string, 0, len(pt))
	for k := range pt {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for k := range keys {
		member := pt[keys[k]]
		preferenceList := member.PreferenceList().String()

		if member.CurrentProposer() != nil {
			currentProposer := member.CurrentProposer().String()

			pattern := regexp.MustCompile(currentProposer)
			newPattern := fmt.Sprintf("%v+", currentProposer)

			preferenceList = pattern.ReplaceAllString(preferenceList, newPattern)
		}

		str = str + fmt.Sprintf("%v\t=>\t%v\n", member, preferenceList)
	}

	return str
}

// UnmatchedMembers returns a list of all members in this table who are still
// unmatched.
func (pt PreferenceTable) UnmatchedMembers() []*Member {
	var unmatched []*Member

	for m := range pt {
		if pt[m].CurrentAcceptor() == nil {
			unmatched = append(unmatched, pt[m])
		}
	}

	return unmatched
}

// IsStable indicates whether this table is considered mathematically stable.
// That is, no member should have an empty preference list.
func (pt PreferenceTable) IsStable() bool {
	for m := range pt {
		if len(pt[m].PreferenceList().members) == 0 {
			return false
		}
	}

	return true
}

// IsComplete indicates whether this table is considered complete. That is,
// every member should have exactly 1 member remaining in its preference list.
func (pt PreferenceTable) IsComplete() bool {
	for m := range pt {
		if len(pt[m].PreferenceList().members) != 1 {
			return false
		}
	}

	return true
}
