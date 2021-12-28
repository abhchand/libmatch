package version

import "testing"

func TestPrint(t *testing.T) {
    want := "v0.0.1"

    if got := Print(); got != want {
        t.Errorf("Print() = %q, want %q", got, want)
    }
}
