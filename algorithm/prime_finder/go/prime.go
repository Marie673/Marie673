package main

import (
	"math"
	"math/big"
)

type Number struct {
	value   int64
	isPrime bool
}

type NumberBig struct {
	value   *big.Int
	isPrime bool
}

func Generate(max int64) []int64 {
	numbers := make([]*Number, max)
	for n := int64(0); n < max; n++ {
		num := &Number{
			value:   n,
			isPrime: true,
		}
		numbers[n] = num
	}

	limit := int64(math.Sqrt(float64(max)))
	for _, num := range numbers {
		if num.value < 2 || limit < num.value {
			continue
		}

		idx := int64(num.value * 2)
		if num.isPrime {
			for i := int64(2); idx < int64(len(numbers)); i++ {
				numbers[idx].isPrime = false
				idx = int64(num.value) * i
			}
		}
	}

	var primes []int64
	for _, num := range numbers {
		if num.value < 2 {
			continue
		}
		if num.isPrime {
			primes = append(primes, num.value)
		}
	}

	return primes
}

func GenerateBig(max int64) []*big.Int {
	numbersBig := make([]*NumberBig, max)
	for n := int64(0); n < max; n++ {
		ng := &NumberBig{
			value:   big.NewInt(n),
			isPrime: true,
		}
		numbersBig[n] = ng
	}
	two := big.NewInt(2)
	limit := big.NewInt(int64(math.Sqrt(float64(max + 1))))
	for _, ng := range numbersBig {
		if ng.value.Cmp(two) == -1 || ng.value.Cmp(limit) == 1 {
			continue
		}

		if ng.isPrime {
			idx := int64(ng.value.Int64() * 2)
			for i := int64(2); idx < int64(len(numbersBig)); i++ {
				numbersBig[idx].isPrime = false
				idx = ng.value.Int64() * i
			}
		}
	}

	var primes []*big.Int
	for _, ng := range numbersBig {
		if ng.value.Cmp(two) == -1 {
			continue
		}
		if ng.isPrime {
			primes = append(primes, ng.value)
		}
	}

	return primes
}
