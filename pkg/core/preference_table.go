package core

type PreferenceTable map[string]PreferenceList

func NewPreferenceTable(entries *[]MatchEntry) PreferenceTable {
	e := *entries
	table := make(PreferenceTable)

	for i := range e {
		name := e[i].Name
		preferences := PreferenceList{e[i].Preferences}

		table[name] = preferences
	}

	return table
}
