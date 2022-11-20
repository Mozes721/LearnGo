package main

import "learnGo/chapter12/contexts"

func main() {
	//ctx := context.Background()
	//result, err := contexts.Logic(ctx, "a string")


	//ss := SlowServer()
	//defer ss.Close()
	//fs := FastServer()
	//defer fs.Close()
	//
	//ctx := context.Background()
	//CallBoth(ctx, os.Args[0], ss.URL, fs.URL)

	bl :=contexts.BusinessLogic{
		RequestDecorator: tracker.Request,
		Logger: tracker.Logger{},
		Remote: "http://www.example.com/query"
	}

}
