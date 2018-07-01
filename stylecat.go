package stylecat

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Config configures stylecat.
type Config struct {

	// RootPath will set stylecat to use a fixed absolute path to determine the
	// paths of imported CSS with absolute URL references.
	// For example, `@import url('/css/base.css')` points to a root URL path that
	// has a filesystem path at `/var/www/.../public/css/base.css`. Setting `RootPath`
	// to `/var/www/.../public`, will resolve the correct path for the CSS references.
	// Note that RootPath is only needed for @import paths that are absolute.
	RootPath string
}

func getImportRegex() (*regexp.Regexp, error) {
	return regexp.Compile("@import (.+);")
}

func getPathRegex() (*regexp.Regexp, error) {
	return regexp.Compile(`\(?['"](?P<URL>.+)['"]\)?;`)
}

func findImportPath(s []byte, rgx *regexp.Regexp) string {
	subs := rgx.FindSubmatch(s)
	n := rgx.SubexpNames()

	paramsMap := make(map[string]string)
	for i, name := range n {
		if i > 0 && i <= len(subs) {
			paramsMap[name] = string(subs[i])
		}
	}
	val, ok := paramsMap["URL"]
	if !ok {
		return ""
	}

	// URLs are invalid.
	if strings.Contains(val, "://") {
		return ""
	}

	return string(val)
}

// Run will concatenate all CSS files referenced by @import statements at the
// given `entryPath`.
func Run(entryPath string, c *Config) ([]byte, error) {
	src, err := ioutil.ReadFile(entryPath)
	if err != nil {
		return nil, err
	}

	importRgx, err := getImportRegex()
	if err != nil {
		return nil, err
	}

	pathRgx, err := getPathRegex()
	if err != nil {
		return nil, err
	}

	concat := importRgx.ReplaceAllFunc(src, func(b []byte) []byte {
		val := findImportPath(b, pathRgx)
		if val == "" {
			return b
		}

		base := filepath.Dir(entryPath)
		if c != nil && c.RootPath != "" {
			base = c.RootPath
		}

		p := path.Join(base, val)

		result, err := Run(p, c)
		if err != nil {
			return b
		}
		return result
	})

	return concat, nil
}
