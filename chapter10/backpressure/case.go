package backpressure

import (
	"errors"
	"time"
)

func For(v) {
	for {
		select {
		case v, ok := <-in:
			if !ok {
				in = nil
				continue
			}
		case v, ok := <-in:
			if !ok {
				in2 = nil
				continue
			}
		}
case <- done:
return
}
}

func TimeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close (done)
	}()
	select {
	case <- done:
		return result, err
	case <- time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {
	
}