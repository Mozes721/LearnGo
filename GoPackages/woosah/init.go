package woosah

import "fmt"

// Private visibility
var test1 = "private variable"

// Public visibility
var Test2 = "public variable"

// Private visiblity
func greet() string {
	return "hello"
}

// Public visibility
func Greet2() string {
	return "hello world"
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Buzz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
