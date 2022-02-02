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
		proposed.Accept(proposer)
		proposer.Accept(proposed)
	} else if proposed.WouldPreferProposalFrom(*proposer) {
		proposed.CurrentProposer().Reject(proposed)
		proposed.Reject(proposed.CurrentProposer())
		proposed.Accept(proposer)
		proposer.Accept(proposed)
	} else {
		proposed.Reject(proposer)
	}
}
