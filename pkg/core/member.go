package core

type Member struct {
	name                 string
	preferenceList       *PreferenceList
	acceptedProposalFrom *Member
}

func NewMember(name string) Member {
	return Member{name: name}
}

func (m Member) Name() string {
	return m.name
}

func (m Member) PreferenceList() *PreferenceList {
	return m.preferenceList
}

func (m *Member) SetPreferenceList(pl *PreferenceList) {
	m.preferenceList = pl
}

func (m Member) CurrentProposer() *Member {
	return m.acceptedProposalFrom
}

func (m Member) CurrentAcceptor() *Member {
	for i := range m.preferenceList.members {
		them := m.preferenceList.members[i]
		theirProposer := them.CurrentProposer()

		if theirProposer != nil && theirProposer.Name() == m.name {
			return m.preferenceList.members[i]
		}
	}

	return nil
}

func (m Member) HasAcceptedProposal() bool {
	return m.acceptedProposalFrom != nil
}

func (m *Member) Accept(member *Member) {
	m.acceptedProposalFrom = member
}

func (m *Member) Reject(member *Member) {

	// Clear "current proposer" if that's who we're rejecting
	if m.CurrentProposer() != nil && m.CurrentProposer().Name() == member.Name() {
		m.acceptedProposalFrom = nil
	}

	// Remove both members from each other's preference lists
	m.preferenceList.Remove(*member)
	member.PreferenceList().Remove(*m)
}

func (m Member) WouldPreferProposalFrom(newProposer Member) bool {
	// If there's no proposal accepted, this member will always prefer
	// a new proposal
	if !m.HasAcceptedProposal() {
		return true
	}

	idxNew := -1
	idxCurrent := -1

	for i := range m.preferenceList.members {
		if m.preferenceList.members[i].Name() == newProposer.Name() {
			idxNew = i
		}
		if m.preferenceList.members[i].Name() == m.CurrentProposer().Name() {
			idxCurrent = i
		}
		if idxNew > -1 && idxCurrent > -1 {
			break
		}
	}

	// A lower index means a higher preference. The new proposal is more
	// attractive if it's index is less than the current.
	return idxNew < idxCurrent
}

func (m Member) FirstPreference() *Member {
	if len(m.preferenceList.members) == 0 {
		return nil
	}

	return m.preferenceList.members[0]
}

func (m Member) SecondPreference() *Member {
	if len(m.preferenceList.members) <= 1 {
		return nil
	}

	return m.preferenceList.members[1]
}

func (m Member) LastPreference() *Member {
	if len(m.preferenceList.members) == 0 {
		return nil
	}

	return m.preferenceList.members[len(m.preferenceList.members)-1]
}
