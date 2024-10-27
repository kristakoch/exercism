package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	registry []string
	used     int
	letters  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

const max = 26 * 26 * 10 * 10 * 10

// Robot holds data about a robot's name.
type Robot struct {
	name string
}

// Name generates the registry of unique names if necessary
// and returns the name of a robot.
func (r *Robot) Name() (string, error) {
	if registry == nil {
		generateRegistry()
	}

	if r.name == "" {
		if used == max {
			return "", errors.New("plum out of unique names")
		}
		r.name = registry[used]
		used++
	}

	return r.name, nil
}

// Reset resets a robot's name to its default
// without recycling the old name.
func (r *Robot) Reset() {
	r.name = ""
}

// generateRegistry generates a randomly ordered registry
// of all possible unique robot names for quick assignment.
func generateRegistry() {
	registry = make([]string, max)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(max)

	var idx int
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters); j++ {
			for k := 0; k < 1000; k++ {
				l1, l2 := string(letters[i]), string(letters[j])
				registry[perm[idx]] = fmt.Sprintf("%s%s%03d", l1, l2, k)
				idx++
			}
		}
	}
}
