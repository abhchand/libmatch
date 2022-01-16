package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatted(t *testing.T) {
	assert.Equal(t, "v0.0.1", Formatted())
}

func TestVersion(t *testing.T) {
	assert.Equal(t, "0.0.1", Version())
}
