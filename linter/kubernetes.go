package linter

import (
	"github.com/arut-ji/individual-project/util"
	"github.com/instrumenta/kubeval/kubeval"
)

func Lint(content []byte) ([]kubeval.ValidationResult, error) {
	return kubeval.Validate(content, kubeval.NewDefaultConfig())
}

func IsKubernetesScriptValid(content string) (bool, error) {
	decodedContent, err := util.DecodeContent(content)
	if err != nil {
		return false, err
	}
	_, err = Lint(decodedContent)
	return err == nil, nil
}
