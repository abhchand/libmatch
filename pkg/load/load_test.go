package load

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/abhchand/libmatch/pkg/core"

	"github.com/stretchr/testify/assert"
)

var testFile = "/tmp/libmatch_test.json"

func TestLoadFromFile(t *testing.T) {
	body := `
  [
    { "name":"A", "preferences": ["B", "C", "D"] },
    { "name":"B", "preferences": ["A", "C", "D"] },
    { "name":"C", "preferences": ["A", "B", "D"] },
    { "name":"D", "preferences": ["A", "B", "C"] }
  ]
	`
	writeToFile(testFile, body)

	got, err := LoadFromFile(testFile)

	wanted := &[]core.MatchPreference{
		{Name: "A", Preferences: []string{"B", "C", "D"}},
		{Name: "B", Preferences: []string{"A", "C", "D"}},
		{Name: "C", Preferences: []string{"A", "B", "D"}},
		{Name: "D", Preferences: []string{"A", "B", "C"}},
	}

	assert.Nil(t, err)
	assert.Equal(t, wanted, got)
}

func TestLoadFromFile_DoesNotExist(t *testing.T) {
	badFile := "/tmp/badfile.json"

	_, err := LoadFromFile(badFile)

	if assert.NotNil(t, err) {
		assert.Equal(t,
			fmt.Sprintf("open %v: no such file or directory", badFile), err.Error())
	}
}

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

	wanted := &[]core.MatchPreference{
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

func writeToFile(filename, body string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(body)
	if err != nil {
		fmt.Printf("Could not create file: %s\n", err.Error())
	}

	writer.Flush()
}
