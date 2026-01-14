package main

import "fmt"

func main() {

	fmt.Println("-------------TASK 1---------------")
	//1
	pointer1()
	fmt.Println()

	fmt.Println("-------------TASK 2---------------")
	// 2
	a := 1
	b := 2
	Swap(&a, &b)
	fmt.Println()

	fmt.Println("-------------TASK 3---------------")
	// 3
	x := 3
	fmt.Println(x)
	Inc(&x)
	fmt.Println(x)
	IncN(&x, 5)
	fmt.Println(x)
	fmt.Println()

	fmt.Println("-------------TASK 4---------------")
	// 4
	Vagon1 := Counter{Value: 12}
	Vagon2 := Counter{Value: 15}
	Vagon1.Add(20)
	fmt.Println(Vagon1.Value)
	Vagon2.AddPtr(20)
	fmt.Println(Vagon2.Value)
	//
	Vagon1.AddPtr(25)
	fmt.Println(Vagon1.Value)
	Vagon2.Add(25)
	fmt.Println(Vagon2.Value)
	fmt.Println()

	fmt.Println("-------------TASK 5---------------")
	// 5
	bob := User{}
	bob.Name = "Bob"
	bob.age = 12
	fmt.Println(bob)
	Birthday(&bob)
	fmt.Println(bob)
	alex := NewUser("Alex", 34)
	fmt.Println(*alex)
	Birthday(alex)
	fmt.Println(*alex)
	fmt.Println()

	fmt.Println("-------------TASK 6---------------")
	// 6
	names := []string{"Ann", "Bob", "Cara"}
	users := []*User6{}
	for _, n := range names {
		users = append(users, &User6{Name: n})
	}
	fmt.Println(users)
	RenameAll(users, "Mr/Mrs.")
	for _, u := range users {
		fmt.Printf("%p -> %s\n", u, u.Name)
	}

}
