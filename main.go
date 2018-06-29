package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	src, err := ioutil.ReadFile("css/master.css")
	if err != nil {
		log.Fatal(err)
	}

	statement, _ := regexp.Compile("@import (.+);")

	statement.ReplaceAllFunc(src, func(b []byte) []byte {
		url, _ := regexp.Compile("url\\((?P<URL>.*?)\\)")
		subs := url.FindSubmatch(b)
		n := url.SubexpNames()

		paramsMap := make(map[string]string)
		for i, name := range n {
			if i > 0 && i <= len(subs) {
				paramsMap[name] = string(subs[i])
			}
		}

		log.Printf("%+v", paramsMap)

		// No URL then maybe its a standalone string.

		return b
	})
}
