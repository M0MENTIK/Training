package main

import (
	"errors"
	"fmt"
	"math"
)

/*

func main() {
	printDivision(10, 5)
	printDivision(10, 0)
	printDivision(10, 0.00000001)
	printDivision(10, 1e-15)
	printDivision(10, 1e-6)
	printDivision(24, 12)
}

*/

var ErrTooSmallDivisionOrZero = errors.New("division is too small or equally zero")

func Divide(a, b float64) (float64, error) {

	const eps = 1e-9

	if math.Abs(b) < eps {
		return 0, ErrTooSmallDivisionOrZero
	}
	return a / b, nil

}

func printDivision(a, b float64) {
	r, err := Divide(a, b)
	if err != nil {
		if errors.Is(err, ErrTooSmallDivisionOrZero) {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Unexpected error:", err)
	}
	fmt.Printf("%.2f / %g = %.2f\n", a, b, r)
}
