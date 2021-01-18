package improper_alignment

import "github.com/instrumenta/kubeval/kubeval"

func Scan(script string) (bool, error) {
	return hasSmell(script)
}

func hasSmell(script string) (bool, error) {
	result, err := kubeval.Validate([]byte(script), kubeval.NewDefaultConfig())
	if err != nil {
		return false, err
	}
	for _, v := range result {
		if len(v.Errors) != 0 {
			return false, nil
		}
	}
	return true, nil
}
