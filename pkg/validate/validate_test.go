package validate

import (
	"fmt"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setupMembers()

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("empty table", func(t *testing.T) {
		pt := core.PreferenceTable{}

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		if assert.NotNil(t, err) {
			assert.Equal(t, "Table must be non-empty", err.Error())
		}
	})

	t.Run("odd number of members", func(t *testing.T) {
		setupMembers()

		pt = core.PreferenceTable{
			"A": &memA,
			"B": &memB,
			"C": &memC,
		}

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		if assert.NotNil(t, err) {
			assert.Equal(t, "Table must have an even number of members", err.Error())
		}
	})

	t.Run("empty member", func(t *testing.T) {
		setupMembers()

		pt = core.PreferenceTable{
			"A": &memA,
			"B": &memB,
			"":  &memC,
			"D": &memD,
		}

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		if assert.NotNil(t, err) {
			assert.Equal(t, "All member names must non-blank", err.Error())
		}
	})

	t.Run("member names are case sensitive", func(t *testing.T) {
		setupMembers()

		memA = core.NewMember("A")
		memB = core.NewMember("B")
		memC = core.NewMember("C")
		memA_ := core.NewMember("a")

		plA = core.NewPreferenceList([]*core.Member{&memB, &memC, &memA_})
		plB = core.NewPreferenceList([]*core.Member{&memA, &memC, &memA_})
		plC = core.NewPreferenceList([]*core.Member{&memA, &memB, &memA_})
		plA_ := core.NewPreferenceList([]*core.Member{&memA, &memB, &memC})

		memA.SetPreferenceList(&plA)
		memB.SetPreferenceList(&plB)
		memC.SetPreferenceList(&plC)
		memA_.SetPreferenceList(&plA_)

		pt = core.PreferenceTable{
			"A": &memA,
			"B": &memB,
			"C": &memC,
			"a": &memA_,
		}

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("asymmetrical empty list", func(t *testing.T) {
		setupMembers()

		plA_ := core.NewPreferenceList([]*core.Member{})
		memA.SetPreferenceList(&plA_)

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		if assert.NotNil(t, err) {
			wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "A")
			assert.Equal(t, wanted, err.Error())
		}
	})

	t.Run("asymmetrical mismatched list", func(t *testing.T) {
		setupMembers()

		plA_ := core.NewPreferenceList([]*core.Member{&memB, &memC})
		memA.SetPreferenceList(&plA_)

		v := Validator{PrimaryTable: pt}
		err := v.Validate()

		if assert.NotNil(t, err) {
			wanted := fmt.Sprintf("Preference list for '%v' does not contain all the required members", "A")
			assert.Equal(t, wanted, err.Error())
		}
	})
}
