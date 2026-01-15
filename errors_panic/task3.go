package main

import "fmt"

/*

func main() {
	SafeExecute(func() {
		TestDivide(10, 0)
	})
	fmt.Println()
	SafeExecute(func() {
		TestDivide(10, 5)
	})
}


*/

func TestDivide(a, b int) {
	fmt.Println(a / b)
}

func SafeExecute(fn func()) {
	defer func() {

		if p := recover(); p != nil {
			fmt.Println("Was panic:", p)
		} else {
			fmt.Println("Panic is none")
		}
	}()

	fmt.Println("Before")
	fn()
	fmt.Println("After")

}
