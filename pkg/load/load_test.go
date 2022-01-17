package load

import (
	"strings"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromIO(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B", "C", "D"] },
    { "name":"B", "preferences": ["A", "C", "D"] },
    { "name":"C", "preferences": ["A", "B", "D"] },
    { "name":"D", "preferences": ["A", "B", "C"] }
  ]
	`

	got, err := LoadFromIO(strings.NewReader(body))

	wanted := &[]core.MatchEntry{
		{Name: "A", Preferences: []string{"B", "C", "D"}},
		{Name: "B", Preferences: []string{"A", "C", "D"}},
		{Name: "C", Preferences: []string{"A", "B", "D"}},
		{Name: "D", Preferences: []string{"A", "B", "C"}},
	}

	assert.Nil(t, err)
	assert.Equal(t, wanted, got)
}

func TestLoadFromIO_UnmarshallError(t *testing.T) {

	// Note missing `:` on final row
	body := `
  [
    { "name":"A", "preferences": ["B", "C", "D"] },
    { "name":"B", "preferences": ["A", "C", "D"] },
    { "name":"C", "preferences": ["A", "B", "D"] },
    { "name":"D", "preferences" ["A", "B", "C"] }
  ]
	`

	_, err := LoadFromIO(strings.NewReader(body))

	if assert.NotNil(t, err) {
		assert.Equal(t, "invalid character '[' after object key", err.Error())
	}
}
