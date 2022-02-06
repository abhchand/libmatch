package version

import (
	"fmt"
)

// The build script parses the version from this line.
// Check the regex in `build.sh` before modifying this!
var version = "0.1.0-beta.1"

func Formatted() string {
	return fmt.Sprintf("v%v", version)
}

func Version() string {
	return version
}
