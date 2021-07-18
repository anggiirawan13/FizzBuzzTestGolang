package main

import "fmt"

func main() {
	FizzBuzz(50)
}

func FizzBuzz(value int) {
	for i := 1; i <= value; i++ {
		if i%2 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%2 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
