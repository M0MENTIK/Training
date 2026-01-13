package main

import "fmt"

type Reader interface {
	Read() string
}
type Writer interface {
	Write(text string)
}

type File struct {
	text string
}

func (f *File) Write(text1 string) {
	f.text = text1
}
func (f *File) Read() string {
	return f.text
}

type ReadWriter interface {
	Reader
	Writer
}

func Process(rm ReadWriter, text string) {
	rm.Write(text)
	fmt.Println(rm.Read())
}
