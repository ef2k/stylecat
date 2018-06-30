package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	src, err := ioutil.ReadFile("css/master.css")
	if err != nil {
		log.Fatal(err)
	}

	statement, _ := regexp.Compile("@import (.+);")

	concat := statement.ReplaceAllFunc(src, func(b []byte) []byte {
		urlRegex, _ := regexp.Compile("url\\((?P<URL>.*?)\\)")
		subs := urlRegex.FindSubmatch(b)
		n := urlRegex.SubexpNames()

		paramsMap := make(map[string]string)
		for i, name := range n {
			if i > 0 && i <= len(subs) {
				paramsMap[name] = string(subs[i])
			}
		}

		// No URL then maybe its a standalone string
		val, ok := paramsMap["URL"]
		if !ok {
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

		// Skip relative paths
		if !path.IsAbs(val) {
			return b
		}

		wd, err := os.Getwd()
		if err != nil {
			return b
		}
		p := path.Join(wd, val)
		css, err := ioutil.ReadFile(p)
		if err != nil {
			return b
		}

		log.Printf("Left with this: %+v | %+v", val, css)
		return css
	})

	log.Printf("The concat: %+v", string(concat))

}
