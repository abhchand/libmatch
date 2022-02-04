package libmatch

import (
	"io"

	"github.com/abhchand/libmatch/pkg/algo/smp"
	"github.com/abhchand/libmatch/pkg/algo/srp"
	"github.com/abhchand/libmatch/pkg/core"
	"github.com/abhchand/libmatch/pkg/load"
	"github.com/abhchand/libmatch/pkg/validate"
)

type MatchEntry = core.MatchEntry
type MatchResult = core.MatchResult

func Load(r io.Reader) (*[]MatchEntry, error) {
	mp, err := load.LoadFromIO(r)
	if err != nil {
		return mp, err
	}

	return mp, err
}

func SolveSMP(prefsA, prefsB *[]MatchEntry) (MatchResult, error) {
	var res MatchResult
	var err error

	tables := core.NewPreferenceTablePair(prefsA, prefsB)
	validator := validate.DoubleTableValidator{
		EntriesList: []*[]core.MatchEntry{prefsA, prefsB},
		Tables:      []*core.PreferenceTable{&tables[0], &tables[1]},
	}

	if err = validator.Validate(); err != nil {
		return res, err
	}

	algoCtx := core.AlgorithmContext{
		TableA: &tables[0],
		TableB: &tables[1],
	}

	res, err = smp.Run(algoCtx)

	return res, err
}

func SolveSRP(prefs *[]MatchEntry) (MatchResult, error) {
	var res MatchResult
	var err error

	table := core.NewPreferenceTable(prefs)
	validator := validate.SingleTableValidator{Entries: prefs, Table: &table}

	if err = validator.Validate(); err != nil {
		return res, err
	}

	algoCtx := core.AlgorithmContext{
		TableA: &table,
	}

	res, err = srp.Run(algoCtx)

	return res, err
}
