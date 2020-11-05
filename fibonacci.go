package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 2; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}

}
