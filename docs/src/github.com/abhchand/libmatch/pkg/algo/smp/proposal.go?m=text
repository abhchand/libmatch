package smp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

// phase1Proposal implements the 1st (and only) phase of the Gale-Shapley (1962)
// algorithm to solve the "Stable Marriage Problem".
//
// Each unmatched member "proposes" to their top remaining preference and each
// member that receives a proposal can accept or reject the incoming proposal.
//
// See smp package documentation for more detail
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

// simulateProposal simulates a proposal between two members
func simulateProposal(proposer, proposed *core.Member) {
	if !proposed.HasAcceptedProposal() {
		// Proposed member does not have a proposal. Blindly accept this one.
		proposed.AcceptMutually(proposer)
	} else if proposed.WouldPreferProposalFrom(*proposer) {
		// Proposed member has a proposal, but the new proposal is better. Reject
		// the existing proposal and accept this new one.
		proposed.RejectMutually(proposed.CurrentProposer())
		proposed.AcceptMutually(proposer)
	} else {
		// Proposed member has a proposal, but prefers to hold on to it. Reject
		// this new proposal.
		proposed.RejectMutually(proposer)
	}
}
