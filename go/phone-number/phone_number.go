package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	area       = "[2-9]([0-9]{2})"
	exchange   = "[2-9]([0-9]{2})"
	subscriber = "[0-9]{4}"
	valid      = "^1?" + area + exchange + subscriber + "$"
)

// Number returns an unformatted number for valid inputs.
func Number(input string) (string, error) {
	var err error

	// Remove punctuation.
	var rxrempunc *regexp.Regexp
	if rxrempunc, err = regexp.Compile("[^0-9]"); nil != err {
		return "", err
	}
	input = rxrempunc.ReplaceAllString(input, "")

	// Check to see if the number is valid.
	var rxvalidnum *regexp.Regexp
	if rxvalidnum, err = regexp.Compile(valid); nil != err {
		return "", err
	}
	if valid := rxvalidnum.MatchString(input); !valid {
		return "", fmt.Errorf("input '%s' is not a valid number", input)
	}

	return strings.TrimPrefix(input, "1"), nil
}

// AreaCode returns the area code for valid numbers.
func AreaCode(input string) (string, error) {
	var err error

	var num string
	if num, err = Number(input); nil != err {
		return "", err
	}

	return num[:3], nil
}

// Format returns the fomatted string for valid numbers.
func Format(input string) (string, error) {
	var err error

	var num string
	if num, err = Number(input); nil != err {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}
