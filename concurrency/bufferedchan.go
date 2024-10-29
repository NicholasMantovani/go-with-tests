package concurrency

import (
	"fmt"
	"sync"
)

func runBufferedChan() []int {
	resCh := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go performSomethingAndSometimesGoWrong(i, resCh, &wg)
	}

	wg.Wait()

	close(resCh)

	var results []int
	for v := range resCh {
		fmt.Println(v)
		results = append(results, v)
	}

	return results
}

func performSomethingAndSometimesGoWrong(input int, resultCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	if input%2 == 0 {
		return
	} else {
		resultCh <- input
	}
}
