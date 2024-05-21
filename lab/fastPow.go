package lab

import "math/big"

func FastPow(x, y, m *big.Int) *big.Int {
	result := Big(1)
	for GE(y, 0) {
		if IsOdd(y) {
			result.Mul(result, x).Mod(result, m)
		}

		y.Rsh(y, 1)
		x.Mul(x, x).Mod(x, m)
	}

	return result
}
