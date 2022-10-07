package concurency

func CountTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func SearchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

func ProcessChannel(ch int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
		var out []int
		for i := 0; i < conc; i++ {
			out = append(out, <-results)
		}
		return out
	}
}
