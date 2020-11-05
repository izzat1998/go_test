package main

import (
	"fmt"
)

func CheckPolindrome(str string) (result string) {
	for _, a := range str {
		result = string(a) + result
	}

	if str == result {
		return "Polindrome"
	} else {
		return "Not Polindrome"
	}
}

func main() {
	var str string

	fmt.Scan(&str)
	println(CheckPolindrome(str))

}
