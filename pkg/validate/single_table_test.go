package validate

import (
	"testing"

	"github.com/abhchand/libmatch/pkg/core"

	"github.com/stretchr/testify/assert"
)

func TestNewSingleTableValidator_Validate(t *testing.T) {
	table := core.PreferenceTable{
		"A": core.PreferenceList{Members: []string{"B", "C", "D"}},
		"B": core.PreferenceList{Members: []string{"A", "C", "D"}},
		"C": core.PreferenceList{Members: []string{"A", "B", "D"}},
		"D": core.PreferenceList{Members: []string{"A", "B", "C"}},
	}

	v := NewSingleTableValidator(table)
	err := v.Validate()

	assert.Nil(t, err)
}
