package main

import "fmt"

func main() {
	matrix := [3][3]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	//fmt.Println(matrix)
	var primaryDiagonalSum, secondaryDiagonalSum, Difference int32
	n := len(matrix)
	fmt.Println("n Is", n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				primaryDiagonalSum += matrix[i][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i + j) == (n - 1) {
				secondaryDiagonalSum += matrix[i][j]
			}
		}
	}

	fmt.Println(primaryDiagonalSum, secondaryDiagonalSum)
	Difference = primaryDiagonalSum - secondaryDiagonalSum
	if Difference < 0 {
		fmt.Println(-(Difference))
	} else {
		fmt.Println(Difference)
	}

}
