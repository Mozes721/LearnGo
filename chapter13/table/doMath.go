package table

import (
	"errors"
	"fmt"
	"testing"
)

//data := []struct {
//	name string
//	num1 int
//	num2 int
//	op string
//	expected int
//	errMsg string
//}{
//	{"addition", 2,2, "+",4,""},
//	{"subtraction",2,2, "-",0,""},
//	{"mulitiplication",2,2, "+",4,""},
//	{"divition",2,2, "/",1,""},
//	{"bad_divition",2,0, "/",0,"divition by zero"},
//}
//
//for _, d := range data {
//	t.Run(d.name, func(t *testing.T) {
//		result, err := doMath(d.num1,d,num2,d.op)
//		if result != d.expected {
//			t.Errorf("Expected %d, got %d", d.expected, result)
//		}
//		var errMsg string
//		if err != nil {
//			errMsg = err.Error()
//		}
//		if errMsg != d.errMsg {
//			t.Errorf("Expected %s, got %s", d.errMsg	)
//}
//	})
//}

func doMath(num1, num2 int, op string) (int, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator %s", op)
	}
}
func TestDoMath(t *testing.T) {
	result, err := doMath(2, 2, "+")
	if result != 4 {
		t.Error("Should have been 4, got", result)
	}
	if err != nil {
		t.Error("Should have been nil, got", err)
	}
	result2, err2 := doMath(2, 2, "-")
	if result2 != 0 {
		t.Error("Should have been 0, got", err2)
	}
	if err2 != nil {
		t.Error("Should have been nil, got", err2)
	}
}
