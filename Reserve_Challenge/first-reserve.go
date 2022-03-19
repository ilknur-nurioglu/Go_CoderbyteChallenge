/* Have the function FirstReserve(str) take the str parameter being passed and return the string in reserved order. 
For example: if the input string is "Hello World and Coders" then your program should return the string "sredoC dna dlroW olleH" */

package main

import "fmt"

func main() {
	fmt.Println(FirstReserve("Hello World and Coders"))
}

func FirstReserve(str string) string {
	ans := ""
	for _, c := range str {
		ans = string(c) + ans
	}
	return ans
}
