package main

import (
	"fmt"
	"strconv"
)

func main() {
	for num := 1; num <= 100; num++ {
		fmt.Println(FizzBuzz(num))
	}
}

func FizzBuzz(num int) string {
	answer := ""
	if num%3 == 0 {
		answer += "Fizz"
	}
	if num%5 == 0 {
		answer += "Buzz"
	}
	if answer != "" {
		return answer
	} else {
		return strconv.Itoa(num)
	}
}
