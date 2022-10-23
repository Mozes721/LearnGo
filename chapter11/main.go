package main

import (
	"fmt"
	"learnGo/chapter11/readers"
	"strings"
)

func main() {
	s := "the quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := readers.CountLetters(sr)
	if err != nil {
		return
	}
	fmt.Println(counts)
}
