/*
Problem 3
---------
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package problems

func Solution003() int {
	num := 600851475143
	return LargestPrime(num)
}

func LargestPrime(n int) int {
	if prime(n) || n == 1 {
		return n
	}

	var factor int

	for k := 1; k <= n; k++ {
		if divisible(n, k) && prime(k) {
			factor = k
			break
		}
	}

	next := LargestPrime(int(n / factor))

	if next > factor {
		return next
	} else {
		return factor
	}
}

func prime(n int) bool {
	if n <= 1 {
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
