package core

import (
	"fmt"
	"regexp"
	"sort"
)

type PreferenceTable map[string]*Member

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
