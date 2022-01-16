package validate

import "github.com/abhchand/libmatch/pkg/core"

type singleTableValidator struct {
	Validator
}

func NewSingleTableValidator(primaryTable core.PreferenceTable) singleTableValidator {
	return singleTableValidator{
		Validator: Validator{PrimaryTable: primaryTable}}
}
