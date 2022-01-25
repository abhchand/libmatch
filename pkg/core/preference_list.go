package core

import (
	"strings"
)

type PreferenceList struct {
	members []*Member
}

func NewPreferenceList(members []*Member) PreferenceList {
	return PreferenceList{members: members}
}

func (pl PreferenceList) String() string {
	names := make([]string, len(pl.members))

	for i := range pl.members {
		names[i] = pl.members[i].String()
	}

	return strings.Join(names, ", ")
}

func (pl PreferenceList) Members() []*Member {
	return pl.members
}

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
