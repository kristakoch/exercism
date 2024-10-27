package palindrome

import (
	"errors"
	"fmt"
	"math"
)

// Product holds data for a palindromic product and a list of all
// possible 2-factor factorizations of Product within rage, in order.
type Product struct {
	Product        int
	Factorizations [][2]int
}

var (
	errFminLargerThanFmax = errors.New("fmin > fmax")
	errNoPalindromes      = errors.New("no palindromes")
)

// Products finds the largest and palindrome
// products within range of beg and end.
func Products(
	fmin int,
	fmax int,
) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, errFminLargerThanFmax
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product, factors := i*j, [2]int{i, j}

			if !isPalindrome(product) {
				continue
			}

			if pmin.Product == 0 || product <= pmin.Product {
				pmin = determineProduct(pmin, product, factors)
			}
			if pmax.Product == 0 || product >= pmax.Product {
				pmax = determineProduct(pmax, product, factors)
			}
		}
	}

	if pmin.Product == 0 && pmax.Product == 0 {
		return pmin, pmax, errNoPalindromes
	}

	return pmin, pmax, nil
}

// determineProduct determines the resulting Product from an original
// and new product, adding to factors if equal and replacing if not.
func determineProduct(
	originalProduct Product,
	newProductNum int,
	factors [2]int,
) Product {
	if newProductNum == originalProduct.Product {
		originalProduct.Factorizations = append(originalProduct.Factorizations, factors)
	} else {
		originalProduct.Factorizations = [][2]int{factors}
	}

	originalProduct.Product = newProductNum

	return originalProduct
}

// isPalindrome is a helper to determine whether
// a number is a palindrome, ignoring the sign.
func isPalindrome(number int) bool {
	numString := fmt.Sprintf("%.0f", math.Abs(float64(number)))
	return numString == reverse(numString)
}

// reverse returns the input string in reverse.
func reverse(str string) string {
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}
