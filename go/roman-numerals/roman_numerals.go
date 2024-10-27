package romannumerals

import (
	"errors"
	"fmt"
	"strconv"
)

var rns = []string{"I", "V", "X", "L", "C", "D", "M", "", ""}

// ToRomanNumeral convets an arabic number
// to roman numeral.
func ToRomanNumeral(arabic int) (roman string, err error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("arabic num not in range of 1-3000")
	}

	arabicRunes := []rune(fmt.Sprintf("%d", arabic))

	for len(arabicRunes) > 0 {
		num, _ := strconv.Atoi(string(arabicRunes[0]))
		place := len(arabicRunes)
		arabicRunes = arabicRunes[1:]

		if num == 0 {
			continue
		}

		roman += toRoman(num, place*2)
	}

	return roman, nil
}

func toRoman(num int, place int) (roman string) {
	start, half, next := rns[place-2], rns[place-1], rns[place]
	switch num {
	case 4:
		roman = start + half
	case 9:
		roman = start + next
	default:
		if num > 3 {
			roman = half
			num -= 5
		}
		for i := 0; i < num; i++ {
			roman += start
		}
	}

	return roman
}
