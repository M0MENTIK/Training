package main

import (
	"fmt"
	"sync"
)

func generator(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func worker(whereFrom <-chan int, where chan<- int, str chan<- struct{}, wg *sync.WaitGroup) {
	str <- struct{}{}
	defer wg.Done()
	for n := range whereFrom {

		where <- n * n
	}
}

func main() {
	var wg sync.WaitGroup
	started := make(chan struct{}, 3)
	jobs := make(chan int)
	results := make(chan int)
	go generator(jobs)

	wg.Add(3)
	go worker(jobs, results, started, &wg)
	go worker(jobs, results, started, &wg)
	go worker(jobs, results, started, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	<-started
	<-started
	<-started

	sum := 0
	for i := range results {
		sum += i
		fmt.Println(i)
	}
	fmt.Println("Sum:", sum)

}

//

//

//

//

//

//

//

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	done := make(chan struct{})

// 	for i := 1; i < 4; i++ {
// 		go func(id int) {
// 			fmt.Printf("Worker #%d started\n", id)
// 			fmt.Printf("Worker #%d finished\n", id)
// 			done <- struct{}{}
// 		}(i)
// 	}
// 	<-done
// 	<-done
// 	<-done

// 	fmt.Println("all workers done")

// }
