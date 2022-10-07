package backpressure

import (
	"errors"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func DoThingThatShouldBeLimted() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func Pressure() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(DoThingThatShouldBeLimted()))
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
