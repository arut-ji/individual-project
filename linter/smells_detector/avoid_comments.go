package smells_detector

import "regexp"

func HasComments(script string) (bool, error) {
	return hasComments(script)
}

func hasComments(script string) (bool, error) {
	match, err := regexp.MatchString("#.*", script)
	if err != nil {
		return false, err
	}
	return match, nil
}
