package concurency

func process(val int) int {
	// TODO:

	return 123
}

func RunThingConcurently(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
	}()
}
