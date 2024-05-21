package lab

import "math/big"

func LucasSequence(k, P, Q, m *big.Int) *big.Int {
	u0 := Big(0)
	u1 := Big(1)

	if EQ(k, 0) {
		return u0
	}

	if EQ(k, 1) {
		return u1
	}

	for i := Big(1); LE(i, k); Inc(i) {
		u := Sub(Mul(P, u1), Mul(Q, u0))
		u0 = u1
		u1 = u.Mod(u, m)
	}

	return u1
}

func LucasProbablePrime(n *big.Int) bool {
	if LEQ(n, 3) {
		return true
	}

	if IsEven(n) {
		return false
	}

	if s := Sqrt(n); EQ(Mul(s, s), n) {
		return false
	}

	D := Big(5)

	for ; big.Jacobi(n, D) != -1; D.Add(D, Big(2*int64(D.Sign()))).Neg(D) {
		if g := GCD(n, D); !EQ(g, 1) && !EQ(g, n) {
			return false
		}
	}

	P := Big(1)
	Q := Div(Add(D.Neg(D), 1), 4)

	s := 0
	d := Add(n, 1)

	for IsEven(d) {
		d.Rsh(d, 1)
		s++
	}

	U := LucasSequence(d, P, Q, n)
	U.Mod(U, n)

	if EQ(U, 0) {
		return true
	}

	v0 := Big(2)
	v1 := new(big.Int).Set(P)
	prevD := Big(1)

	for range s {
		for i := prevD; LE(i, d); Inc(i) {
			v := Sub(Mul(P, v1), Mul(Q, v0))

			v0 = v1
			v1 = v.Mod(v, n)
		}

		if EQ(v1, 0) {
			return true
		}
		d.Lsh(d, 1)
	}

	return false
}
