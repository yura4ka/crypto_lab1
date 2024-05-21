package lab

import (
	"fmt"
	"math/big"
	"testing"
)

func TestLucasSequnce(t *testing.T) {
	tests := []int{0, 1, 3, 10, 33, 109, 360, 1189, 3927, 12970, 42837, 141481, 467280, 1543321, 5097243, 16835050}
	P := Big(3)
	Q := Big(-1)
	m := Big(1e9)
	for i, v := range tests {
		t.Run(fmt.Sprintf("LucasSequnce U(%v)=%v", i, v), func(t *testing.T) {
			U := LucasSequence(Big(int64(i)), P, Q, m)
			if !EQ(U, v) {
				t.Errorf("got %v, want %v", U, v)
			}
		})
	}
}

func TestLucasProbablePrime(t *testing.T) {
	tests := []struct {
		n       *big.Int
		isPrime bool
	}{
		{Big(1), true},
		{Big(2), true},
		{Big(3), true},
		{Big(4), false},
		{Big(5), true},
		{Big(6), false},
		{Big(7), true},
		{Big(9), false},
		{Big(11), true},
		{Big(13), true},
		{Big(15), false},
		{Big(17), true},
		{Big(727), true},
		{Big(6961), true},
		{Big(7247), true},
		{Big(7753), true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("LucasProbablePrime(%v)=%v", tt.n, tt.isPrime), func(t *testing.T) {
			ans := LucasProbablePrime(tt.n)
			if ans != tt.isPrime {
				t.Errorf("got %v, want %v", ans, tt.isPrime)
			}
		})
	}
}

func BailliePSW(n *big.Int) bool {
	return MillerRabin(n) && LucasProbablePrime(n)
}
