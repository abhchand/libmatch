<div align="center">
  <h1>libmatch</h1>

  <a href="https://github.com/abhchand/libmatch">
    <img
      width="100"
      alt="libmatch"
      src="https://raw.githubusercontent.com/abhchand/libmatch/master/meta/logo.png"
    />
  </a>

  <p>libmatch is a library for solving matching problems.</p>
</div>

---

[![Build Status][ci-badge]][ci]

`libmatch` can be used as a **Go package** ([docs](https://pkg.go.dev/github.com/abhchand/libmatch)) or as a **standalone executable** (CLI).

It supports solving the following problems:

| Matching Problem | Shorthand | Description | Example |
|---|---|---|---|
| [Stable Marriage Problem](https://en.wikipedia.org/wiki/Stable_marriage_problem) | `SMP` | Matching between two groups of members | [example→](https://pkg.go.dev/github.com/abhchand/libmatch#example-SolveSMP) |
| [Stable Roommates Problem](https://en.wikipedia.org/wiki/Stable_roommates_problem) | `SRP` | Matching within a group of members | [example→](https://pkg.go.dev/github.com/abhchand/libmatch#example-SolveSRP) |

---

- [What Does This Do?](#what-does-this-do)
- [Go Package](#go-package)
- [CLI](#cli)
  * [Installation](#installation)
  * [Usage](#usage)
- [Miscellaneous](#miscellaneous)


## <a name="what-does-this-do">What Does This Do?

Matching algorithms find an optimal matching between members, given one or more sets of *member preferences*.

<div align="center">
  <img src="https://github.com/abhchand/libmatch/raw/main/meta/matching.png" width="400px" />
</div>

`libmatch` provides solutions to solve this and variations of this classic matching problem, which have a wide range of real-world applications.

## <a name="go-package">Go Package

[View Go Package Documentation](https://pkg.go.dev/github.com/abhchand/libmatch#section-documentation).

Add `libmatch`:

```shell
go get github.com/abhchand/libmatch
```

Use it:

```go
package main

import (
  "fmt"
  "os"

  "github.com/abhchand/libmatch"
)

func main() {
  prefTable := []libmatch.MatchPreference{
    {Name: "A", Preferences: []string{"B", "D", "C"}},
    {Name: "B", Preferences: []string{"D", "A", "C"}},
    {Name: "C", Preferences: []string{"D", "A", "B"}},
    {Name: "D", Preferences: []string{"C", "A", "B"}},
  }

  // Call the solver for the type of matching algorithm you'd like to solve.
  // In this case `SolveSRP` solves the "Stable Roommate Problem".
  result, err := libmatch.SolveSRP(&prefTable)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // => MatchResult{
  //   Mapping: map[string]string{
  //     "A": "F",
  //     "B": "E",
  //     "C": "D",
  //     "D": "C",
  //     "E": "B",
  //     "F": "A",
  //   }
  // }
}
```

## <a name="cli">CLI

### <a name="installation"></a>Installation

Download the the [latest release](https://github.com/abhchand/libmatch/releases/latest) for your platform.

Or alternatively, build it from source:

```shell
git clone git@github.com:abhchand/libmatch.git
cd libmatch/

make all
```

### <a name="usage">Usage

Define member preferences as JSON data:

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

See `libmatch --help` for more options and examples.

## <a name="miscellaneous">Miscellaneous

* [Create an issue](https://github.com/abhchand/libmatch/issues/new) to report a bug or request a feature
* Contributions are welcome! Please [open an Issue](https://github.com/abhchand/libmatch/issues/new) to discuss your changes first
* The Changelog can be found in the [releases](https://github.com/abhchand/libmatch/releases)

[ci-badge]:
  https://github.com/abhchand/libmatch/actions/workflows/test.yml/badge.svg?branch=main
[ci]:
  https://github.com/abhchand/libmatch/actions
