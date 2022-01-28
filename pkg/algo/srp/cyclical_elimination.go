package srp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

type cyclePair struct {
	x, y *core.Member
}

func phase3CyclicalElimnation(pt *core.PreferenceTable) core.MatchResult {
	res := phase3CyclicalElimnationWithSeed(pt, "")
	return res
}

func phase3CyclicalElimnationWithSeed(pt *core.PreferenceTable, seed string) core.MatchResult {
	var startingMember *core.Member
	var loopIdx int
	res := core.MatchResult{}

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
			return res
		}

		loopIdx++
	}

	res.Mapping = make(map[string]string)
	for name, member := range *pt {
		res.Mapping[name] = member.PreferenceList().Members()[0].Name()
	}

	return res
}

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

func eliminateCycle(pt *core.PreferenceTable, pairs []cyclePair) {
	for p := range pairs {
		(pairs[p].x).Reject(pairs[p].y)
	}
}
