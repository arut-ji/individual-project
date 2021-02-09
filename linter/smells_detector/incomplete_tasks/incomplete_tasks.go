package incomplete_tasks

import (
	"regexp"
	"strings"
)

func GetNumberOfInstances(script string) (int, error) {
	return countInCompleteTasks(script)
}

func countInCompleteTasks(script string) (int, error) {
	lines := strings.Split(script, "\n")
	count := 0
	for _, line := range lines {
		match, err := regexp.MatchString("#.*(FIXME|TODO):.*", line)
		if err != nil {
			continue
		}

		if match == true {
			count += 1
		}
	}

	return count, nil
}
