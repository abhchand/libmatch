package core

type MatchEntry struct {
	Name        string   `json:"name"`
	Preferences []string `json:"preferences"`
}

type PreferenceTable map[string]PreferenceList

type PreferenceList struct {
	Members []string
}
