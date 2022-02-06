package libmatch

import (
	"testing"

	"github.com/abhchand/libmatch/pkg/core"
)

func BenchmarkSolveSRP(b *testing.B) {
	entries := []core.MatchEntry{
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
		SolveSRP(&entries)
	}
}
