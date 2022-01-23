package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSingleTableValidator(t *testing.T) {
	setupMembers()

	v := NewSingleTableValidator(pt)
	err := v.Validate()

	assert.Nil(t, err)
}
