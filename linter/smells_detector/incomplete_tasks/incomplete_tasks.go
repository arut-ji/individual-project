package incomplete_tasks

import "regexp"

func HasInCompleteTasks(script string) (bool, error) {
	return inCompleteTasks(script)
}

func inCompleteTasks(script string) (bool, error) {
	match, err := regexp.MatchString("#.*(FIXME|TODO):.*", script)
	if err != nil {
		return false, err
	}
	return match, nil
}
