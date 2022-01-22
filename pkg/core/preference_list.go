package core

type PreferenceList struct {
	members []*Member
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
