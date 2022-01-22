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
