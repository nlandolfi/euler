/*
Problem 3
---------
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package problems

func Solution003() int {
	num := 600851475143

	largest := 1
	for k := 1; k < num; k++ {
		if divisible(num, k) && prime(k) && k > largest {
			largest = k
		}
	}
	return largest
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
