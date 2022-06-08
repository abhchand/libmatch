package core

import (
	"strings"
)

// PreferenceList models an ordered list of preferences for other members for
// any given Member.
type PreferenceList struct {
	members []*Member
}

// NewPreferenceList returns a new preference list given an array of initial
// ordered members.
func NewPreferenceList(members []*Member) PreferenceList {
	return PreferenceList{members: members}
}

// String returns a human readable representation of this preference list
func (pl PreferenceList) String() string {
	names := make([]string, len(pl.members))

	for i := range pl.members {
		names[i] = pl.members[i].String()
	}

	return strings.Join(names, ", ")
}

// Members returns the raw list of preferred members
func (pl PreferenceList) Members() []*Member {
	return pl.members
}

// Remove removes a specific member from the preference list
func (pl *PreferenceList) Remove(member Member) {
	idx := -1

	// Find index of `member`
	for m := range pl.members {
		if member.Name() == pl.members[m].Name() {
			idx = m
			break
		}
	}

	if idx == -1 {
		return
	}

	// Remove `member`
	newMembers := make([]*Member, len(pl.members)-1)
	copy(newMembers[:idx], pl.members[:idx])
	copy(newMembers[idx:], pl.members[idx+1:])

	pl.members = newMembers
}
