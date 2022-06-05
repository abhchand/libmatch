// Package load is responsible for loading preference data from streams and
// files
package load

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/abhchand/libmatch/pkg/core"
)

// LoadFromFile loads preference data from a file containing JSON data.
//
// The structure of the JSON file should be of format:
//
//    [
//      { "name":"A", "preferences": ["B", "D", "F", "C", "E"] },
//      { "name":"B", "preferences": ["D", "E", "F", "A", "C"] },
//      { "name":"C", "preferences": ["D", "E", "F", "A", "B"] },
//      { "name":"D", "preferences": ["F", "C", "A", "E", "B"] },
//      { "name":"E", "preferences": ["F", "C", "D", "B", "A"] },
//      { "name":"F", "preferences": ["A", "B", "D", "C", "E"] },
//    ]
//
// The return value is an array of `MatchPreference` structs containing the
// loaded JSON data
//
//    *[]core.MatchPreference{
//      {Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
//      {Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
//      {Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
//      {Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
//      {Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
//      {Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}},
//    }
func LoadFromFile(filename string) (*[]core.MatchPreference, error) {
	var data *[]core.MatchPreference

	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	data, err = LoadFromIO(bufio.NewReader(file))
	return data, err
}

// LoadFromIO reads match preference data from an `io.Reader`.
//
// The expected data is a JSON formatted preference table of the format:
//
//    [
//      { "name":"A", "preferences": ["B", "D", "F", "C", "E"] },
//      { "name":"B", "preferences": ["D", "E", "F", "A", "C"] },
//      { "name":"C", "preferences": ["D", "E", "F", "A", "B"] },
//      { "name":"D", "preferences": ["F", "C", "A", "E", "B"] },
//      { "name":"E", "preferences": ["F", "C", "D", "B", "A"] },
//      { "name":"F", "preferences": ["A", "B", "D", "C", "E"] },
//    ]
//
// The return value is an array of `MatchPreference` structs containing the
// loaded JSON data
//
//    *[]libmatch.MatchPreference{
//      {Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
//      {Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
//      {Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
//      {Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
//      {Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
//      {Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}},
//    }
func LoadFromIO(r io.Reader) (*[]core.MatchPreference, error) {
	var data []core.MatchPreference

	rawJson, err := io.ReadAll(r)
	if err != nil {
		return &data, err
	}

	if err := json.Unmarshal(rawJson, &data); err != nil {
		return &data, err
	}

	return &data, nil
}
