package main

import (
	"fmt"
)

var xq1, xq2, xq3 int

func main() {
	var x1, x2, x3, x4 = 1, 2, 3, 4
	var q1, q2, q3, q4 int
	q1, q2, q3, q4 = x1, x2, x3, x4
	fmt.Printf("%v%v%v%d\n", x1, x2, x3, x4)
	fmt.Printf("%v%v%v%d\n", q1, q2, q3, q4)

	fmt.Println("hello world!!" + "every day\n")
	fmt.Print("-------------画个爱心玩玩---------\n")
	test(5, 20, 2)
	test(11, 20, -1)
}
func threeCount(souce bool, a int, b int) int {
	if souce {
		return a
	} else {
		return b
	}
	return a
}

func test(row int, col int, is int) {
	if is == 1 {
		for i := 0; i < row; i++ {
			let := threeCount(col%2 == 0, col+1, col)
			for j := 0; j < let; j++ {
				if j >= let/2-i && j <= let/2+i {
					fmt.Print("○")
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Print("\n")
		}
	} else if is == 2 {
		for i := 0; i < row; i++ {
			let := threeCount(col%2 == 0, col+1, col)
			for j := 0; j < let; j++ {
				if ( j >= let/2-(let/2/2)-i && j <= let/2-(let/2/2)+i) || (j >= let/2+(let/2/2)-i && j <= let/2+(let/2/2)+i) {
					fmt.Print("○")
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Print("\n")
		}
	} else if is == -1 {
		for i := row - 1; i >= 0; i-- {
			let := threeCount(col%2 == 0, col+1, col)
			for j := let - 1; j >= 0; j-- {
				if j >= let/2-i && j <= let/2+i {
					fmt.Print("○")
				} else {
					fmt.Print("  ")
				}
			}
			fmt.Print("\n")
		}
	}

}
