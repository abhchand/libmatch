package smp

/*
 * Solves the "Stable Marriage Problem" (SMP) for a set of given inputs.
 *
 *
 * See: https://en.wikipedia.org/wiki/Stable_marriage_problem
 *
 * Implements the Gale-Shapley (1962) algorithm, which is guranteed to
 * return a stable matching for any same-sized inputs.
 *
 * Note that:
 *
 * 	1. Multiple stable matchings may exist, and this process will only return
 *     one of them.
 *
 *  3. The algorithm itself prioritizes the preferences of the first preference
 *     table over the second (the group)
 *
 *  2. While the algorithm is deterministic, the `libmatch` implementation below
 *     is not, It is not guranteed to return the same matching everytime.
 *
 * To expand on #3 - This is due to the underlying storage mechanism, which
 * uses a `map` type. Go randomizes the access keys of a `map` and so the
 * algorithm may produces different results on different runs.
 *
 */

import (
	"github.com/abhchand/libmatch/pkg/core"
)

func Run(algoCtx core.AlgorithmContext) (core.MatchResult, error) {
	ptA := algoCtx.TableA
	ptB := algoCtx.TableB

	phase1Proposal(ptA, ptB)

	return buildResult(ptA, ptB), nil
}

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
