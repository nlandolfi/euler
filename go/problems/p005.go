/*
Problem 4
---------

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package problems

func Solution005() int {
	for k := 1; ; k++ {
		if divisibleThrough(k, 20) {
			return k
		}
	}
}

func divisibleThrough(n int, k int) bool {
	for i := 1; i <= k; i++ {
		if n%i != 0 {
			return false
		}
	}

	return true
}
