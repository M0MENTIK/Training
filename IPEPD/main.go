package main

import "fmt"

func main() {

	// // 1
	// defer func() {
	// 	p := recover()
	// 	if p != nil {
	// 		fmt.Println("Was panec:", p)
	// 	}
	// }()
	// b := 5
	// var a *int = &b
	// fmt.Println(SafeDivide(6, 2, a))
	// fmt.Println(b)

	// 2
	store := NewMempryStore()
	if err := Upsert(store, "name", "Alex"); err != nil {
		fmt.Println("Upsert error:", err)
		return
	}

	name, err := store.Get("name")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	fmt.Println("name:", name)

	if err := store.Save("", "test"); err != nil {
		fmt.Println("Save error (expected):", err)
	}
	if _, err := store.Get("age"); err != nil {
		fmt.Println("Get error (expected):", err)
	}

	var s Saver = store
	_ = s.Save("city", "Berlin")
	city, _ := s.Get("city")
	fmt.Println("city:", city)
}
