package improper_alignment

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/instrumenta/kubeval/kubeval"
)

func GetNumberOfInstances(script string) (int, error) {
	return countSmellInstances(script)
}

func countSmellInstances(script string) (int, error) {
	result, err := kubeval.Validate([]byte(script), kubeval.NewDefaultConfig())
	if err != nil {
		spew.Dump(err)
		return -1, err
	}
	for _, v := range result {
		if len(v.Errors) != 0 {
			return 1, nil
		}
	}
	return 0, nil
}
