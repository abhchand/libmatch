package srp

import (
	"errors"
	"github.com/abhchand/libmatch/pkg/core"
)

func Run(algoCtx core.AlgorithmContext) (core.MatchResult, error) {
	var res core.MatchResult
	pt := algoCtx.TableA

	if !phase1Proposal(pt) {
		return res, errors.New("No stable solution exists")
	}

	phase2Rejection(pt)
	phase3CyclicalElimnation(pt)

	res.Mapping = make(map[string]string)
	for name, member := range *pt {
		res.Mapping[name] = member.PreferenceList().Members()[0].Name()
	}

	return res, nil
}
