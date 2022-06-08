package srp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

// phase2Rejection implements the 2nd phase of the Irving (1985) algorithm to
// solve the "Stable Roommate Problem".
//
// Each member that has accepted a proposal will remove those they prefer less
// than their current proposer.
//
// See srp package documentation for more detail
func phase2Rejection(pt *core.PreferenceTable) {
	for _, member := range *pt {
		idx := -1
		prefs := member.PreferenceList().Members()

		// Find the index of the current proposer
		for i := range prefs {
			if prefs[i].Name() == member.CurrentProposer().Name() {
				idx = i
				break
			}
		}

		// Reject all members less preferred than the current proposer
		membersToReject := prefs[(idx + 1):]
		for i := range membersToReject {
			member.Reject(membersToReject[i])
		}
	}
}
