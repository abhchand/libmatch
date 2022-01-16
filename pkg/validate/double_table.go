package validate

import "github.com/abhchand/libmatch/pkg/core"

type doubleTableValidator struct {
	Validator
	SecondaryTable core.PreferenceTable
}

func NewDoubleTableValidator(primaryTable, secondaryTable core.PreferenceTable) doubleTableValidator {
	return doubleTableValidator{
		Validator: Validator{PrimaryTable: primaryTable}, SecondaryTable: secondaryTable}
}
