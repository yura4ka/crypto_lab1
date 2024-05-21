package lab

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMillerRabin(t *testing.T) {
	big1, _ := Big(0).SetString("5210644015679228794060694325390955853335898483908056458352183851018372555735221", 10)
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
		{Big(28871271685163), true},
		{big1, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("millerRabin(%v)=%v", tt.n, tt.isPrime), func(t *testing.T) {
			ans := MillerRabin(tt.n)
			if ans != tt.isPrime {
				t.Errorf("got %v, want %v", ans, tt.isPrime)
			}
		})
	}
}
