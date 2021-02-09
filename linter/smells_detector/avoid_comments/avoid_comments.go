package avoid_comments

import (
	"regexp"
	"strings"
)

func GetNumberOfInstances(script string) (int, error) {
	return countComments(script)
}

func countComments(script string) (int, error) {
	lines := strings.Split(script, "\n")
	count := 0
	for _, line := range lines {
		match, err := regexp.MatchString("#.*", line)
		if err != nil {
			continue
		}
		if match == true {
			count += 1
		}
	}
	return count, nil
}
