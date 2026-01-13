package main

import "fmt"

type Speaker interface {
	Speak()
}

func AllSpeak(speaker Speaker) {
	fmt.Println("You can to talk")
	speaker.Speak()
}

//

//

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

//

//

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("May!")
}
