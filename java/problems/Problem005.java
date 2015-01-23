/*
Problem 4
---------

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package problems;

public class Problem005 {
    public static int Solution() {
        for (int k = 10; ; k++) {
            if (divisibleThrough(k, 20)) {
                return k;
            }
        }
    }

    public static boolean divisibleThrough(int n, int k) {
        for (int i = 2; i <= k; i++) {
            if (n%i != 0) {
                return false;
            }
        }

        return true;
    }
}
