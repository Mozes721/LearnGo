package main

import (
	"context"
	. "learnGo/chapter12/cancelation"
	"os"
)

func main() {
	//ctx := context.Background()
	//result, err := contexts.Logic(ctx, "a string")
	ss := SlowServer()
	defer ss.Close()
	fs := FastServer()
	defer fs.Close()

	ctx := context.Background()
	CallBoth(ctx, os.Args[0], ss.URL, fs.URL)
}
