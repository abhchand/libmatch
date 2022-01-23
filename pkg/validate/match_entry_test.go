package validate

import (
	"fmt"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestMatchEntryValidator__Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		entries := &[]core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "D", Preferences: []string{"A", "B", "C"}},
		}

		v := MatchEntryValidator{Entries: entries}
		err := v.Validate()

		assert.Nil(t, err)
	})

	t.Run("duplicate member name", func(t *testing.T) {
		entries := &[]core.MatchEntry{
			{Name: "A", Preferences: []string{"B", "C", "D"}},
			{Name: "B", Preferences: []string{"A", "C", "D"}},
			{Name: "C", Preferences: []string{"A", "B", "D"}},
			{Name: "A", Preferences: []string{"A", "B", "C"}},
		}

		v := MatchEntryValidator{Entries: entries}
		err := v.Validate()

		if assert.NotNil(t, err) {
			msg := fmt.Sprintf("Member names must be unique. Found duplicate entry '%v'", "A")
			assert.Equal(t, msg, err.Error())
		}
	})
}
