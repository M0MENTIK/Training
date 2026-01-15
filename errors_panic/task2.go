package main

import (
	"fmt"
	"io"
	"os"
)

/*

func main() {
	var nameFile string
	fmt.Print("Input file name: ")
	fmt.Fscan(os.Stdin, &nameFile)

	if err := ReadFile(nameFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File not found: %s\n", nameFile)
		} else {
			fmt.Println("Error:", err)
		}
		os.Exit(1)
	}

}

*/

func ReadFile(nameFile string) error {
	file, err := os.Open(nameFile)
	if err != nil {
		return err
	}

	defer file.Close()
	data := make([]byte, 128)
	for {
		n, err := file.Read(data)
		if n > 0 {
			fmt.Println(string(data[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

/*

- read file
- out body file
- if file is none, out understand message
- realise error "File doesn't find" and other

*/
