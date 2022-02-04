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

	res = phase3CyclicalElimnation(pt)

	return res, nil
}
