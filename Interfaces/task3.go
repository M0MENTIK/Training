package main

import "fmt"

type Printer interface {
	Print()
}

func PrintAll(items []Printer) {
	for _, item := range items {
		item.Print()
	}
}

//

//

type Document struct {
	DocumentText string
}

func (d Document) Print() {
	fmt.Println(d.DocumentText)
}

//

//

type Image struct {
	ImageText string
}

func (i Image) Print() {
	fmt.Println("Printing image:", i.ImageText)
}
