package long_statement

import "strings"

const (
	CharactersPerLine = 100
)

func GetNumberOfInstances(script string) (int, error) {
	return countLongStatement(script)
}

/*
	Associated with Long Statement smells
	- The script may contain one or more very long statement.
	- The limit is set to 100 characters per line.
	- However, to have such characters per line is very rare.
*/

func countLongStatement(script string) (int, error) {
	count := 0
	lines := strings.Split(script, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > CharactersPerLine {
			count += 1
		}
	}
	return count, nil
}
