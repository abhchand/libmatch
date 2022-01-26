package srp

import (
	"fmt"

	"github.com/abhchand/libmatch/pkg/core"
)

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

func isStable(pt *core.PreferenceTable) bool {
	for i := range *pt {
		member := (*pt)[i]

		if len(member.PreferenceList().Members()) == 0 {
			return false
		}
	}

	return true
}

func simulateProposal(proposer, proposed *core.Member) {
	fmt.Printf("%v proposes to %v,", proposer, proposed)

	if !proposed.HasAcceptedProposal() {
		fmt.Printf("%v accepts %v\n", proposed, proposer)
		proposed.Accept(proposer)
	} else if proposed.WouldPreferProposalFrom(*proposer) {
		fmt.Printf("%v rejects %v, accepts %v\n", proposed, proposed.CurrentProposer(), proposer)
		proposed.Reject(proposed.CurrentProposer())
		proposed.Accept(proposer)
	} else {
		fmt.Printf("%v rejects %v\n", proposed, proposer)
		proposed.Reject(proposer)
	}
}
