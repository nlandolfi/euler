/*
Problem 3
---------
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package problems

import "fmt"

func factor(n int, k int, factors []int) []int {
	if k > n {
		return factors
	}

	kk := k + 1

	if n%k == 0 {
		return factor(int(n/k), kk, append(factors, k))
	}

	return factor(int(n/k), kk, factors)
}

func Solution003() int {
	num := 600851475143

	factors := factor(num, 1, make([]int, 0))
	fmt.Print(factors)
	return largest(primes(factors))
}

func primes(ints []int) []int {
	primes := []int{}

	for i := range ints {
		if prime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func prime(n int) bool {
	if n < 1 {
		return false
	}

	for k := 2; k < n; k++ {
		if divisible(n, k) {
			return false
		}
	}

	return true
}

func divisible(n int, k int) bool {
	if n%k == 0 {
		return true
	}
	return false
}

func largest(ints []int) int {
	l := ints[0]

	for _, k := range ints {
		if k > l {
			l = k
		}
	}

	return l
}
