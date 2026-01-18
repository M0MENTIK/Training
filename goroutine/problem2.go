package main

import "fmt"

func send(chf chan int, n int) {
	chf <- n
}

func problem2() {
	ch := make(chan int)

	go send(ch, 1)
	go send(ch, 2)
	go send(ch, 3)
	//close(ch)
	go send(ch, 4)
	go send(ch, 5)
	go send(ch, 6)
	go send(ch, 8)

	for v := range ch {
		fmt.Println(v)
	}

}
