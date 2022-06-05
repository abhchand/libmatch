// Package version stores version information.
package version

import (
	"fmt"
)

// The build script parses the version from this line.
// Check the regex in `build.sh` before modifying this!
var version = "0.1.0-beta.1"

// Formatted formats the version as a printable string.
func Formatted() string {
	return fmt.Sprintf("v%v", version)
}

// Version returns the current version.
func Version() string {
	return version
}
