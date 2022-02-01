package libmatch

import (
	"io"

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

func SolveSRP(prefs *[]MatchEntry) (MatchResult, error) {
	var res MatchResult
	var err error

	table := core.NewPreferenceTable(prefs)
	validator := validate.SingleTableValidator{Entries: prefs, Table: &table}

	if err = validator.Validate(); err != nil {
		return res, err
	}

	algoCtx := core.AlgorithmContext{
		PrimaryTable: &table,
	}

	res, err = srp.Run(algoCtx)

	return res, err
}
