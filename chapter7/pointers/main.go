package main 

import (
	"fmt"
)

func main() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("in main", c.String())
	doUpdateRight(&c)
	fmt.Println("in main", c.String())
}
