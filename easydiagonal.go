package kata

import "fmt"

/* # Easy diagonal

In the drawing below we have a part of the Pascal's triangle, lines are numbered from zero (top).
The left diagonal in pale blue with only numbers equal to 1 is diagonal zero, then in dark green (1, 2, 3, 4, 5, 6, 7)
is diagonal 1, then in pale green (1, 3, 6, 10, 15, 21) is diagonal 2 and so on.

			 ,-- diagonal zero
			1  ,-- diagonal one
		      1   1  ,-- diagonal two
		   1    2   1
		 1   3    3   1
	      1    4    6    4   1
	    1   5    10   10   5   1
	  1   6   15   20   15   6    1
	1   7   21   35   35   21   7   1

We want to calculate the sum of the binomial coefficients on a given diagonal. The sum on diagonal 0 is 8
(we'll write it S(7, 0), 7 is the number of the line where we start, 0 is the number of the diagonal).
In the same way S(7, 1) is 28, S(7, 2) is 56.

Can you write a program which calculate S(n, p) where n is the line where we start and p is the number of the diagonal?

The function will take n and p (with: n >= p >= 0) as parameters and will return the sum.

## Examples:

diagonal(20, 3) => 5985
diagonal(20, 4) => 20349

## Hint
When following a diagonal from top to bottom have a look at the numbers on the diagonal at its right.

## Ref
http://mathworld.wolfram.com/BinomialCoefficient.html
*/

func Diagonal(n, p int) int {
	return SumDiagonal(PascalData(n, p), p)
}

/* Provided with a depth and a diagonal, it is possible to locate the element in the diagonal from which the sum
computation will begin. For instance, with depth 5 and diagonal 3, the element would be located as follows:

			     1
  Begin  	          1      1
   here		      1       2       1   ,-- diagonal
    |		  1       3       3      1
    v	   ,-|1|-. .-|4|-. .-|6|-.   4      1
  depth-> |1|    |5|     |10|   *10*     5     1
	 1    6      15       20     15     6    1
      1    7     21      35       35     21    7    1

Note that locating this element is only a function of depth and diagonal, and not of the actual values of the triangle.
This means that the pivot can be calculated before even constructing the triangle.
This element is denoted as pivot and allows to reduce the scope of the sum computation on the pascal triangle to the
bare minimum. For the example above, the minimum required data from the triangle would be:

			1  ,-- diagonal one
		      1   1  ,-- diagonal two
		   1    2   1
		     3    3    1
	                6    4
	                 |10|

 */
func findPivot(depth, diagonal int) (pivoti, pivotj int) {
	pivoti = depth
	pivotj = 0
	for pivotj != diagonal {
		pivoti--
		pivotj++
	}
	return
}

/*
Computes the minimum amount of pascal data that is sufficient to calculate the requested sum S(n,p)
 */
func PascalData(depth int, diagonal int) [][]int {
	pivoti, pivotj := findPivot(depth, diagonal)

	pyramid := make([][]int, pivoti + 1)
	for i := 0; i <= pivoti; i++ {
		pyramid[i] = make([]int, pivotj + 1)
		for j := 0; j <= pivotj; j++ {
			if i == 0 || j == 0 {
				pyramid[i][j] = 1
			} else {
				pyramid[i][j] = pyramid[i][j-1] + pyramid[i-1][j]
			}
		}
	}
	return pyramid
}

func SumDiagonal(pyramid [][]int, diagonal int) int {
	sum := 0
	fmt.Println("Pyramid: ", pyramid)
	for i := 0; i < len(pyramid); i++ {
		sum += pyramid[i][diagonal]
	}

	return sum
}
