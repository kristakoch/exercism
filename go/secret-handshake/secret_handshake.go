package secret

// Handshake returns the secret handshake
// based on the input number.
func Handshake(num uint) []string {
	var handshake []string

	if num&1 == 1 {
		handshake = append(handshake, "wink")
	}
	if num&2 == 2 {
		handshake = append(handshake, "double blink")
	}
	if num&4 == 4 {
		handshake = append(handshake, "close your eyes")
	}
	if num&8 == 8 {
		handshake = append(handshake, "jump")
	}
	if num&16 == 16 {
		handshake = reverse(handshake)
	}

	return handshake
}

// reverse returns the reversal of a string slice.
func reverse(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}
