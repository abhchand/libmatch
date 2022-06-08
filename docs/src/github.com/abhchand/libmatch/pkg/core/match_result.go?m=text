package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MatchResult stores the result of executing a mapping algorithm
type MatchResult struct {
	Mapping map[string]string `json:"mapping"`
}

// Print prints formatted match results in a sepcified format. The `format` can
// be specified as one of the following:
// 		* csv
//    * json
func (mr MatchResult) Print(format string) error {
	switch format {
	case "csv":
		for a, b := range mr.Mapping {
			fmt.Printf("%v,%v\n", a, b)
		}
	case "json":
		json, _ := json.Marshal(mr)
		fmt.Println(string(json))
	default:
		return errors.New(fmt.Sprintf("Unknown format '%v'", format))
	}

	return nil
}
