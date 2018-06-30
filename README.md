# stylecat :cat2:

A Go package to concatenate CSS loaded by `@import` into a single stylesheet.

**Usage**

```go
import (
  "github.com/ef2k/stylecat"
)

func main() {
  src, err := stylecat.Run("/path/to/master.css")
}
```

## Contributions

Fixes? Ideas? Improvements? Jump in. All are welcome!

## License

MIT.
