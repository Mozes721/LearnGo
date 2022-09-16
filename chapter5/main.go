package main

import (
	"errors"
	"io"
	"log"
	"os"
)

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("canno divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func add(i, j int) int { return i + j }

func main() {
	// result, remainder, err := divAndRemainder(5, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(result, remainder)
	// fmt.Println(add(2, 12))

	if len(os.Args) < 2 {
		log.Fatal("no file specified")

	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
