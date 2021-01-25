package util

import (
	"fmt"
	"io/ioutil"
)

type LoadFixtureFunc func(string) string

func LoadFixture(fixtureDir string) LoadFixtureFunc {
	return func(filename string) string {
		content, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", fixtureDir, filename))
		if err != nil {
			panic(err)
		}
		return string(content)
	}
}
