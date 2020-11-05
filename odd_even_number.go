package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	if number%2 == 0 {
		println("%V is Even Number", number)
	} else {
		println("%V is Odd Number", number)
	}

}
