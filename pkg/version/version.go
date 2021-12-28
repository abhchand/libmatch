package version

import (
  "fmt"
)

var version = "0.0.1"

func Print() {
  fmt.Printf("v%v\n", version)
}
