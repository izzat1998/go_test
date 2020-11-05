package main

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")

		}
		if i%3 == 0 {
			println("Fizz")
		}
		if i%5 == 0 {
			println("Buzz")
		}

	}

}
