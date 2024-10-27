package clock

import (
	"fmt"
)

// Clock holds information about times.
type Clock struct {
	hours   int
	minutes int
}

const minutesInADay = 1440

// New creates a new clock.
func New(hours int, minutes int) Clock {

	// Get the total minutes left over after dividing
	// the minutes up into days.
	totalMinutes := (hours*60 + minutes) % minutesInADay

	// Negative minute values are counted backwards
	// from the total number of minutes in a day.
	if totalMinutes < 0 {
		totalMinutes += minutesInADay
	}

	return Clock{totalMinutes / 60, totalMinutes % 60}
}

// String is a stringer method for clocks.
func (c Clock) String() string {
	return fmt.Sprintf("%0.2v:%0.2v", c.hours, c.minutes)
}

// Add returns the clock with minutes added.
func (c Clock) Add(minutes int) Clock {
	return New(c.hours, c.minutes+minutes)
}

// Subtract returns the clock with minutes subtracted.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hours, c.minutes-minutes)
}
