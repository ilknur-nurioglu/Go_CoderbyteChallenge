/*Have the function FirstFactorial(num) take the num parameter being passed and return the factorial of it. 
For example: if num=4, then your program should return (4 * 3 * 2 * 1 = 24).*/

package main

import "fmt"

func main() {
	fmt.Println(FirstFactorial(4))
}

func FirstFactorial(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * FirstFactorial(num-1)
	}
}
