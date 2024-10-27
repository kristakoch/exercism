package piglatin

// import (
// 	"regexp"
// 	"strings"
// )

// // Sentence translates a sentence into pig latin.
// func Sentence(input string) string {
// 	words := strings.Split(input, " ")

// 	var translated []string
// 	for _, word := range words {
// 		translated = append(translated, translate(word))
// 	}

// 	return strings.Join(translated, " ")
// }

// // translate translates a word into pig latin.
// func translate(word string) string {

// 	// Handle vowel sound prefixed words.
// 	leadingVwSoundRX := regexp.MustCompile("^[aeiou]|^(xr)|^(yt)")
// 	vprefix := leadingVwSoundRX.FindString(word)

// 	if vprefix != "" {
// 		return word + "ay"
// 	}

// 	// Non-yt y prefixes count as consonants.
// 	if strings.HasPrefix(word, "y") {
// 		return word[1:] + "y" + "ay"
// 	}

// 	// Handle consonant sound prefixed words.
// 	leadingCsSoundRX := regexp.MustCompile("^[^aeiouyAEIOUY]*")
// 	cprefix := leadingCsSoundRX.FindString(word)

// 	translated := word[len(cprefix):] + word[:len(cprefix)]

// 	// u's following q's should be grouped with the q.
// 	if translated[len(translated)-1] == 'q' && translated[0] == 'u' {
// 		translated = translated[1:] + "u"
// 	}

// 	return translated + "ay"
// }
