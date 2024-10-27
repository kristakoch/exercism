// Package twofer communicates about twofers.
package twofer

import "fmt"

// ShareWith communicates a twofor message with someone 
// named or anonymous. 
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
