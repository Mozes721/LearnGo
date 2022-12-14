package time

import (
	"fmt"
	"time"
)

func Time(t time.Time) error {
	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
	if err != nil {
		return err
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
}
