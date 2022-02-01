package core

import (
	"fmt"
	"regexp"
	"sort"
)

type PreferenceTable map[string]*Member

func NewPreferenceTable(entries *[]MatchEntry) PreferenceTable {
	e := *entries

	table := make(PreferenceTable, len(*entries))

	// First pass: build a list of members as a lookup table
	for i := range e {
		m := NewMember(e[i].Name)
		table[e[i].Name] = &m
	}

	// Second pass, build preference list for each member
	// that contains references to other members
	for i := range e {
		name := e[i].Name
		m := table[name]
		plMembers := make([]*Member, len(e[i].Preferences))

		for p := range e[i].Preferences {
			prefName := e[i].Preferences[p]
			pref := table[prefName]
			plMembers[p] = pref
		}

		m.preferenceList = &PreferenceList{members: plMembers}
		table[name] = m
	}

	return table
}

func NewPreferenceTablePair(entriesA, entriesB *[]MatchEntry) []PreferenceTable {
	entriesList := []*[]MatchEntry{entriesA, entriesB}

	tables := make([]PreferenceTable, 2)
	tables[0] = make(PreferenceTable, len(*entriesA))
	tables[1] = make(PreferenceTable, len(*entriesB))

	// First pass: build a list of members as a lookup table
	for i := range entriesList {
		entries := entriesList[i]

		for j := range *entries {
			name := (*entries)[j].Name
			m := NewMember(name)
			tables[i][name] = &m
		}
	}

	// Second pass, build preference list for each member
	// that contains references to the other table's members
	for i := range entriesList {
		e := *entriesList[i]

		table := tables[i]
		otherTable := tables[1-i]

		for j := range e {
			name := e[j].Name
			m := table[name]
			plMembers := make([]*Member, len(e[j].Preferences))

			for p := range e[j].Preferences {
				prefName := e[j].Preferences[p]
				pref := otherTable[prefName]
				plMembers[p] = pref
			}

			m.preferenceList = &PreferenceList{members: plMembers}
			tables[i][name] = m
		}
	}

	return tables
}

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

func (pt PreferenceTable) UnmatchedMembers() []*Member {
	var unmatched []*Member

	for m := range pt {
		if pt[m].CurrentAcceptor() == nil {
			unmatched = append(unmatched, pt[m])
		}
	}

	return unmatched
}

func (pt PreferenceTable) IsStable() bool {
	for m := range pt {
		if len(pt[m].PreferenceList().members) == 0 {
			return false
		}
	}

	return true
}

func (pt PreferenceTable) IsComplete() bool {
	for m := range pt {
		if len(pt[m].PreferenceList().members) != 1 {
			return false
		}
	}

	return true
}
