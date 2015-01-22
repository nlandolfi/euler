/*
Problem 3
---------
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package problems;

public class Problem003  {
    public static int Solution(){
        double x = 600851475143.0;
        return (int) (highestPrime(x));
    }

    public static double highestPrime(double n) {
        if (isPrime(n) || n == 1) {
            return n;
        }

        double current = 1;

        for (double k = 1; k < n; k++) {
            if (n % k == 0 && isPrime(k)) {
                current = k;
                break;
            }
        }

        double next = highestPrime(n / current);

        return Math.max(current, next);
    }

    public static boolean isPrime(double n) {
        if (n < 2) {
            return false;
        }

        for (int k = 2; k < n; k++) {
            if (n%k == 0) {
                return false;
            }
        }

        return true;
    }
}
