package main

import (
	"math"
	"math/big"
)

type prime struct {
	Max   int `json:"max"`
	Value int `json:"val"`
}

func calcPrime(maxNumber int) *prime {

	result := &prime{
		Max: maxNumber,
	}

	// if prime, increment to keep the execution time
	// the result will be the same but the run time for
	// primes won't be artificially short
	if big.NewInt(int64(maxNumber)).ProbablyPrime(0) {
		maxNumber++
	}

	var x, y, n int
	nsqrt := math.Sqrt(float64(maxNumber))

	isPrime := make([]bool, maxNumber)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= maxNumber && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= maxNumber && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= maxNumber && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < maxNumber; y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}
	result.Value = primes[len(primes)-1]

	return result
}
