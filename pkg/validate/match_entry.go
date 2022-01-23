package validate

import (
	"errors"
	"fmt"

	"github.com/abhchand/libmatch/pkg/core"
)

type MatchEntryValidator struct {
	Entries *[]core.MatchEntry
	Err     error
}

func (v *MatchEntryValidator) Validate() error {
	var err error

	err = v.validateUniqueness(v.Entries)
	if err != nil {
		return err
	}

	return nil
}

func (v MatchEntryValidator) validateUniqueness(entries *[]core.MatchEntry) error {
	cache := make(map[string]bool, 0)

	for i := range *entries {
		name := (*entries)[i].Name

		if cache[name] {
			msg := fmt.Sprintf("Member names must be unique. Found duplicate entry '%v'", name)
			return errors.New(msg)
		}

		cache[name] = true
	}

	return nil
}
