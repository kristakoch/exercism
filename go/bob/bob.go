// Package bob generates responses from bob.
package bob

import (
	"strings"
)

// Hey returns bob's response given an input phrase.
func Hey(remark string) string {

	// Clean the string by removing whitespace.
	remark = strings.Trim(remark, "\n\r\t ")

	// Repond, based on input.
	switch {
	case isAsking(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"

	case isYelling(remark):
		return "Whoa, chill out!"

	case isAsking(remark):
		return "Sure."
	
	case isSilent(remark):
		return "Fine. Be that way!"
		
	default:
		return "Whatever."
	}
}

// The functions below take in string and return the 
// type of remark for the input.
func isYelling(remark string) bool {
	return remark == strings.ToUpper(remark) && 
	remark != strings.ToLower(remark)
}
func isAsking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
func isSilent(remark string) bool {
	return remark == ""
}