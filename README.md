# stylecat :cat2:

A Go package to concatenate `@import` CSS references into a single stylesheet.

**Usage**

```go
import (
  "github.com/ef2k/stylecat"
)

func main() {
    src, err := stylecat.Run("/the/path/to/public/css/master.css", nil)
}
```

## Contributions

Fixes? Ideas? Improvements? Jump in. All are welcome!

## License

MIT.
