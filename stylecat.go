package stylecat

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

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
	return string(val)
}

func Run(entryPath string) ([]byte, error) {
	src, err := ioutil.ReadFile(entryPath)
	if err != nil {
		return nil, err
	}

	importPattern, err := getImportRegex()
	if err != nil {
		return nil, err
	}

	concat := importPattern.ReplaceAllFunc(src, func(b []byte) []byte {
		rgx, err := getPathRegex()
		if err != nil {
			return b
		}

		val := findImportPath(b, rgx)
		if val == "" {
			return b
		}

		// Clean up
		val = strings.ToLower(val)
		val = strings.TrimSpace(val)
		val = strings.Replace(val, "\"", "", -1)
		val = strings.Replace(val, "'", "", -1)

		// Skip URLs
		if strings.Contains(val, "://") {
			return b
		}

		wd, err := os.Getwd()
		if err != nil {
			return b
		}

		p := path.Join(wd, val)
		if !path.IsAbs(p) {
			var err error
			if p, err = filepath.Abs(p); err != nil {
				return b
			}
		}
		result, err := Run(p)
		if err != nil {
			return b
		}
		return result
	})

	return concat, nil
}
