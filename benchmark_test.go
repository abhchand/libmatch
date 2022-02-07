package libmatch

import (
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
)

func BenchmarkSolveSMP(b *testing.B) {
	prefsA := []core.MatchPreference{
		{Name: "A", Preferences: []string{"R", "L", "M", "S", "Q", "T", "P", "N", "O", "K"}},
		{Name: "B", Preferences: []string{"S", "O", "N", "K", "T", "M", "Q", "P", "R", "L"}},
		{Name: "C", Preferences: []string{"N", "O", "M", "S", "L", "K", "R", "P", "T", "Q"}},
		{Name: "D", Preferences: []string{"S", "Q", "T", "N", "K", "P", "R", "L", "M", "O"}},
		{Name: "E", Preferences: []string{"P", "S", "T", "N", "R", "M", "Q", "K", "O", "L"}},
		{Name: "F", Preferences: []string{"Q", "N", "K", "M", "S", "L", "P", "O", "R", "T"}},
		{Name: "G", Preferences: []string{"T", "R", "K", "Q", "N", "M", "S", "P", "L", "O"}},
		{Name: "H", Preferences: []string{"N", "Q", "P", "L", "M", "O", "S", "K", "T", "R"}},
		{Name: "I", Preferences: []string{"K", "O", "M", "L", "Q", "N", "S", "P", "T", "R"}},
		{Name: "J", Preferences: []string{"L", "N", "Q", "S", "T", "K", "P", "R", "O", "M"}},
	}

	prefsB := []core.MatchPreference{
		{Name: "K", Preferences: []string{"C", "B", "F", "A", "J", "I", "G", "D", "H", "E"}},
		{Name: "L", Preferences: []string{"F", "J", "D", "I", "H", "E", "A", "C", "G", "B"}},
		{Name: "M", Preferences: []string{"D", "J", "H", "C", "F", "G", "B", "I", "E", "A"}},
		{Name: "N", Preferences: []string{"B", "F", "D", "A", "H", "G", "J", "E", "I", "C"}},
		{Name: "O", Preferences: []string{"D", "J", "F", "A", "H", "B", "C", "I", "E", "G"}},
		{Name: "P", Preferences: []string{"A", "D", "C", "B", "J", "G", "I", "H", "F", "E"}},
		{Name: "Q", Preferences: []string{"J", "E", "I", "A", "F", "H", "G", "B", "C", "D"}},
		{Name: "R", Preferences: []string{"B", "H", "J", "A", "C", "I", "G", "F", "D", "E"}},
		{Name: "S", Preferences: []string{"E", "G", "B", "D", "C", "I", "H", "F", "A", "J"}},
		{Name: "T", Preferences: []string{"H", "C", "A", "F", "G", "B", "D", "E", "J", "I"}},
	}

	for i := 0; i < b.N; i++ {
		SolveSMP(&prefsA, &prefsB)
	}
}

func BenchmarkSolveSRP(b *testing.B) {
	prefs := []core.MatchPreference{
		{Name: "A", Preferences: []string{"H", "J", "E", "B", "D", "I", "C", "G", "F"}},
		{Name: "B", Preferences: []string{"E", "I", "G", "D", "A", "J", "C", "F", "H"}},
		{Name: "C", Preferences: []string{"J", "A", "B", "H", "F", "I", "G", "D", "E"}},
		{Name: "D", Preferences: []string{"I", "C", "E", "G", "B", "A", "J", "F", "H"}},
		{Name: "E", Preferences: []string{"F", "J", "G", "B", "C", "H", "A", "D", "I"}},
		{Name: "F", Preferences: []string{"C", "I", "D", "E", "G", "H", "A", "J", "B"}},
		{Name: "G", Preferences: []string{"E", "H", "F", "A", "J", "C", "D", "B", "I"}},
		{Name: "H", Preferences: []string{"F", "G", "J", "B", "I", "E", "C", "A", "D"}},
		{Name: "I", Preferences: []string{"J", "G", "B", "D", "A", "C", "E", "F", "H"}},
		{Name: "J", Preferences: []string{"C", "F", "A", "B", "I", "G", "H", "D", "E"}},
	}

	for i := 0; i < b.N; i++ {
		SolveSRP(&prefs)
	}
}
