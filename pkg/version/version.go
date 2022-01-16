package version

import (
	"fmt"
)

var version = "0.0.1"

func Formatted() string {
	return fmt.Sprintf("v%v", version)
}
