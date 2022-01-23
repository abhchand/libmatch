package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMember(t *testing.T) {
	m := NewMember("A")

	var zero Member
	wanted := reflect.TypeOf(zero).Kind()

	assert.Equal(t, wanted, reflect.TypeOf(m).Kind())
	assert.Equal(t, "A", m.name)
	assert.Nil(t, m.preferenceList)
	assert.Nil(t, m.acceptedProposalFrom)
}

func TestName(t *testing.T) {
	memA = Member{name: "A"}

	assert.Equal(t, "A", memA.Name())
}

func TestCurrentProposer(t *testing.T) {

	t.Run("returns current proposer", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}

		memA.acceptedProposalFrom = &memB

		assert.Equal(t, "B", memA.CurrentProposer().Name())
	})

	t.Run("passed by reference", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}

		memA.acceptedProposalFrom = &memB

		assert.Equal(t, "B", memA.CurrentProposer().Name())
		memB.name = "SomethingElse"
		assert.Equal(t, "SomethingElse", memA.CurrentProposer().Name())
	})
}

func TestCurrentAcceptor(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupMembers()

		// Ensure all users have an accepted proposal for this test. This avoids
		// any nil reference errors for this specific test
		memA.acceptedProposalFrom = &memB
		memB.acceptedProposalFrom = &memA
		memC.acceptedProposalFrom = &memD
		memD.acceptedProposalFrom = &memC

		testCases := map[Member]string{
			memA: "B",
			memB: "A",
			memC: "D",
			memD: "C",
		}

		for member, acceptorName := range testCases {
			assert.Equal(t, acceptorName, member.CurrentAcceptor().Name())
		}
	})

	t.Run("handles members with no accepted proposal", func(t *testing.T) {
		setupMembers()

		// Set only some members to have an accepted proposal
		memA.acceptedProposalFrom = nil
		memB.acceptedProposalFrom = &memA
		memC.acceptedProposalFrom = nil
		memD.acceptedProposalFrom = &memC

		testCases := map[Member]string{
			memA: "B",
			memB: "",
			memC: "D",
			memD: "",
		}

		for member, acceptorName := range testCases {
			if acceptorName == "" {
				assert.Nil(t, member.CurrentAcceptor())
			} else {
				assert.Equal(t, acceptorName, member.CurrentAcceptor().Name())
			}
		}
	})
}

func TestHasAcceptedProposal(t *testing.T) {
	memA = Member{name: "A"}
	memB = Member{name: "B"}

	memA.acceptedProposalFrom = &memB

	assert.True(t, memA.HasAcceptedProposal())
	assert.False(t, memB.HasAcceptedProposal())
}

func TestAccept(t *testing.T) {
	setupMembers()

	assert.Nil(t, memA.acceptedProposalFrom)

	memA.Accept(&memB)
	assert.Equal(t, memB, *memA.acceptedProposalFrom)

	memA.Accept(&memC)
	assert.Equal(t, memC, *memA.acceptedProposalFrom)
}

func TestReject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupMembers()

		memA.acceptedProposalFrom = &memB
		memB.acceptedProposalFrom = &memA
		memC.acceptedProposalFrom = &memD
		memD.acceptedProposalFrom = &memC

		memA.Reject(&memC)

		assert.Equal(t, PreferenceList{members: []*Member{&memB, &memD}}, plA)
		assert.Equal(t, PreferenceList{members: []*Member{&memA, &memC, &memD}}, plB)
		assert.Equal(t, PreferenceList{members: []*Member{&memB, &memD}}, plC)
		assert.Equal(t, PreferenceList{members: []*Member{&memA, &memB, &memC}}, plD)

		assert.Equal(t, &memB, memA.CurrentProposer())
		assert.Equal(t, &memA, memB.CurrentProposer())
		assert.Equal(t, &memD, memC.CurrentProposer())
		assert.Equal(t, &memC, memD.CurrentProposer())
	})

	t.Run("Rejecting current proposer", func(t *testing.T) {
		setupMembers()

		memA.acceptedProposalFrom = &memB
		memB.acceptedProposalFrom = &memA
		memC.acceptedProposalFrom = &memD
		memD.acceptedProposalFrom = &memC

		memA.Reject(&memB)

		assert.Equal(t, PreferenceList{members: []*Member{&memC, &memD}}, plA)
		assert.Equal(t, PreferenceList{members: []*Member{&memC, &memD}}, plB)
		assert.Equal(t, PreferenceList{members: []*Member{&memA, &memB, &memD}}, plC)
		assert.Equal(t, PreferenceList{members: []*Member{&memA, &memB, &memC}}, plD)

		/*
		 * A rejects B (it's current proposer), so A also has that value reset
		 * B does not lose its current proposer, regardless of who is it
		 */
		assert.Nil(t, memA.CurrentProposer())
		assert.Equal(t, &memA, memB.CurrentProposer())
		assert.Equal(t, &memD, memC.CurrentProposer())
		assert.Equal(t, &memC, memD.CurrentProposer())
	})
}

func TestWouldPreferProposalFrom(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupMembers()

		memA.acceptedProposalFrom = &memC

		assert.True(t, memA.WouldPreferProposalFrom(memB))
		assert.False(t, memA.WouldPreferProposalFrom(memD))
	})

	t.Run("has not accepted proposal", func(t *testing.T) {
		setupMembers()

		assert.True(t, memA.WouldPreferProposalFrom(memB))
		assert.True(t, memA.WouldPreferProposalFrom(memD))
	})
}

func TestFirstPreference(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}
		memC = Member{name: "C"}

		plA = PreferenceList{members: []*Member{&memB, &memC}}
		memA.SetPreferenceList(&plA)

		assert.Equal(t, &memB, memA.FirstPreference())
	})

	t.Run("preference list is empty", func(t *testing.T) {
		memA = Member{name: "A"}

		plA = PreferenceList{members: []*Member{}}
		memA.SetPreferenceList(&plA)

		assert.Nil(t, memA.SecondPreference())
	})
}

func TestSecondPreference(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}
		memC = Member{name: "C"}

		plA = PreferenceList{members: []*Member{&memB, &memC}}
		memA.SetPreferenceList(&plA)

		assert.Equal(t, &memC, memA.SecondPreference())
	})

	t.Run("preference list has one element", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}

		plA = PreferenceList{members: []*Member{&memB}}
		memA.SetPreferenceList(&plA)

		assert.Nil(t, memA.SecondPreference())
	})
}

func TestLastPreference(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		memA = Member{name: "A"}
		memB = Member{name: "B"}
		memC = Member{name: "C"}

		plA = PreferenceList{members: []*Member{&memB, &memC}}
		memA.SetPreferenceList(&plA)

		assert.Equal(t, &memC, memA.LastPreference())
	})

	t.Run("preference list is empty", func(t *testing.T) {
		memA = Member{name: "A"}

		plA = PreferenceList{members: []*Member{}}
		memA.SetPreferenceList(&plA)

		assert.Nil(t, memA.LastPreference())
	})
}
