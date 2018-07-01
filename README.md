`stylecat` [![Go Report Card](https://goreportcard.com/badge/github.com/ef2k/stylecat)](https://goreportcard.com/report/github.com/ef2k/stylecat) [![Build Status](https://travis-ci.org/ef2k/stylecat.svg?branch=master)](https://travis-ci.org/ef2k/stylecat) [![GoDoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ef2k/stylecat)
=======
Concatenate `@import` CSS references into a single stylesheet.

## Usage

```go
import (
  "github.com/ef2k/stylecat"
)

func main() {
    src, err := stylecat.Run("/the/path/to/public/css/master.css", nil)
}
```

When stylesheets link to **absolute paths** `e.g: /css/master.css`, set a `RootPath`:

```go
src, err := stylecat.Run("/the/path/to/public/css/master.css", &stylecat.Config{
  RootPath: "/the/path/to/public"
})
```

## Contributions

Fixes? Ideas? Improvements? Jump in. All are welcome!

## License

MIT.
