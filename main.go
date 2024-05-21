package main

import (
	"fmt"

	"github.com/yura4ka/crypto_lab1/lab"
)

func handlePow() {
	fmt.Println("Enter a, b, m: ")
	var a, b, m string
	fmt.Scan(&a, &b, &m)
	r := lab.FastPow(lab.FromString(a), lab.FromString(b), lab.FromString(m))
	println("Result: ", r)
}

func handlePrime() {
	fmt.Println("Enter n: ")
	var n string
	fmt.Scan(&n)
	r := lab.MillerRabin(lab.FromString(n))
	println("Result: ", r)
}

func main() {
	fmt.Println("(1) Pow(a, b, m), (2) IsPrime(n)")
	var action int
	fmt.Scan(&action)
	switch action {
	case 1:
		handlePow()
	case 2:
		handlePrime()
	}
}
