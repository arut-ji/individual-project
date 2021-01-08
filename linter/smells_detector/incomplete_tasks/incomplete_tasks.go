package incomplete_tasks

import "regexp"

func Scan(script string) (bool, error) {
	return hasInCompleteTasks(script)
}

func hasInCompleteTasks(script string) (bool, error) {
	match, err := regexp.MatchString("#.*(FIXME|TODO):.*", script)
	if err != nil {
		return false, err
	}
	return match, nil
}
