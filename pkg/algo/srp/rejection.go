package srp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

func phase2Rejection(pt *core.PreferenceTable) {
	for _, member := range *pt {
		idx := -1
		prefs := member.PreferenceList().Members()

		for i := range prefs {
			if prefs[i].Name() == member.CurrentProposer().Name() {
				idx = i
				break
			}
		}

		membersToReject := prefs[(idx + 1):]
		for i := range membersToReject {
			member.Reject(membersToReject[i])
		}
	}
}
