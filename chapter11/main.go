package main

import (
	"fmt"
	"learnGo/chapter11/readers"
)

func main() {
	r, close, err := readers.BuildGZipReader("my_data.txt.gz")
	if err != nil {
		return
	}
	defer close()
	counts, err := readers.CountLetters(r)
	if err != nil {
		return
	}
	fmt.Println(counts)
}

//s := "the quick brown fox jumped over the lazy dog"
//sr := strings.NewReader(s)
//counts, err := readers.CountLetters(sr)
//if err != nil {
//return
//}
//fmt.Println(counts)
