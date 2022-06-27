package main

import "testing"

func PrimeNumberGenerator(limit int) []int {
	var buffer []int = make([]int, limit)
	var primes []int = make([]int, 0)

	for index, item := range buffer {
		if index < 2 {
			continue
		}
		if item == 1 {
			continue
		} else {
			for i := index * 2; i < limit; i += index {
				buffer[i] = 1
			}
		}
	}

	for i, val := range buffer {
		if val == 0 {
			primes = append(primes, i)
		}
	}

	return primes

}

func IsPrime(number int) bool {
	primes := PrimeNumberGenerator(number)
	for _, n := range primes {
		if n == number {
			return true
		}
	}
	return false
}

var num int = 1000

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeNumberGenerator(num)
	}
}
