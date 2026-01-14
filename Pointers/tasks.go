package main

import (
	"fmt"
)

// 1
func pointer1() {
	x := 10
	p := &x
	fmt.Println("Value x:", *p)
	fmt.Println("Adress x:", p)
	fmt.Println("Value p:", p)
	fmt.Println("Adress p:", &p)

	*p += 20
	fmt.Println("Value x:", *p)
	fmt.Println("Adress x:", p)
	fmt.Println("Value p:", p)
	fmt.Println("Adress p:", &p)
}

// 2
func Swap(a, b *int) {
	fmt.Println("Before: ", *a, *b)
	c := *a
	*a = *b
	*b = c
	fmt.Println("After: ", *a, *b)
}

// 3
func Inc(x *int) {
	*x++
}
func IncN(x *int, n int) {
	*x += n
}

// 4
type Counter struct{ Value int }

func (c Counter) Add(n int) {
	c.Value = n
}
func (c *Counter) AddPtr(n int) {
	c.Value = n
}

// 5
type User struct {
	Name string
	age  int
}

func NewUser(name string, age int) *User {
	u := &User{
		Name: name,
		age:  age,
	}

	return u
}

func Birthday(u *User) {
	if u != nil {
		u.age++
	} else {
		fmt.Println("Error. Pointer empty")
	}

}

// 6
type User6 struct {
	Name string
}

func RenameAll(users []*User6, prefix string) {
	for _, u := range users {
		u.Name = prefix + u.Name
	}
}
