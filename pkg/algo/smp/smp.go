package smp

import (
	"github.com/abhchand/libmatch/pkg/core"
)

func Run(algoCtx core.AlgorithmContext) (core.MatchResult, error) {
	res := core.MatchResult{}

	ptA := algoCtx.PrimaryTable
	ptB := algoCtx.PartnerTable

	phase1Proposal(ptA, ptB)

	res.Mapping = make(map[string]string)
	for name, member := range *ptA {
		res.Mapping[name] = member.CurrentProposer().Name()
	}
	for name, member := range *ptB {
		res.Mapping[name] = member.CurrentProposer().Name()
	}

	return res, nil
}
