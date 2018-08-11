# Named Regexp Capture

[![CircleCI](https://circleci.com/gh/mdouchement/nregexp.svg?style=shield)](https://circleci.com/gh/mdouchement/nregexp)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mdouchement/nregexp)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdouchement/nregexp)](https://goreportcard.com/report/github.com/mdouchement/nregexp)
[![License](https://img.shields.io/github/license/mdouchement/nregexp.svg)](http://opensource.org/licenses/MIT)

NRegexp is a wrapper around `regexp.Regexp` that allow you to use named captures. It exposes `regexp.Regexp` methods.

## Usage

```go
package main

import (
	"fmt"

	"github.com/mdouchement/nregexp"
)

func main() {
	re := nregexp.MustCompile(`.*\s.*\s(?P<my_capture>[^\s]*)\s.*`)
	fmt.Printf("%v\n", re.NamedCaptureString("Example of nregexp package"))
}
```

## License

**MIT**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

As possible, run the following commands to format and lint the code:

```sh
# Format
find . -name '*.go' -not -path './vendor*' -exec gofmt -s -w {} \;

# Lint
golangci-lint run --enable-all
```
