package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPreferenceTable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		prefs := []MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		setupSingleTable()
		wanted := pt

		assert.True(t, reflect.DeepEqual(wanted, NewPreferenceTable(&prefs)))
	})

	t.Run("empty table", func(t *testing.T) {
		prefs := []MatchPreference{}

		wanted := PreferenceTable{}

		assert.Equal(t, wanted, NewPreferenceTable(&prefs))
	})

	t.Run("undefined preference", func(t *testing.T) {
		prefs := []MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "X"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		setupSingleTable()

		plA = PreferenceList{members: []*Member{&memB, &memC, nil}}
		memA.SetPreferenceList(&plA)

		assert.True(t, reflect.DeepEqual(pt, NewPreferenceTable(&prefs)))
	})

	t.Run("case sensitive", func(t *testing.T) {
		prefs := []MatchPreference{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
			{Name: "a", Preferences: []string{"A", "B", "C"}},
		}

		setupSingleTable()

		_memA := Member{name: "a"}
		_plA := PreferenceList{members: []*Member{&memA, &memB, &memC}}
		_memA.SetPreferenceList(&_plA)

		wanted := PreferenceTable{
			"A": &memA,
			"B": &memB,
			"C": &memC,
			"D": &memD,
			"a": &_memA,
		}

		assert.Equal(t, wanted, NewPreferenceTable(&prefs))
	})
}

func TestNewPreferenceTablePair(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		prefsA := []MatchPreference{
			{Name: "A", Preferences: []string{"K", "L", "M"}},
			{Name: "B", Preferences: []string{"L", "M", "K"}},
			{Name: "C", Preferences: []string{"M", "L", "K"}},
		}
		prefsB := []MatchPreference{
			{Name: "K", Preferences: []string{"B", "C", "A"}},
			{Name: "L", Preferences: []string{"A", "C", "B"}},
			{Name: "M", Preferences: []string{"A", "B", "C"}},
		}

		setupDoubleTable()

		tables := NewPreferenceTablePair(&prefsA, &prefsB)

		assert.True(t, reflect.DeepEqual(ptA, tables[0]))
		assert.True(t, reflect.DeepEqual(ptB, tables[1]))
	})

	t.Run("order of prefs does not matter", func(t *testing.T) {
		prefsA := []MatchPreference{
			{Name: "A", Preferences: []string{"K", "L", "M"}},
			{Name: "B", Preferences: []string{"L", "M", "K"}},
			{Name: "C", Preferences: []string{"M", "L", "K"}},
		}
		prefsB := []MatchPreference{
			{Name: "K", Preferences: []string{"B", "C", "A"}},
			{Name: "L", Preferences: []string{"A", "C", "B"}},
			{Name: "M", Preferences: []string{"A", "B", "C"}},
		}

		setupDoubleTable()

		tables := NewPreferenceTablePair(&prefsB, &prefsA)

		assert.True(t, reflect.DeepEqual(ptB, tables[0]))
		assert.True(t, reflect.DeepEqual(ptA, tables[1]))
	})

	t.Run("empty table", func(t *testing.T) {
		prefsA := []MatchPreference{}
		prefsB := []MatchPreference{}

		tables := NewPreferenceTablePair(&prefsA, &prefsB)

		assert.True(t, reflect.DeepEqual(PreferenceTable{}, tables[0]))
		assert.True(t, reflect.DeepEqual(PreferenceTable{}, tables[1]))
	})

	t.Run("undefined preference", func(t *testing.T) {
		prefsA := []MatchPreference{
			{Name: "A", Preferences: []string{"K", "L", "X"}},
			{Name: "B", Preferences: []string{"L", "M", "K"}},
			{Name: "C", Preferences: []string{"M", "L", "K"}},
		}
		prefsB := []MatchPreference{
			{Name: "K", Preferences: []string{"B", "C", "A"}},
			{Name: "L", Preferences: []string{"A", "Z", "B"}},
			{Name: "M", Preferences: []string{"A", "B", "C"}},
		}

		setupDoubleTable()

		plA = PreferenceList{members: []*Member{&memK, &memL, nil}}
		memA.SetPreferenceList(&plA)

		plL = PreferenceList{members: []*Member{&memA, nil, &memB}}
		memL.SetPreferenceList(&plL)

		tables := NewPreferenceTablePair(&prefsA, &prefsB)

		assert.True(t, reflect.DeepEqual(ptA, tables[0]))
		assert.True(t, reflect.DeepEqual(ptB, tables[1]))
	})

	t.Run("case sensitive", func(t *testing.T) {
		prefsA := []MatchPreference{
			{Name: "A", Preferences: []string{"K", "L", "M"}},
			{Name: "B", Preferences: []string{"L", "M", "K"}},
			{Name: "C", Preferences: []string{"M", "L", "K"}},
			{Name: "a", Preferences: []string{"L", "K", "M"}},
		}
		prefsB := []MatchPreference{
			{Name: "K", Preferences: []string{"B", "C", "A"}},
			{Name: "L", Preferences: []string{"A", "C", "B"}},
			{Name: "M", Preferences: []string{"A", "B", "C"}},
			{Name: "k", Preferences: []string{"C", "B", "A"}},
		}

		setupDoubleTable()

		_memA := Member{name: "a"}
		_plA := PreferenceList{members: []*Member{&memL, &memK, &memM}}
		_memA.SetPreferenceList(&_plA)

		_memK := Member{name: "k"}
		_plK := PreferenceList{members: []*Member{&memC, &memB, &memA}}
		_memK.SetPreferenceList(&_plK)

		ptA = PreferenceTable{
			"A": &memA,
			"B": &memB,
			"C": &memC,
			"a": &_memA,
		}

		ptB = PreferenceTable{
			"K": &memK,
			"L": &memL,
			"M": &memM,
			"k": &_memK,
		}

		tables := NewPreferenceTablePair(&prefsA, &prefsB)

		assert.True(t, reflect.DeepEqual(ptA, tables[0]))
		assert.True(t, reflect.DeepEqual(ptB, tables[1]))
	})
}

