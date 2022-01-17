package core

type MatchEntry struct {
	Name        string   `json:"name"`
	Preferences []string `json:"preferences"`
}

type PreferenceList struct {
	Members []string
}
