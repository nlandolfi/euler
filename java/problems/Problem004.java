/*
Problem 4
---------
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package problems;

public class Problem004 {
    public static int Solution() {
        int largest =-1;
        for (int k = 100; k < 1000; k++) {
            for (int c = 100; c < 1000; c++) {
                int p = c*k;
                if (isPalindrome(p) && p > largest) {
                    largest = p;
                }
            }
        }

        return largest;
    }

    public static boolean isPalindrome(int n) {
        String ns = Integer.toString(n);
        int i = 0;
        int j = ns.length()-1;
        while (i < j) {
            if (ns.charAt(i) != ns.charAt(j)) {
                return false;
            }

            i++;
            j--;
        }
        return true;
    }
}
