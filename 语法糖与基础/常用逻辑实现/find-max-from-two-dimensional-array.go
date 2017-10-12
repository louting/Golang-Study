package main

import "fmt"

func main() {
	var row, col, max int

	var a = [3][4]int{{1, 3, 7, 3}, {2, 3, 7, 9}, {22, 3, 5, 10}}

	max = a[0][0]

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 3; j++ {
			if a[i][j] > max {
				max = a[i][j]
				row = i
				col = j
			}
		}
	}

	fmt.Println("max = %d, row = %d, col = %d\n", max, row, col)
}
