package core

type PreferenceTable map[string]PreferenceList

type PreferenceList struct {
	Members []string
}
