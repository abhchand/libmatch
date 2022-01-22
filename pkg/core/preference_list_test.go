package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMembers(t *testing.T) {
	memA = Member{name: "A"}
	memB = Member{name: "B"}
	memC = Member{name: "C"}

	plA = PreferenceList{members: []*Member{&memB, &memC}}
	memA.SetPreferenceList(&plA)

	assert.Equal(t, []*Member{&memB, &memC}, plA.Members())
}

func TestRemove(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}
		memC = Member{name: "C"}
		memD = Member{name: "D"}

		plA = PreferenceList{members: []*Member{&memB, &memC, &memD}}
		memA.SetPreferenceList(&plA)

		plA.Remove(memC)
		assert.Equal(t, []*Member{&memB, &memD}, plA.members)

		plA.Remove(memD)
		assert.Equal(t, []*Member{&memB}, plA.members)

		plA.Remove(memB)
		assert.Equal(t, []*Member{}, plA.members)
	})

	t.Run("handles missing member", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}
		memC = Member{name: "C"}

		plA = PreferenceList{members: []*Member{&memB}}
		memA.SetPreferenceList(&plA)

		// Removing a missing element raises no error, just returns
		plA.Remove(memC)
		assert.Equal(t, []*Member{&memB}, plA.members)
	})
}