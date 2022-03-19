/* Have the function SimpleAdding(num) add up all the numbers from 1 to num. For example: if the input is 4 then your program should return 10 because 1 + 2 + 3 + 4 = 10. */

package main

import "fmt"

func main() {
	fmt.Println(SimpleAdding(4))
}

func SimpleAdding(num int) int {
	result := 0
	for i := 0; i <= num; i++ {
		result += i
	}
	return result
}
