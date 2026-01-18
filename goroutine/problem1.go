package main

import (
	"fmt"
	"sync"
)

func prombel1() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Hello from #%d\n", n)
			fmt.Printf("func #%d finished\n", n)
			fmt.Println()
		}(i)
	}

	wg.Wait()
	fmt.Println("Wait to finish all goroutine")
}
