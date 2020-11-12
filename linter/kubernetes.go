package linter

import (
	"github.com/arut-ji/individual-project/util"
	"github.com/instrumenta/kubeval/kubeval"
)

func Lint(content []byte) ([]kubeval.ValidationResult, error) {
	return kubeval.Validate(content, kubeval.NewDefaultConfig())
}

func IsKubernetesScriptValid(content string) bool {
	decodedContent, err := util.DecodeContent(content)
	if err != nil {
		panic(err)
	}
	_, err = Lint(decodedContent)
	return err == nil
}
