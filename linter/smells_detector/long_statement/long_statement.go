package long_statement

import "strings"

const (
	CharactersPerLine = 100
)

func Scan(script string) (bool, error) {
	return hasLongStatement(script)
}

/*
	Associated with Long Statement smells
	- The script may contain one or more very long statement.
	- The limit is set to 100 characters per line.
	- However, to have such characters per line is very rare.
*/

func hasLongStatement(script string) (bool, error) {
	lines := strings.Split(script, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > CharactersPerLine {
			return true, nil
		}
	}
	return false, nil
}
