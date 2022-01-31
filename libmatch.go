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
	var mp *[]MatchEntry
	var err error

	mp, err = load.LoadFromIO(r)
	if err != nil {
		return mp, err
	}

	return mp, err
}

func SolveSRP(prefs *[]MatchEntry) (MatchResult, error) {
	var res MatchResult
	var err error

	// Validate input data
	meValidator := validate.MatchEntryValidator{Entries: prefs}

	if err = meValidator.Validate(); err != nil {
		return res, err
	}

	// Build and validate preference table
	table := core.NewPreferenceTable(prefs)
	ptValidator := validate.PreferenceTableValidator{PrimaryTable: table}

	if err = ptValidator.Validate(); err != nil {
		return res, err
	}

	algoCtx := core.AlgorithmContext{
		PrimaryTable: &table,
	}

	res, err = srp.Run(algoCtx)

	return res, err
}
