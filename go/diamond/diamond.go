package diamond

import (
	"fmt"
	"strings"
)

// Gen generates a diamond based on an input letter.
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", fmt.Errorf("char %v is not in range of A-Z", string(b))
	}

	var rows []string
	var inc byte = 1

	var current byte = 'A'
	for current >= 'A' {
		letter := string(current)
		outerSpace := strings.Repeat(" ", int(b-current))
		innerSpace := strings.Repeat(" ", int(((current-65)*2)-1))

		line := outerSpace + letter + innerSpace + letter + outerSpace
		if current == 'A' {
			line = outerSpace + letter + outerSpace
		}

		rows = append(rows, line)

		if current == b {
			inc = -inc
		}
		current += inc
	}

	return strings.Join(rows, "\n") + "\n", nil
}
