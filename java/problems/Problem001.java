/*
Problem 1
---------
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package problems;

public class Problem001 extends Problem {

    public static int Solution() {
        int sum = 0;
        for (int k = 1; k < 1000; k++) {
            if (k % 3 == 0 || k % 5 == 0) {
                sum += k;
            }
        }
        return sum;
    }

}
