package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}


func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdate)
}
