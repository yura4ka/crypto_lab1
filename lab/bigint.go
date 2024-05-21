package lab

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
)

var ErrWrongType = errors.New("error wrong type: excpected int64 or bigint")

type Int interface {
	big.Int | int
}

func Big(x int64) *big.Int {
	return big.NewInt(x)
}

func toBigInt(n interface{}) *big.Int {
	switch v := n.(type) {
	case int:
		return Big(int64(v))
	case *big.Int:
		return v
	default:
		panic(ErrWrongType)
	}
}

func Cmp(x *big.Int, y interface{}) int {
	return x.Cmp(toBigInt(y))
}

func GE(x *big.Int, y interface{}) bool {
	return Cmp(x, y) == 1
}

func LE(x *big.Int, y interface{}) bool {
	return Cmp(x, y) == -1
}

func LEQ(x *big.Int, y interface{}) bool {
	return Cmp(x, y) <= 0
}

func EQ(x *big.Int, y interface{}) bool {
	return Cmp(x, y) == 0
}

func IsOdd(x *big.Int) bool {
	temp := new(big.Int)
	temp.And(x, Big(1))
	return EQ(temp, 1)
}

func IsEven(x *big.Int) bool {
	return !IsOdd(x)
}

func RandomBigint(a, b interface{}) *big.Int {
	x := toBigInt(a)
	y := toBigInt(b)

	if EQ(x, y) {
		return Big(0).Set(x)
	}
	max := Sub(y, x)
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return r.Add(r, x)
}

func Add(x *big.Int, y interface{}) *big.Int {
	return Big(0).Add(x, toBigInt(y))
}

func Sub(x *big.Int, y interface{}) *big.Int {
	return Big(0).Sub(x, toBigInt(y))
}

func Mul(x *big.Int, y interface{}) *big.Int {
	return Big(0).Mul(x, toBigInt(y))
}

func Div(x *big.Int, y interface{}) *big.Int {
	return Big(0).Div(x, toBigInt(y))
}

func Mod(x *big.Int, y interface{}) *big.Int {
	return Big(0).Mod(x, toBigInt(y))
}

func Sqrt(x *big.Int) *big.Int {
	return Big(0).Sqrt(x)
}

func Inc(x *big.Int) {
	x.Add(x, Big(1))
}

func GCD(x, y *big.Int) *big.Int {
	return Big(0).GCD(nil, nil, x, y)
}

func FromString(s string) *big.Int {
	r, _ := Big(0).SetString(s, 10)
	return r
}

func ToBase64(n *big.Int) string {
	return base64.RawStdEncoding.EncodeToString([]byte(n.String()))
}
