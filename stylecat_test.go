package stylecat

import "testing"

// Reference from: https://developer.mozilla.org/en-US/docs/Web/CSS/@import

var cases = []struct {
	// Import is the `@import...` string.
	Import string

	// Expected is the extracted path value.
	Expected string

	// Match is true when a `@import...;` string is found.
	Match bool

	// Valid is true for acceptably formatted `@import...` statements.
	Valid bool
}{
	{`@import url("fineprint.css") print;`, ``, true, false},
	{`@import url("bluish.css") speech;`, ``, true, false},
	{`@import url("bluish.css");`, `bluish.css`, true, true},
	{`@import url("/css/bluish.css");`, `/css/bluish.css`, true, true},
	{`@import url('/css/bluish.css');`, `/css/bluish.css`, true, true},
	{`@import url(/css/bluish.css);`, ``, true, true},
	{`@import 'custom.css';`, `custom.css`, true, true},
	{`@import '/css/custom.css';`, `/css/custom.css`, true, true},
	{`@import '../custom.css';`, `../custom.css`, true, true},
	{`@import url("chrome://communicator/skin/");`, `chrome://communicator/skin/`, true, false},
	{`@import "common.css" screen;`, ``, true, false},
	{`@import url('landscape.css') screen and (orientation:landscape);`, ``, true, false},
	{`@import 'invalid;`, ``, true, false},
	{`@import url('https://fonts.googleapis.com/css?family=Roboto');`, `https://fonts.googleapis.com/css?family=Roboto`, true, false},
	{`@import url(https://fonts.googleapis.com/css?family=Roboto);`, ``, true, false},
}

func TestImportRegex(t *testing.T) {
	rgx, err := getImportRegex()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		if rgx.Match([]byte(c.Import)) != c.Match {
			t.Errorf("Expected %s not to %v", c.Import, c.Match)
		}
	}
}

func TestFindImportPath(t *testing.T) {
	rgx, err := getPathRegex()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		result := findImportPath([]byte(c.Import), rgx)
		if result != c.Expected {
			t.Errorf("Expected (%s) got (%s) for (%s)", c.Expected, result, c.Import)
		}
	}
}
