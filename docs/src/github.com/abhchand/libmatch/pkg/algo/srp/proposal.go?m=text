package srp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

// phase1Proposal implements the 1st phase of the Irving (1985) algorithm to
// solve the "Stable Roommate Problem".
//
// Each unmatched member "proposes" to their top remaining preference and each
// member that receives a proposal can accept or reject the incoming proposal.
//
// See srp package documentation for more detail
func phase1Proposal(pt *core.PreferenceTable) bool {
	for true {
		unmatchedMembers := pt.UnmatchedMembers()

		if len(unmatchedMembers) == 0 {
			break
		}

		if !isStable(pt) {
			return false
		}

		member := unmatchedMembers[0]
		topChoice := member.FirstPreference()
		simulateProposal(member, topChoice)
	}

	// Check for stability once more since final iteration may have left the
	//table unstable
	if !isStable(pt) {
		return false
	}

	return true
}

// isStable evaluates the preference table and determines whether it is
// "stable". A table is stable when all members' preference lists are non-empty.
func isStable(pt *core.PreferenceTable) bool {
	for i := range *pt {
		member := (*pt)[i]

		if len(member.PreferenceList().Members()) == 0 {
			return false
		}
	}

	return true
}

// simulateProposal simulates a proposal between two members
func simulateProposal(proposer, proposed *core.Member) {
	if !proposed.HasAcceptedProposal() {
		// Proposed member does not have a proposal. Blindly accept this one.
		proposed.Accept(proposer)
	} else if proposed.WouldPreferProposalFrom(*proposer) {
		// Proposed member has a proposal, but the new proposal is better. Reject
		// the existing proposal and accept this new one.
		proposed.Reject(proposed.CurrentProposer())
		proposed.Accept(proposer)
	} else {
		// Proposed member has a proposal, but prefers to hold on to it. Reject
		// this new proposal.
		proposed.Reject(proposer)
	}
}
