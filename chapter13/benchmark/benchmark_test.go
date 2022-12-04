package benchmark

import (
	"fmt"
	"os"
	"testing"
)

func FileLen(f string, bufsize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	count := 0

	for {
		buf := make([]byte, bufsize)
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}

	return count, nil
}

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}

	if result != 65204 {
		t.Error("Expected 65204, got", result)
	}
}

var blackhole int

func BenchmarkFileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("testdata/data.txt0=", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testingdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
