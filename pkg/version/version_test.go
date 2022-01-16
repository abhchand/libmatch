package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatted(t *testing.T) {
	assert.Equal(t, "v0.0.1", Formatted())
}
