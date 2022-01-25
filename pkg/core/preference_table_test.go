package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPreferenceTable(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entries := []MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		setupMembers()
		wanted := pt

		assert.True(t, reflect.DeepEqual(wanted, NewPreferenceTable(&entries)))
	})

	t.Run("empty table", func(t *testing.T) {
		entries := []MatchEntry{}

		wanted := PreferenceTable{}

		assert.Equal(t, wanted, NewPreferenceTable(&entries))
	})

	t.Run("case sensitive", func(t *testing.T) {
		entries := []MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
			{Name: "a", Preferences: []string{"A", "B", "C"}},
		}

		setupMembers()

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

		assert.Equal(t, wanted, NewPreferenceTable(&entries))
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
			entries := []MatchEntry{
				{Name: "A", Preferences: []string{"B", "C", "D"}},
				{Name: "B", Preferences: []string{"A", "C", "D"}},
				{Name: "C", Preferences: []string{"A", "B", "D"}},
				{Name: "D", Preferences: []string{"A", "B", "C"}},
			}

			pt = NewPreferenceTable(&entries)

			if len(testCase[1]) > 0 {
				proposed := testCase[1]
				proposer := testCase[2]
				pt[proposed].Accept(pt[proposer])
			}

			assert.Equal(t, testCase[3], pt.String())
		})
	}

	t.Run("empty table", func(t *testing.T) {
		entries := []MatchEntry{}

		pt = NewPreferenceTable(&entries)

		wanted := ""

		assert.Equal(t, wanted, pt.String())
	})
}

func TestUnmatchedMembers(t *testing.T) {
	setupMembers()

	memA.acceptedProposalFrom = &memB
	memB.acceptedProposalFrom = &memA
	memC.acceptedProposalFrom = &memD
	memD.acceptedProposalFrom = nil

	assert.Equal(t, []*Member{&memC}, pt.UnmatchedMembers())
}

func TestIsComplete(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		setupMembers()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{&memA}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.True(t, pt.IsComplete())
	})

	t.Run("returns false", func(t *testing.T) {
		setupMembers()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{&memA, &memD}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.False(t, pt.IsComplete())
	})

	t.Run("handles empty lists", func(t *testing.T) {
		setupMembers()

		plA = PreferenceList{members: []*Member{&memB}}
		plB = PreferenceList{members: []*Member{&memA}}
		plC = PreferenceList{members: []*Member{}}
		plD = PreferenceList{members: []*Member{&memA}}

		assert.False(t, pt.IsComplete())
	})
}
