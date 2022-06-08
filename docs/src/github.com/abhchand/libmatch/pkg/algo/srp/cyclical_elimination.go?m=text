package srp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

type cyclePair struct {
	x, y *core.Member
}

// phase3CyclicalElimnation implements the 3rd phase of the Irving (1985)
// algorithm to solve the "Stable Roommate Problem".
//
// In this last phase we attempt to find any preference cycles and reject them.
//
// See srp package documentation for more detail
func phase3CyclicalElimnation(pt *core.PreferenceTable) {
	phase3CyclicalElimnationWithSeed(pt, "")
}

// phase3CyclicalElimnationWithSeed implements the 3rd phase of the Irving
// (1985) algorithm to solve the "Stable Roommate Problem".
//
// It accepts a seed value with which to being processing members
// deterministically. This is useful for testing.
func phase3CyclicalElimnationWithSeed(pt *core.PreferenceTable, seed string) {
	var startingMember *core.Member
	var loopIdx int

	for !pt.IsComplete() {
		if loopIdx == 0 && seed != "" {
			startingMember = (*pt)[seed]
		} else {
			// Find the first memeber with at least two preferences
			for _, member := range *pt {
				if len(member.PreferenceList().Members()) >= 2 {
					startingMember = member
					break
				}
			}
		}

		pairs := detectCycle(pt, startingMember)
		eliminateCycle(pt, pairs)

		if !pt.IsStable() {
			return
		}

		loopIdx++
	}
}

// detectCycle detects preference cycles in a preference table, given a starting
// member.
func detectCycle(pt *core.PreferenceTable, startingMember *core.Member) []cyclePair {
	pairs := []cyclePair{
		{x: startingMember},
	}

	lastSeenAt := make(map[string]int, 0)
	lastSeenAt[startingMember.Name()] = 1
	currentMemberIdx := 0

	for true {
		currentMember := pairs[currentMemberIdx].x

		newPair := cyclePair{
			x: currentMember.SecondPreference().LastPreference(),
			y: currentMember.SecondPreference(),
		}

		pairs = append(pairs, newPair)

		if idx := lastSeenAt[newPair.x.Name()]; idx > 0 {
			pairs = pairs[idx:]
			break
		}

		lastSeenAt[newPair.x.Name()] = currentMemberIdx + 1
		currentMemberIdx = currentMemberIdx + 1
	}

	return pairs
}

// eliminateCycle removes an identified preference cycle in a preference table
func eliminateCycle(pt *core.PreferenceTable, pairs []cyclePair) {
	for p := range pairs {
		(pairs[p].x).Reject(pairs[p].y)
	}
}
