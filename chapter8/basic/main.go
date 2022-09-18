package main

import ("fmt"
		"errors"
		"bytes"
		"archive/zip")

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0{
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil 
}

func doubleEven(i int) (int, error) {
	if i % 2 != 0 {
		return i, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEvenFmt(i int) (int, error) {
	if i % 2 != 0 {
		return i, fmt.Errorf("%d is not an even number", i)
	}
	return i * 2, nil
}

// func main() {
// 	numerator := 20 
// 	denominator := 3 

// 	remainder, mod, err := calcRemainderAndMod(numerator, denominator) 
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(remainder, mod)
// }
func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
