package smp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

func phase1Proposal(ptA, ptB *core.PreferenceTable) {
	for true {
		unmatchedMembers := ptA.UnmatchedMembers()

		if len(unmatchedMembers) == 0 {
			break
		}

		for i := range unmatchedMembers {
			member := unmatchedMembers[i]
			topChoice := member.FirstPreference()
			simulateProposal(member, topChoice)
		}
	}
}

func simulateProposal(proposer, proposed *core.Member) {
	if !proposed.HasAcceptedProposal() {
		proposed.AcceptMutually(proposer)
	} else if proposed.WouldPreferProposalFrom(*proposer) {
		proposed.RejectMutually(proposed.CurrentProposer())
		proposed.AcceptMutually(proposer)
	} else {
		proposed.RejectMutually(proposer)
	}
}
