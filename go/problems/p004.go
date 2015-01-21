/*
Problem 4
---------
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package problems

import "strconv"

func Solution004() int {
	largest := -1

	for k := 100; k < 1000; k++ {
		for c := 100; c < 1000; c++ {
			p := c * k
			if palindrome(p) {
				if p > largest {
					largest = p
				}
			}
		}
	}

	return largest
}

func palindrome(x int) bool {
	sx := strconv.Itoa(x)
	for i, j := 0, len(sx)-1; i < j; i, j = i+1, j-1 {
		if sx[i] != sx[j] {
			return false
		}
	}
	return true
}
