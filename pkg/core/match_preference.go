package core

// MatchPreference stores information about a member and their preference list
// of other members.
//
// This data model is designed to be a container for information read directly
// from a stream of JSON data (e.g. a file on disk).
type MatchPreference struct {
	Name        string   `json:"name"`
	Preferences []string `json:"preferences"`
}
