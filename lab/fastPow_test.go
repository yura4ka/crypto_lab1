package lab

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFastPow(t *testing.T) {
	tests := []struct {
		x, y, m, res *big.Int
	}{
		{Big(5), Big(5), Big(10000), Big(3125)},
		{Big(5), Big(5), Big(13), Big(5)},
		{Big(5), Big(5), Big(18), Big(11)},
		{Big(3), Big(11), Big(127), Big(109)},
		{Big(2), Big(19), Big(18), Big(2)},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v ^ %v mod %v = %v", tt.x, tt.y, tt.m, tt.res), func(t *testing.T) {
			ans := FastPow(tt.x, tt.y, tt.m)
			if !EQ(ans, tt.res) {
				t.Errorf("got %s, want %s", ans, tt.res)
			}
		})
	}
}
