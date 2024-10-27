package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a key greater than 1 and less than p.
func PrivateKey(p *big.Int) *big.Int {
	pk := new(big.Int)

	if p == nil {
		return pk
	}

	min := big.NewInt(2)
	max := big.NewInt(0).Sub(p, min)

	// rand.Int ranges from 0 -> max, so get the range by subtracting
	// min from max, then adding min to the result to make it in range.
	pk, err := rand.Int(rand.Reader, max)
	if err != nil {
		return pk
	}

	return pk.Add(pk, min)
}

// PublicKey calculates a public key using the formula g**a % p.
func PublicKey(a, p *big.Int, g int64) *big.Int {
	pk := new(big.Int)

	if a == nil || p == nil {
		return pk
	}

	return pk.Exp(big.NewInt(g), a, p)
}

// SecretKey calculates a secret key using the formula B**a % p.
func SecretKey(a, B, p *big.Int) *big.Int {
	sk := new(big.Int)

	if a == nil || B == nil || p == nil {
		return sk
	}

	return sk.Exp(B, a, p)
}

// NewPair returns a new pair of keys, private and public.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private, public := new(big.Int), new(big.Int)

	if p == nil {
		return private, public
	}

	private = PrivateKey(p)
	public = PublicKey(private, p, g)

	return private, public
}
