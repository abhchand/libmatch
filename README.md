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
  * [Installation](#pkg-installation)
  * [Examples](#pkg-examples)
    * [Stable Marriage Example](#pkg-stable-marriage-example)
    * [Stable Roommates Example](#pkg-stable-roommates-example)
- [CLI](#cli)
  * [Installation](#cliinstallation)
  * [Examples](#cli-examples)
    * [Stable Marriage Example](#cli-stable-marriage-example)
    * [Stable Roommates Example](#cli-stable-roommates-example)
- [Miscellaneous](#miscellaneous)


## <a name="what-does-this-do">What Does This Do?

Matching algorithms find an optimal matching between members, given one or more sets of *member preferences*.

<div align="center">
  <img src="https://github.com/abhchand/libmatch/raw/main/meta/matching.png" width="400px" />
</div>

Algorithmic solutions to these problems have a wide range of real-world applications, from economics to computational mathematics.

`libmatch` provides solutions to solve this and variations of this classic matching problem at scale.

## <a name="go-package">Go Package

[View Go Package Documentation](https://pkg.go.dev/github.com/abhchand/libmatch#section-documentation).

### <a name="pkg-installation"></a>Installation

```shell
go get github.com/abhchand/libmatch
```

### <a name="pkg-examples">Examples

#### <a name="pkg-stable-marriage-example">Stable Marriage Example

```go
import (
  "github.com/abhchand/libmatch"
)

prefTableA := []libmatch.MatchPreference{
  {Name: "A", Preferences: []string{"E", "F", "G", "H"}},
  {Name: "B", Preferences: []string{"F", "G", "H", "E"}},
  {Name: "C", Preferences: []string{"G", "F", "H", "E"}},
  {Name: "D", Preferences: []string{"H", "E", "G", "F"}},
}

prefTableB := []libmatch.MatchPreference{
  {Name: "E", Preferences: []string{"A", "B", "C", "D"}},
  {Name: "F", Preferences: []string{"C", "B", "D", "A"}},
  {Name: "G", Preferences: []string{"D", "A", "C", "B"}},
  {Name: "H", Preferences: []string{"B", "C", "D", "A"}},
}

result, err := libmatch.SolveSMP(&prefTableA, &prefTableB)
if err != nil {
  fmt.Println(err)
  os.Exit(1)
}

// => MatchResult{
//   Mapping: map[string]string{
//     "H": "D",
//     "B": "F",
//     "C": "G",
//     "D": "H",
//     "A": "E",
//     "E": "A",
//     "F": "B",
//     "G": "C",
//   }
// }
```

#### <a name="pkg-stable-roommates-example">Stable Roommates Example

```go
import (
  "github.com/abhchand/libmatch"
)

prefTable := []libmatch.MatchPreference{
  {Name: "A", Preferences: []string{"B", "C", "D"}},
  {Name: "B", Preferences: []string{"A", "C", "D"}},
  {Name: "C", Preferences: []string{"A", "B", "D"}},
  {Name: "D", Preferences: []string{"A", "B", "C"}}
}

result, err := libmatch.SolveSRP(&prefTable)
if err != nil {
  fmt.Println(err)
  os.Exit(1)
}

// => MatchResult{
//   Mapping: map[string]string{
//     "D": "C",
//     "A": "B",
//     "B": "A",
//     "C": "D",
//   }
// }
```

## <a name="cli">CLI

### <a name="cli-installation"></a>Installation

Download the the [latest release](https://github.com/abhchand/libmatch/releases/latest) for your platform.

Or alternatively, build it from source:

```shell
git clone git@github.com:abhchand/libmatch.git
cd libmatch/

make all
```


### <a name="cli-examples">Examples

#### <a name="cli-stable-marriage-example">Stable Marriage Example
```shell
$ cat <<EOF > prefs-a.json
[
  { "name": "A", "preferences": ["E", "F", "G", "H"] },
  { "name": "B", "preferences": ["F", "G", "H", "E"] },
  { "name": "C", "preferences": ["G", "F", "H", "E"] },
  { "name": "D", "preferences": ["H", "E", "G", "F"] }
]
EOF

$ cat <<EOF > prefs-b.json
[
  { "name": "E", "preferences": ["A", "B", "C", "D"] },
  { "name": "F", "preferences": ["C", "B", "D", "A"] },
  { "name": "G", "preferences": ["D", "A", "C", "B"] },
  { "name": "H", "preferences": ["B", "C", "D", "A"] }
]
EOF

$ libmatch solve --algorithm SMP --file prefs-a.json --file prefs-b.json
H,D
B,F
C,G
D,H
A,E
E,A
F,B
G,C
```

#### <a name="cli-stable-roommates-example">Stable Roommates Example

```shell
$ cat <<EOF > prefs.json
[
  { "name": "A", "preferences": ["B", "C", "D"] },
  { "name": "B", "preferences": ["A", "C", "D"] },
  { "name": "C", "preferences": ["A", "B", "D"] },
  { "name": "D", "preferences": ["A", "B", "C"] }
]
EOF

$ libmatch solve --algorithm SRP --file prefs.json
D,C
A,B
B,A
C,D
```

## <a name="miscellaneous">Miscellaneous

* [Create an issue](https://github.com/abhchand/libmatch/issues/new) to report a bug or request a feature
* Contributions are welcome! Please [open an Issue](https://github.com/abhchand/libmatch/issues/new) to discuss your changes first
* The Changelog can be found in the [releases](https://github.com/abhchand/libmatch/releases)

[ci-badge]:
  https://github.com/abhchand/libmatch/actions/workflows/test.yml/badge.svg?branch=main
[ci]:
  https://github.com/abhchand/libmatch/actions
