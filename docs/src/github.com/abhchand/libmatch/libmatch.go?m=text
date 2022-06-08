/*
Package libmatch provides an API for solving matching problems.

Each matching algorithm has a shorthand acronym that can be used to invoke
the solver. For example "Stable Marriage Problem" has a shorthand of "SMP" and
to invoke the solver, you can call:

	libmatch.SolveSMP(...)

For a full list of available matching algorithms and their shorthands, see
https://github.com/abhchand/libmatch#readme.

If you have the libmatch command line utility installed, you can also run

	libmatch ls

to see this list.
*/
package libmatch

import (
	"io"

	"github.com/abhchand/libmatch/pkg/algo/smp"
	"github.com/abhchand/libmatch/pkg/algo/srp"
	"github.com/abhchand/libmatch/pkg/core"
	"github.com/abhchand/libmatch/pkg/load"
	"github.com/abhchand/libmatch/pkg/validate"
)

type MatchPreference = core.MatchPreference
type MatchResult = core.MatchResult

// Load reads match preference data from an `io.Reader`.
//
// The expected data is a JSON formatted preference table of the format:
//
// 		[
// 		  { "name":"A", "preferences": ["B", "D", "F", "C", "E"] },
// 		  { "name":"B", "preferences": ["D", "E", "F", "A", "C"] },
// 		  { "name":"C", "preferences": ["D", "E", "F", "A", "B"] },
// 		  { "name":"D", "preferences": ["F", "C", "A", "E", "B"] },
// 		  { "name":"E", "preferences": ["F", "C", "D", "B", "A"] },
// 		  { "name":"F", "preferences": ["A", "B", "D", "C", "E"] },
// 		]
//
// The return value is an array of `MatchPreference` structs containing the
// loaded JSON data
//
// 		*[]libmatch.MatchPreference{
// 		  {Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
// 		  {Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
// 		  {Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
// 		  {Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
// 		  {Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
// 		  {Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}},
//    }
func Load(r io.Reader) (*[]MatchPreference, error) {
	mp, err := load.LoadFromIO(r)
	if err != nil {
		return mp, err
	}

	return mp, err
}

// SolveSMP solves the Stable Marriage Problem for a set of preferences.
//
// See: https://en.wikipedia.org/wiki/Stable_marriage_problem
//
// The algorithm finds a stable matching between two same-sized sets.
// Implements the Gale-Shapley (1962) algorithm. A stable solution is always
// guranteed, but it is non-deterministic and potentially one of many.
//
// Example:
//
// SolveSMP takes a pair of preference tables as inputs. Each preference table
// is an array of match preferences.
//
// 		prefTableA := []libmatch.MatchPreference{
// 			{Name: "A", Preferences: []string{"F", "J", "H", "G", "I"}},
// 			{Name: "B", Preferences: []string{"F", "J", "H", "G", "I"}},
// 			{Name: "C", Preferences: []string{"F", "G", "H", "J", "I"}},
// 			{Name: "D", Preferences: []string{"H", "J", "F", "I", "G"}},
// 			{Name: "E", Preferences: []string{"H", "F", "G", "I", "J"}},
// 		}
//
// 		prefTableB := []libmatch.MatchPreference{
// 			{Name: "F", Preferences: []string{"A", "E", "C", "B", "D"}},
// 			{Name: "G", Preferences: []string{"D", "E", "C", "B", "A"}},
// 			{Name: "H", Preferences: []string{"A", "B", "C", "D", "E"}},
// 			{Name: "I", Preferences: []string{"B", "E", "C", "D", "A"}},
// 			{Name: "J", Preferences: []string{"E", "A", "D", "B", "C"}},
// 		}
//
// On success, the return value will be a MatchResult containing the stable the
// mapping between pairs of members.
//
// 		MatchResult{
// 			Mapping: map[string]string{
// 				"A": "F",
// 				"B": "H",
// 				"C": "I",
// 				"D": "J",
// 				"E": "G",
// 				"F": "A",
// 				"G": "E",
// 				"H": "B",
// 				"I": "C",
// 				"J": "D",
// 			},
// 		}
func SolveSMP(prefsA, prefsB *[]MatchPreference) (MatchResult, error) {
	var res MatchResult
	var err error

	tables := core.NewPreferenceTablePair(prefsA, prefsB)
	validator := validate.DoubleTableValidator{
		PrefsSet: []*[]core.MatchPreference{prefsA, prefsB},
		Tables:   []*core.PreferenceTable{&tables[0], &tables[1]},
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

// SolveSRP solves the Stable Roommates Problem for a set of preferences.
//
// See: https://en.wikipedia.org/wiki/Stable_roommates_problem
//
// The algorithm finds a stable matching within an even-sized set. A stable
// solution is not guranteed, but is always deterministic if exists.
// Implements Irving's (1985) algorithm.
//
// Example:
//
// SolveSRP takes a single preference table as an input. The preference table
// is an array of match preferences.
//
// 		prefs := *[]libmatch.MatchPreference{
// 		  {Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
// 		  {Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
// 		  {Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
// 		  {Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
// 		  {Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
// 		  {Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}}
//    }
//
// On success, the return value will be a MatchResult containing the stable the
// mapping between pairs of members.
//
// 		MatchResult{
// 			Mapping: map[string]string{
// 				"A": "F",
// 				"B": "E",
// 				"C": "D",
// 				"D": "C",
// 				"E": "B",
// 				"F": "A",
// 			},
// 		}
func SolveSRP(prefs *[]MatchPreference) (MatchResult, error) {
	var res MatchResult
	var err error

	table := core.NewPreferenceTable(prefs)
	validator := validate.SingleTableValidator{Prefs: prefs, Table: &table}

	if err = validator.Validate(); err != nil {
		return res, err
	}

	algoCtx := core.AlgorithmContext{
		TableA: &table,
	}

	res, err = srp.Run(algoCtx)

	return res, err
}
