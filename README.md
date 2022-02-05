# libmatch

A utility library for solving matching problems.

`libmatch` supports solving the following classes of problems:

| Matching Problem | Shorthand | Description |
|---|---|---|
| [Stable Marriage Problem](https://en.wikipedia.org/wiki/Stable_marriage_problem) | `SMP` | Matching between two groups of members |
| [Stable Roommates Problem](https://en.wikipedia.org/wiki/Stable_roommates_problem) | `SRP` | Matching within a group of members |

`libmatch` can be used as a **standalone executable** or included **as a library** in your Go program.

# Overview

All algorithms take one or more sets of *preference tables* as an input.

These tables define an ordered preference for each *member*, whose name can be any unique `string` value.

<div align="center">
  <img src="https://github.com/abhchand/libmatch/raw/main/meta/matching.png" width="400px" />
</div>

The output represents the mathematically optimal matching ("stable matching"), where no two members both prefer each other over their current matches.

# Usage

## As a utility

Download the `libmatch` release:

```shell
tbd
```

Run it

```shell
$ cat <<EOF > input.json
[
  { "name":"A", "preferences": ["B", "D", "F", "C", "E"] },
  { "name":"B", "preferences": ["D", "E", "F", "A", "C"] },
  { "name":"C", "preferences": ["D", "E", "F", "A", "B"] },
  { "name":"D", "preferences": ["F", "C", "A", "E", "B"] },
  { "name":"E", "preferences": ["F", "C", "D", "B", "A"] },
  { "name":"F", "preferences": ["A", "B", "D", "C", "E"] }
]
EOF

$ libmatch solve --algorithm SRP --file input.json
A,F
B,E
C,D
D,C
E,B
F,A
```

See `libmatch --help` for more options and examples

## As a Go Library

Add `libmatch`:

```shell
go get github.com/abhchand/libmatch
```

Call the `libmatch` API

```go
package main

import (
  "fmt"
  "os"

  "github.com/abhchand/libmatch"
)

func main() {
  prefTable := []libmatch.MatchEntry{
    {Name: "A", Preferences: []string{"B", "D", "F", "C", "E"}},
    {Name: "B", Preferences: []string{"D", "E", "F", "A", "C"}},
    {Name: "C", Preferences: []string{"D", "E", "F", "A", "B"}},
    {Name: "D", Preferences: []string{"F", "C", "A", "E", "B"}},
    {Name: "E", Preferences: []string{"F", "C", "D", "B", "A"}},
    {Name: "F", Preferences: []string{"A", "B", "D", "C", "E"}}
  }

  result, err := libmatch.SolveSRP(&prefTable)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for x, y := result.Mapping() {
    fmt.Printf("%v => %v", x, y)
  }
}
```

See documentation [`libmatch.go`](libmatch.go) for the full available API.

# Miscellaneous

* [Create an issue](https://github.com/abhchand/libmatch/issues/new) to report a bug or request a feature
* Contributions are welcome! Please [open an Issue](https://github.com/abhchand/libmatch/issues/new) to discuss your changes first
* The Changelog can be found in the [releases](https://github.com/abhchand/libmatch/releases)
