/*
Problem 7
---------
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package problems

func Solution007() int {

	cnt := 0

	for k := 1; ; k++ {
		if prime(k) {
			cnt++
			if cnt == 10001 {
				return k
			}
		}
	}

}

// prime function defined in p003.go
