# libmatch

![build status](https://github.com/abhchand/libmatch/actions/workflows/test.yml/badge.svg?branch=main)

A utility library for solving matching problems.

`libmatch` supports solving the following classes of problems:

| Matching Problem | Shorthand | Description |
|---|---|---|
| [Stable Marriage Problem](https://en.wikipedia.org/wiki/Stable_marriage_problem) | `SMP` | Matching between two groups of members |
| [Stable Roommates Problem](https://en.wikipedia.org/wiki/Stable_roommates_problem) | `SRP` | Matching within a group of members |

`libmatch` can be used as a **standalone executable** or included **as a library** in your Go program.

---

- [Overview](#overview)
- [Installation](#installation)
  * [Build from source](#build-from-source)
- [Usage](#usage)
  * [As an executable](#as-an-executable)
  * [As a Go Library](#as-a-go-library)
- [Miscellaneous](#miscellaneous)

## <a name="overview">Overview

Given one or more sets of *member preferences*, `libmatch` finds an optimal matching between members.

<div align="center">
  <img src="https://github.com/abhchand/libmatch/raw/main/meta/matching.png" width="400px" />
</div>

The output will be mathematically "stable", meaning no two members will prefer each other over their existing matches.

## <a name="installation"></a>Installation

Download the the [latest release](https://github.com/abhchand/libmatch/releases/latest) for your platform.

### <a name="build-from-source"></a>Build from source

```shell
git clone git@github.com:abhchand/libmatch.git
cd libmatch/

make all
```

## <a name="usage">Usage

### <a name="as-an-executable">As an executable


Define member preferences as a JSON input file:

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
```

Run `libmatch`:

```shell
$ libmatch solve --algorithm SRP --file input.json
A,F
B,E
C,D
D,C
E,B
F,A
```

See `libmatch --help` for more options and examples

### <a name="as-a-go-library">As a Go Library

Add `libmatch` as a dependency:

```shell
go get github.com/abhchand/libmatch
```

Then `import` and use it:

```go
package main

import (
  "fmt"
  "os"

  "github.com/abhchand/libmatch"
)

func main() {
  prefTable := []libmatch.MatchEntry{
    {Name: "A", Preferences: []string{"B", "D", "C"}},
    {Name: "B", Preferences: []string{"D", "A", "C"}},
    {Name: "C", Preferences: []string{"D", "A", "B"}},
    {Name: "D", Preferences: []string{"C", "A", "B"}},
  }

  // libmatch returns a `libmatch.MatchResult` type.
  // The mapping can be accessed via the `.Mapping` property
  result, err := libmatch.SolveSRP(&prefTable)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for x, y := range result.Mapping {
    fmt.Printf("%v => %v\n", x, y)
  }
}
```

See [the `examples/` directory](examples/) for a full list of working examples

## <a name="miscellaneous">Miscellaneous

* [Create an issue](https://github.com/abhchand/libmatch/issues/new) to report a bug or request a feature
* Contributions are welcome! Please [open an Issue](https://github.com/abhchand/libmatch/issues/new) to discuss your changes first
* The Changelog can be found in the [releases](https://github.com/abhchand/libmatch/releases)
