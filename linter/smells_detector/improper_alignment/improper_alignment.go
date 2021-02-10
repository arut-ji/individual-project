package improper_alignment

import (
	"github.com/instrumenta/kubeval/kubeval"
)

func GetNumberOfInstances(script string) (int, error) {
	return countSmellInstances(script)
}

func countSmellInstances(script string) (int, error) {
	result, err := kubeval.Validate([]byte(script), kubeval.NewDefaultConfig())
	if err != nil {
		return 1, nil
	}
	for _, v := range result {
		if len(v.Errors) != 0 {
			return 1, nil
		}
	}
	return 0, nil
}
