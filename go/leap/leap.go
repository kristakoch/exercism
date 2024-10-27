package leap

// IsLeapYear reports if the year is a leap year.
func IsLeapYear(yr int) bool {
	if yr%4 != 0 {
		return false
	}

	if yr%100 == 0 && yr%400 != 0 {
		return false
	}

	return true
}
