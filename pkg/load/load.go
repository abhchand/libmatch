/*
 * Loads and transforms preference data.
 *
 * The input can be a data structure of JSON file in the following format:
 *
 * Input format:
 *
 *     [
 *         [
 *             { "name": "A", "preferences": ["B", "C", "D"] },
 *             { "name": "B", "preferences": ["A", "C", "D"] },
 *             { "name": "C", "preferences": ["A", "B", "D"] },
 *             { "name": "D", "preferences": ["A", "B", "C"] }
 *         ],
 *         [
 *             { "name": "E", "preferences": ["F", "G", "H"] }
 *         ]
 *     ]
 *
 * The output will be a list of `PreferenceTable` structs, mapping a name
 * to a list of preferences:
 *
 *     [
 *         PreferenceTable{
 *             "A": PreferenceList{"B", "C", "D"},
 *             "B": PreferenceList{"A", "C", "D"},
 *             "C": PreferenceList{"A", "B", "D"},
 *             "D": PreferenceList{"A", "B", "C"}
 *         },
 *         PreferenceTable{
 *             "E": PreferenceList{"F", "G", "H"},
 *         }
 *     ]
 *
 */

package load

import (
	"encoding/json"
	"io"

	"github.com/abhchand/libmatch/pkg/core"
)

func LoadFromIO(r io.Reader) (core.PreferenceTable, error) {
	rawJson, err := io.ReadAll(r)
	if err != nil {
		return make(core.PreferenceTable), nil
	}

	var data []core.MatchEntry

	if err := json.Unmarshal(rawJson, &data); err != nil {
		return make(core.PreferenceTable), nil
	}

	return LoadFromDataType(&data), nil
}

func LoadFromDataType(pEntries *[]core.MatchEntry) core.PreferenceTable {
	entries := *pEntries
	table := make(core.PreferenceTable)

	for i := range entries {
		name := entries[i].Name
		preferences := core.PreferenceList{entries[i].Preferences}

		table[name] = preferences
	}

	return table
}
