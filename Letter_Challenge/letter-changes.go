/* Have the function LetterChanges(str) take the str parameter being passesd and modify it using the following algorithm. Replace every letter in the string with the letter following it in the alphabet (ie. c becomes d, z becomes a). Then capitalize every vowel in this new string (a,e,i,o,u) and finally return this modified string. */

package main

import (
	"fmt"
	"strings"
)

func LetterChanges(txt string) string {
	newStr := []string{}
	for i := 0; i < len(txt); i++ {
		b := txt[i]

		if b >= 97 && b <= 122 {
			b = b + 1

		}
		l := string(b)

		switch {
		case l == "a":
			fallthrough
		case l == "e":
			fallthrough
		case l == "i":
			fallthrough
		case l == "o":
			fallthrough
		case l == "u":
			l = strings.ToUpper(l)
		}
		newStr = append(newStr, string(l))
	}
	result := strings.Join(newStr, "")
	return result
}

func main() {
	fmt.Println(LetterChanges("hello world"))
}
