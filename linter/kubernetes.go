package linter

import "github.com/instrumenta/kubeval/kubeval"

func Lint(content []byte) ([]kubeval.ValidationResult, error) {
	return kubeval.Validate(content, kubeval.NewDefaultConfig())
}