func TestString__PreferenceTable(t *testing.T) {

	cases := [][]string{
		{
			"success",
			"",
			"",
			"'A'\t=>\t'B', 'C', 'D'\n'B'\t=>\t'A', 'C', 'D'\n'C'\t=>\t'A', 'B', 'D'\n'D'\t=>\t'A', 'B', 'C'\n",
		},
		{
			"displays current proposer (middle element)",
			"A",
			"C",
			"'A'\t=>\t'B', 'C'+, 'D'\n'B'\t=>\t'A', 'C', 'D'\n'C'\t=>\t'A', 'B', 'D'\n'D'\t=>\t'A', 'B', 'C'\n",
		},
		{
			"displays current proposer (first element)",
			"A",
			"B",
			"'A'\t=>\t'B'+, 'C', 'D'\n'B'\t=>\t'A', 'C', 'D'\n'C'\t=>\t'A', 'B', 'D'\n'D'\t=>\t'A', 'B', 'C'\n",
		},
		{
			"displays current proposer (last element)",
			"A",
			"D",
			"'A'\t=>\t'B', 'C', 'D'+\n'B'\t=>\t'A', 'C', 'D'\n'C'\t=>\t'A', 'B', 'D'\n'D'\t=>\t'A', 'B', 'C'\n",
		},
	}

	for i := range cases {
		testCase := cases[i]

		t.Run(testCase[0], func(t *testing.T) {
			prefs := []MatchPreference{
				{Name: "A", Preferences: []string{"B", "C", "D"}},
				{Name: "B", Preferences: []string{"A", "C", "D"}},
				{Name: "C", Preferences: []string{"A", "B", "D"}},
				{Name: "D", Preferences: []string{"A", "B", "C"}},
			}

			pt = NewPreferenceTable(&prefs)

			if len(testCase[1]) > 0 {
				proposed := testCase[1]
				proposer := testCase[2]
				pt[proposed].Accept(pt[proposer])
			}

			assert.Equal(t, testCase[3], pt.String())
		})
	}

	t.Run("empty table", func(t *testing.T) {
		prefs := []MatchPreference{}

		pt = NewPreferenceTable(&prefs)

		wanted := ""

		assert.Equal(t, wanted, pt.String())
	})
}

func TestUnmatchedMembers(t *testing.T) {
	setupSingleTable()

	memA.acceptedProposalFrom = &memB
	memB.acceptedProposalFrom = &memA
	memC.acceptedProposalFrom = &memD
	memD.acceptedProposalFrom = nil

	assert.Equal(t, []*Member{&memC}, pt.UnmatchedMembers())
}

func TestIsStable(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		setupSingleTable()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{&memA, &memD}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.True(t, pt.IsStable())
	})

	t.Run("returns false", func(t *testing.T) {
		setupSingleTable()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.False(t, pt.IsStable())
	})
}

func TestIsComplete(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		setupSingleTable()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{&memA}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.True(t, pt.IsComplete())
	})

	t.Run("returns false", func(t *testing.T) {
		setupSingleTable()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{&memA, &memD}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.False(t, pt.IsComplete())
	})

	t.Run("handles empty lists", func(t *testing.T) {
		setupSingleTable()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.False(t, pt.IsComplete())
	})
}
