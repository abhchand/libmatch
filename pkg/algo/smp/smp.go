package smp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

// Run executes the algorithm to solve the "Stable Marriage Problem" (SMP) for
// a set of given preference inputs.
//
// See smp package documentation for more detail
func Run(algoCtx core.AlgorithmContext) (core.MatchResult, error) {
	ptA := algoCtx.TableA
	ptB := algoCtx.TableB

	phase1Proposal(ptA, ptB)

	return buildResult(ptA, ptB), nil
}

// buildResult constructs a Match Result from a Preference Table that has
// been reduced by the algorithm run.
func buildResult(ptA, ptB *core.PreferenceTable) core.MatchResult {
	res := core.MatchResult{}

	res.Mapping = make(map[string]string)

	for _, pt := range []*core.PreferenceTable{ptA, ptB} {
		for name, member := range *pt {
			res.Mapping[name] = member.CurrentProposer().Name()
		}
	}

	return res
}
