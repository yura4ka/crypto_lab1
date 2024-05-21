package lab

import (
	"math/big"
)

// 4^-k
func MillerRabin(n *big.Int, k ...int) bool {
	rounds := 10
	if len(k) != 0 {
		rounds = k[0]
	}

	if LE(n, 3) {
		return true
	}

	if IsEven(n) {
		return false
	}

	s := 0
	d := Sub(n, 1)
	for IsEven(d) {
		d.Rsh(d, 1)
		s++
	}

	n1 := Sub(n, 1)
	for range rounds {
		a := RandomBigint(2, n1)
		x := FastPow(a, d, n)
		for range s {
			y := Mul(x, x)
			y.Mod(y, n)
			if EQ(y, 1) && !EQ(x, 1) && !EQ(x, n1) {
				return false
			}
			x = y
		}
		if !EQ(x, 1) {
			return false
		}
	}

	return true
}
