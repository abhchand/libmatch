package core

// Member models an individual element that is to be matched with another
// member. Each Member will have an ordered preference list of other members.
//
// As part of the Algorithm execution, each Member can hold a conditional
// "proposal" from another Member. Over the course of the algorithm run, the
// Member may keep this proposal or reject this proposal for a more preferred
// proposal.
type Member struct {
	name                 string
	preferenceList       *PreferenceList
	acceptedProposalFrom *Member
}

// NewMember builds a new member from a unique name
func NewMember(name string) Member {
	return Member{name: name}
}

// String returns a human readable representation of this member
func (m Member) String() string {
	return "'" + m.name + "'"
}

// Name returns the name of this member
func (m Member) Name() string {
	return m.name
}

// PreferenceList returns the current list of other members in order of
// preference. It may change over time as the algorithm runs and eliminates
// certain elements of the list.
func (m Member) PreferenceList() *PreferenceList {
	return m.preferenceList
}

// SetPreferenceList sets the preference list for this member
func (m *Member) SetPreferenceList(pl *PreferenceList) {
	m.preferenceList = pl
}

// CurrentProposer returns the member who currently holds an accepted proposal
// from this Member.
func (m Member) CurrentProposer() *Member {
	return m.acceptedProposalFrom
}

// CurrentAcceptor returns the Member who has currently accepted a proposal from
// this Member.
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

// HasAcceptedProposal indicates whether any member currently holds an accepted
// proposal from this member.
func (m Member) HasAcceptedProposal() bool {
	return m.acceptedProposalFrom != nil
}

// Accept accepts an incoming proposal from another member
func (m *Member) Accept(member *Member) {
	m.acceptedProposalFrom = member
}

// AcceptMutually marks both this member and another specified member as holding
// accepted proposals from each other.
func (m *Member) AcceptMutually(member *Member) {
	m.Accept(member)
	member.Accept(m)
}

// Reject removes the specified member and this member from each others'
// preference lists, and clear's this member's current registered proposer if
// applicable.
//
// Even though this action is two-way (mutual), it only clears this member's
// current registered proposer. Therefore we view this action as "one way", and
// there is a separate `RejectMutually` method that clears both members'
// current registered proposers.
//
// This is needed because from the perspective of this member, it is possible
// to reject another member but still have the other member hold an accepted
// proposal from this member. The algorithm will eventually eliminate this
// pair since at least one of the members has rejected the other, but we don't
// want to change the algorithm internal state pre-maturely.
func (m *Member) Reject(member *Member) {

	// Clear "current proposer" if that's who we're rejecting
	if m.CurrentProposer() != nil && m.CurrentProposer().Name() == member.Name() {
		m.acceptedProposalFrom = nil
	}

	// Remove both members from each other's preference lists
	m.preferenceList.Remove(*member)
	member.PreferenceList().Remove(*m)
}

// RejectMutually marks both this member and another specified member as having
// rejected each other. They will both be removed from each others' preference
// lists and both will remove the other as the currently registered proposer
// (if needed).
func (m *Member) RejectMutually(member *Member) {
	m.Reject(member)
	member.Reject(m)
}

// WouldPreferProposalFrom indicates whether this member would prefer a new
// proposal from the specified member over a proposal it already holds.
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

// FirstPreference returns the first preferred member on this member's
// preference list.
func (m Member) FirstPreference() *Member {
	if len(m.preferenceList.members) == 0 {
		return nil
	}

	return m.preferenceList.members[0]
}

// SecondPreference returns the second preferred member on this member's
// preference list.
func (m Member) SecondPreference() *Member {
	if len(m.preferenceList.members) <= 1 {
		return nil
	}

	return m.preferenceList.members[1]
}

// LastPreference returns the lowest preferred member on this member's
// preference list.
func (m Member) LastPreference() *Member {
	if len(m.preferenceList.members) == 0 {
		return nil
	}

	return m.preferenceList.members[len(m.preferenceList.members)-1]
}
