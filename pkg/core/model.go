package core

type MatchEntry struct {
	Name        string   `json:"name"`
	Preferences []string `json:"preferences"`
}

type MatchResult struct {
	Mapping map[string]string `json:"values"`
}
