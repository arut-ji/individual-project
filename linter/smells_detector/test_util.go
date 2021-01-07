package smells_detector

import (
	"fmt"
	"io/ioutil"
)

type loadFixtureFunc func(string) string

func LoadFixture(fixtureDir string) loadFixtureFunc {
	return func(filename string) string {
		content, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", fixtureDir, filename))
		if err != nil {
			panic(err)
		}
		return string(content)
	}
}
