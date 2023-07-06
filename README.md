# types

[![Static Badge](https://img.shields.io/badge/project%20use%20codesystem-green?link=https%3A%2F%2Fgithub.com%2Fgofast-pkg%2Fcodesystem)](https://github.com/gofast-pkg/codesystem)
![Build](https://github.com/gofast-pkg/types/actions/workflows/ci.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/gofast-pkg/types.svg)](https://pkg.go.dev/github.com/gofast-pkg/types)
[![codecov](https://codecov.io/gh/gofast-pkg/types/branch/main/graph/badge.svg?token=7TCE3QB21E)](https://codecov.io/gh/gofast-pkg/types)
[![Release](https://img.shields.io/github/release/gofast-pkg/types?style=flat-square)](https://github.com/gofast-pkg/types/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofast-pkg/types)](https://goreportcard.com/report/github.com/gofast-pkg/types)
[![codebeat badge](https://codebeat.co/badges/342befce-c894-4cff-9fca-acdecb1d6c65)](https://codebeat.co/projects/github-com-gofast-pkg-types-main)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/gofast-pkg/types/blob/main/LICENSE)

Package types group common operation on standard type in Go

## Install

``` bash
$> go get github.com/gofast-pkg/types@latest
```

## Usage

``` Golang
import github.com/gofast-pkg/types

func main() {
  s1 := []int32{15, 22, 8, 74, 23}
  x1 := int32(42)
  s2 := []string{"15", "22", "8", "74", "23"}
  x2 := "22"

  ok := SliceFinder(s1, x1)
  // false
  ok = SliceFinder(s2, x2)
  // true
}
```

Check the [go documentation](https://pkg.go.dev/github.com/gofast-pkg/types) for more details.

## Contributing

&nbsp;:grey_exclamation:&nbsp; Use issues for everything

Read more informations with the [CONTRIBUTING_GUIDE](./.github/CONTRIBUTING.md)

For all changes, please update the CHANGELOG.txt file by replacing the existant content.

Thank you &nbsp;:pray:&nbsp;&nbsp;:+1:&nbsp;

<a href="https://github.com/gofast-pkg/types/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gofast-pkg/types" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## Licence

[MIT](https://github.com/gofast-pkg/types/blob/main/LICENSE)
